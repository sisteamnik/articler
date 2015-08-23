package adapters

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

func extractLinks(ctx *url.URL, doc *goquery.Document) []*url.URL {
	var res []*url.URL
	var dup = map[string]bool{}
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("href")
		// Resolve address
		u, err := ctx.Parse(val)
		if err != nil {
			fmt.Printf("error: resolve URL %s - %s\n", val, err)
			return
		}
		if !dup[u.String()] {
			res = append(res, u)
			dup[u.String()] = true
		}
	})
	return res
}
