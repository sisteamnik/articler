package main

import (
	"fmt"
	"github.com/sisteamnik/articler/urltrack"
	"time"
)

func cb(in []byte) {
	fmt.Printf("%s", in)
}

func main() {
	s, e := urltrack.NewWithCommonInterval([]string{"http://golang.org", "https://zhuharev.ru"},
		5*time.Second, cb)
	if e != nil {
		panic(e)
	}
	s.Run()
}
