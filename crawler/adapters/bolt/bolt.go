package bolt

import (
	"bytes"
	"compress/gzip"
	"github.com/boltdb/bolt"
	"github.com/fatih/color"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/sisteamnik/articler/crawler"
	"io/ioutil"
	"os"
	"time"
)

var (
	LinkBucketName []byte = []byte("link")
	SiteBucketName []byte = []byte("site")
	BlobBucketName []byte = []byte("blob")
)

type BoltStore struct {
	db     *bolt.DB
	opened bool
}

func (b *BoltStore) Connect(filepath string) (e error) {
	b.db, e = bolt.Open(filepath, 0600, nil)
	if e != nil {
		return e
	}
	e = b.db.Update(func(tx *bolt.Tx) error {
		_, e = tx.CreateBucketIfNotExists(LinkBucketName)
		if e != nil {
			return e
		}
		_, e = tx.CreateBucketIfNotExists(SiteBucketName)
		if e != nil {
			return e
		}
		_, e = tx.CreateBucketIfNotExists(BlobBucketName)
		if e != nil {
			return e
		}
		return nil
	})

	go func() {
		// Grab the initial stats.
		prev := b.db.Stats()

		for {

			// Grab the current stats and diff them.
			stats := b.db.Stats()
			diff := stats.Sub(&prev)

			// Encode stats to JSON and print to STDERR.
			ffjson.NewEncoder(os.Stderr).Encode(diff)
			// Save stats for the next loop.
			prev = stats

			// Wait for 10s.
			time.Sleep(60 * time.Second)
		}
	}()

	return
}

func (b *BoltStore) SaveLink(l *crawler.Link) error {
	bts, e := ffjson.Marshal(l)
	if e != nil {
		color.Red(e.Error())
		return e
	}
	e = b.db.Update(func(tx *bolt.Tx) error {
		bu := tx.Bucket(LinkBucketName)
		return bu.Put([]byte(l.RequestUri), bts)
	})
	return e
}

func (b *BoltStore) GetLink(rawurl string) (*crawler.Link, bool) {
	var (
		bts []byte
		res *crawler.Link = &crawler.Link{}
	)
	e := b.db.View(func(tx *bolt.Tx) error {
		bu := tx.Bucket(LinkBucketName)
		bts = bu.Get([]byte(rawurl))
		if len(bts) == 0 {
			return crawler.ErrNotFound
		}
		return nil
	})
	if e != nil {
		return nil, false
	}
	e = ffjson.Unmarshal(bts, res)
	if e != nil {
		color.Red(e.Error())
		return nil, false
	}
	return res, true
}

func (b *BoltStore) CountLink() (count int) {
	b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(LinkBucketName)
		b.ForEach(func(_, _ []byte) error {
			count++
			return nil
		})
		return nil
	})
	return
}

func (b *BoltStore) Visited() (res []*crawler.Link, e error) {
	b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(LinkBucketName)
		e = b.ForEach(func(_, data []byte) error {
			l := &crawler.Link{}
			e = ffjson.Unmarshal(data, l)
			if e != nil {
				return e
			}
			if !l.Visited.IsZero() {
				res = append(res, l)
			}
			return nil
		})
		return e
	})
	return
}

func (b *BoltStore) NotVisited() (res []*crawler.Link, e error) {
	b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(LinkBucketName)
		e = b.ForEach(func(_, data []byte) error {
			l := &crawler.Link{}
			e = ffjson.Unmarshal(data, l)
			if e != nil {
				return e
			}
			if l.Visited.IsZero() {
				res = append(res, l)
			}
			return nil
		})
		return e
	})
	return
}

func (b *BoltStore) SaveSite(s *crawler.Site) error {
	bts, e := ffjson.Marshal(s)
	if e != nil {
		return e
	}
	e = b.db.Update(func(tx *bolt.Tx) error {
		bu := tx.Bucket(SiteBucketName)
		return bu.Put([]byte(s.Host), bts)
	})
	return e
}

func (b *BoltStore) GetSite(host string) (*crawler.Site, bool) {
	var (
		bts []byte
		res *crawler.Site = &crawler.Site{}
	)
	e := b.db.View(func(tx *bolt.Tx) error {
		bu := tx.Bucket(SiteBucketName)
		bts = bu.Get([]byte(host))
		if len(bts) == 0 {
			return crawler.ErrNotFound
		}
		return nil
	})
	if e != nil {
		return nil, false
	}
	e = ffjson.Unmarshal(bts, res)
	if e != nil {
		color.Red(e.Error())
		return nil, false
	}
	return res, true
}

func (b *BoltStore) CountSite() (count int) {
	b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(SiteBucketName)
		b.ForEach(func(_, _ []byte) error {
			count++
			return nil
		})
		return nil
	})
	return
}

func (b *BoltStore) AllowedSites() (res []*crawler.Site, e error) {
	b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(SiteBucketName)
		e = b.ForEach(func(k, v []byte) error {
			var s = &crawler.Site{}
			e = ffjson.Unmarshal(v, s)
			if e != nil {
				return e
			}
			if s.Allowed {
				res = append(res, s)
			}
			return e
		})
		return e
	})
	return
}

func (b *BoltStore) SaveBlob(key string, in []byte) (e error) {
	e = b.db.Update(func(tx *bolt.Tx) error {
		bu := tx.Bucket(BlobBucketName)
		gzipped, e := gz(in)
		if e != nil {
			return e
		}
		return bu.Put([]byte(key), gzipped)
	})
	return e
}

func (b *BoltStore) GetBlob(key string) ([]byte, error) {
	var (
		res []byte
		e   error
	)
	e = b.db.View(func(tx *bolt.Tx) error {
		bu := tx.Bucket(BlobBucketName)
		bts := bu.Get([]byte(key))
		if len(bts) == 0 {
			return crawler.ErrNotFound
		}
		res, e = guz(bts)
		if e != nil {
			return e
		}
		return nil
	})
	if e != nil {
		return nil, e
	}
	return res, nil
}

func init() {
	crawler.Register("bolt", &BoltStore{})
}

func gz(in []byte) (out []byte, e error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, e = w.Write(in)
	if e != nil {
		return
	}
	e = w.Close()
	if e != nil {
		return
	}
	return b.Bytes(), nil
}

func guz(in []byte) (out []byte, e error) {
	rdr := bytes.NewReader(in)
	r, e := gzip.NewReader(rdr)
	if e != nil {
		return nil, e
	}
	out, e = ioutil.ReadAll(r)
	if e != nil {
		return nil, e
	}
	r.Close()
	return
}
