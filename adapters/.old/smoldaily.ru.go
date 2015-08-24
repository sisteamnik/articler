package adapters

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/sisteamnik/articler"
	"github.com/sisteamnik/sitemap"
	"github.com/ungerik/go-dry"
	"log"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	SmoldailyBaseUrl = "smoldaily.ru"
)

type SmoldailyParser struct {
}

func (s *SmoldailyParser) Name() string {
	return "smoldaily"
}

func (s *SmoldailyParser) Domain() string {
	return SmoldailyBaseUrl
}

func (s *SmoldailyParser) IsArticle(u string) bool {
	//todo log error
	matched, _ := regexp.MatchString("^/[a-z-_0-9]*$", u)
	return matched
}

func (s *SmoldailyParser) LastArticles() ([]*url.URL, error) {
	bts, err := dry.FileGetBytes("http://" + SmoldailyBaseUrl + "/sitemap.xml")
	if err != nil {
		return nil, err
	}
	items, err := sitemap.ParseIndex(bts)
	if err != nil {
		return nil, err
	}
	needFetch := ""
	sort.Sort(sort.Reverse(sitemap.ByTime(items)))
	for _, v := range items {
		if strings.Contains(v.Loc, "sitemap-misc.xml") ||
			strings.Contains(v.Loc, "sitemap-tax-category.xml") {
			continue
		}
		needFetch = v.Loc
		break
	}
	log.Println(needFetch)
	bts, err = dry.FileGetBytes(needFetch)
	if err != nil {
		return nil, err
	}
	items, err = sitemap.Parse(bts)
	if err != nil {
		return nil, err
	}
	var res []*url.URL
	for _, v := range items {
		u, _ := url.Parse(v.Loc)
		if u != nil {
			res = append(res, u)
		}
	}
	return res, nil
}

func (s *SmoldailyParser) Parse(bts []byte) (*articler.Article, error) {
	a := &articler.Article{}
	r := bytes.NewReader(bts)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	a.Title = doc.Find("h1").Text()

	body := ""
	sel := doc.Find(".mainpost p")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		if i != 0 {
			body += "\n"
		}
		body += single.Text()
	}
	a.Body = []byte(strings.TrimSpace(body))

	strTime := doc.Find(".meta-post .date span").Text()
	t, err := time.Parse("02.01.2006, 15:04", strTime)
	if err != nil {
		log.Println(err)
	}
	a.Published = t

	return a, nil
}
