package crawler

import (
	"fmt"
	"net/url"
	"time"
)

type Site struct {
	Host    string
	Allowed bool
}

func NewSite(host string, allowed bool) *Site {
	return &Site{host, allowed}
}

type Link struct {
	SiteHost   string
	RequestUri string
	CrawlTime  time.Duration
	Created    time.Time
	Visited    time.Time
}

func NewLink(u *url.URL) *Link {
	return &Link{
		SiteHost:   u.Host,
		RequestUri: u.RequestURI(),
		Created:    time.Now(),
	}
}

func (l *Link) URL() *url.URL {
	u, _ := url.Parse("http://" + l.SiteHost + l.RequestUri)
	return u
}

func (l *Link) GetBlob(db Store) ([]byte, error) {
	if l.Visited.IsZero() {
		return nil, fmt.Errorf("not visited")
	}
	return db.GetBlob(l.URL().String())
}

/*func (l *Link) SetData(in []byte) error {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(in)
	e := w.Close()
	if e != nil {
		return e
	}
	l.Data = b.Bytes()
	return nil
}

func (l *Link) GetData() (res []byte, exists bool) {
	if l.Visited.IsZero() || len(l.Data) == 0 {
		return
	}
	var b bytes.Buffer
	r, e := gzip.NewReader(&b)
	if e != nil {
		color.Red(e.Error())
		return
	}
	bts, e := ioutil.ReadAll(r)
	return bts, e == nil
}*/
