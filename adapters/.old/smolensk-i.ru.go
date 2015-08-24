package adapters

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/advancedlogic/GoOse"
	"github.com/sisteamnik/articler"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	SmolenskIBaseUrl = "smolensk-i.ru"
)

type SmolenskIParser struct {
	*articler.DefaultParser
}

func NewSmolenskIParser() *SmolenskIParser {
	dp := articler.NewDefaultParser(SmolenskIBaseUrl)
	return &SmolenskIParser{DefaultParser: dp}
}

func (s *SmolenskIParser) Name() string {
	return "smolenski"
}

func (s *SmolenskIParser) IsArticle(u string) bool {
	//todo log error
	matched, _ := regexp.MatchString("^/[a-z-_0-9]*/[a-z-_0-9]*$", u)
	return matched
}

func (s *SmolenskIParser) LastArticles() ([]*url.URL, error) {
	f := time.Now().Format("/date/2006/01/02")
	doc, err := goquery.NewDocument("http://" + SmolenskIBaseUrl + f)
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse("http://" + SmolenskIBaseUrl)
	var res []*url.URL
	sel := doc.Find("article h1 a")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		ur, _ := u.Parse(single.AttrOr("href", "/"))
		res = append(res, ur)
	}
	return res, nil
}

func (p *SmolenskIParser) Parse(bts []byte) (*articler.Article, error) {
	a, _ := p.DefaultParser.Parse(bts)

	r := bytes.NewReader(bts)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sBody, _ := doc.Find(".entry-content").Html()

	g := goose.New()
	art := g.ExtractFromRawHtml("http://"+SmolenskIBaseUrl, sBody)

	a.Body = []byte(strings.TrimSpace(art.CleanedText))

	strTime := doc.Find("time").AttrOr("datetime", "")
	t, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		log.Println(err)
	}
	a.Published = t
	return a, err
}
