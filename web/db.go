package main

import (
	"encoding/json"
	"github.com/sisteamnik/articler"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"strings"
)

type DB struct {
	db *leveldb.DB
}

func NewDb(fname string) (*DB, error) {
	db, err := leveldb.OpenFile(fname, nil)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, err
}

func (db *DB) Visited(u string) bool {
	if data, err := db.db.Get([]byte(u), nil); err == nil &&
		"true" == string(data) {
		return true
	}
	return false
}

func (db *DB) Visit(u string) error {
	return db.db.Put([]byte(u), []byte("true"), nil)
}

var ArticlePrefix = "article"

func (db *DB) Save(a *articler.Article) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return db.db.Put([]byte(ArticlePrefix+a.Source), data, nil)
}

func (db *DB) Get(url string) (art *articler.Article, err error) {
	art = &articler.Article{}
	data, err := db.db.Get([]byte(ArticlePrefix+url), nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, art)
	return
}

func (db *DB) GetAll() (res []*articler.Article) {
	iter := db.db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()

		if strings.HasPrefix(string(key), ArticlePrefix) {
			var art *articler.Article = &articler.Article{}
			err := json.Unmarshal(value, art)
			if err != nil {
				log.Println("ERROR", err)
				break
			}
			res = append(res, art)
		}
	}
	iter.Release()
	return res
}
