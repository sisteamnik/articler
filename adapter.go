package articler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/advancedlogic/GoOse"
	"net/url"
)

type Adapter interface {
	Domain() string
	LastArticles() ([]*url.URL, error)
	//ExtractFeedLinks([]byte) ([]*url.URL, error)
	Parse([]byte) (*Article, error)
	IsArticle(string) bool
	Name() string

	//BaseUrl() *url.Url
	//FeedType() string //html or rss
	//FeedUrl() string //for external fetcher
	//ExtractLinks(in []byte) ([]*url.URL, error)

	/*
			Example init new adapter
		adapter,err :=  NewAdapter(Config{
		Name:"smolenski",
		BaseUrl :""
		})


	*/
}

type DefaultParser struct {
	baseUrl string
}

func NewDefaultParser(baseUrl string) *DefaultParser {
	return &DefaultParser{baseUrl}
}

func (s *DefaultParser) Domain() string {
	return s.baseUrl
}

func (s *DefaultParser) Name() string {
	return "default"
}

func (s *DefaultParser) IsArticle(u string) bool {
	return true
}

func (s *DefaultParser) LastArticles() ([]*url.URL, error) {
	doc, err := goquery.NewDocument(s.baseUrl)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(s.baseUrl)
	if err != nil {
		return nil, err
	}
	return extractLinks(u, doc), nil
}

func (s *DefaultParser) Parse(bts []byte) (*Article, error) {
	a := &Article{}
	g := goose.New()
	art := g.ExtractFromRawHtml(s.baseUrl, string(bts))

	a.Title = art.Title

	a.Text = art.CleanedText

	/*doc, err := readability.NewDocument(string(bts))
	if err != nil {
		log.Println(err)
	}

	a.Body = []byte(doc.Content())*/
	return a, nil
}

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
