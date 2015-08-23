package adapters

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sisteamnik/articler"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	FasBaseUrl = "smolensk.fas.gov.ru"
)

type FasParser struct {
	*articler.DefaultParser
}

func NewFasParser() *FasParser {
	dp := articler.NewDefaultParser(FasBaseUrl)
	return &FasParser{DefaultParser: dp}
}

func (s *FasParser) Name() string {
	return "fas"
}

func (s *FasParser) IsArticle(u string) bool {
	//todo log error
	matched, _ := regexp.MatchString("^/news/[0-9]*$", u)
	return matched
}

func (s *FasParser) LastArticles() ([]*url.URL, error) {
	doc, err := goquery.NewDocument("http://" + FasBaseUrl + "/news")
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse("http://" + FasBaseUrl)
	var res []*url.URL
	sel := doc.Find(".news-container h3 a")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		ur, _ := u.Parse(single.AttrOr("href", "/"))
		res = append(res, ur)
	}
	return res, nil
}

func (p *FasParser) Parse(bts []byte) (*articler.Article, error) {
	a, _ := p.DefaultParser.Parse(bts)

	r := bytes.NewReader(bts)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	sBody := doc.Find(".text-version").Text()

	a.Body = []byte(strings.TrimSpace(sBody))

	t, err := p.getDate(doc.Find(".submitted").Text())
	if err != nil {
		log.Println(err)
	}
	a.Published = t
	return a, err
}

func (p *FasParser) getDate(in string) (time.Time, error) {
	in = strings.ToLower(strings.TrimSpace(in))
	var months = map[string]string{
		"января":   "01",
		"февраля":  "02",
		"марта":    "03",
		"апреля":   "04",
		"мая":      "05",
		"июня":     "06",
		"июля":     "07",
		"августа":  "08",
		"сентября": "09",
		"октября":  "10",
		"ноября":   "11",
		"декабря":  "12",
	}
	var rep = []string{}
	for k, v := range months {
		rep = append(rep, k, v)
	}
	r := strings.NewReplacer(rep...)
	in = r.Replace(in)
	re := regexp.MustCompile("[0-9]{1,2} [0-9]{1,2} [0-9]{4}, [0-9]{1,2}:[0-9]{1,2}")
	res := fmt.Sprintf("%s", re.FindString(in))
	return time.Parse("02 01 2006, 15:04", res)
}
