package main

import (
	"github.com/sisteamnik/articler/crawler"
	"github.com/sisteamnik/articler/urltrack"
	"log"
	"net/url"
	"sync"
	"time"
)

var (
	dupMu sync.Mutex
	dup   = map[string]struct{}{}
)

func cb(rawurl string, data []byte) {
	u, e := url.Parse(rawurl)
	if e != nil {
		log.Fatalln(e)
		return
	}
	urls := crawler.ExtractLinks(u, data)
	for _, v := range urls {
		dupMu.Lock()
		if _, ok := dup[v.String()]; !ok {
			dup[v.String()] = struct{}{}
			if u.Host == v.Host {
				log.Printf("[new url] %s", v)
			}
		}
		dupMu.Unlock()
	}
}

func main() {
	s, e := urltrack.NewWithCommonInterval([]string{"http://golang.org", "https://zhuharev.ru"},
		5*time.Second, cb)
	if e != nil {
		panic(e)
	}
	s.Run()
}
