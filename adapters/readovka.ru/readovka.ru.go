package adapters

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sisteamnik/articler"
	prss "github.com/sisteamnik/articler/rss"
	"github.com/ungerik/go-dry"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	ReadovkaBaseUrl = "readovka.ru"
)

type ReadovkaParser struct {
	*articler.DefaultParser
}

func NewReadovkaParser() *ReadovkaParser {
	dp := articler.NewDefaultParser(ReadovkaBaseUrl)
	return &ReadovkaParser{DefaultParser: dp}
}

func (s *ReadovkaParser) Name() string {
	return "readovka.ru"
}

func (s *ReadovkaParser) IsArticle(u string) bool {
	//todo
	return true

	matched, _ := regexp.MatchString("^/[a-z-_0-9]*/[a-z-_0-9]*$", u)
	return matched
}

func (s *ReadovkaParser) LastArticles() ([]*url.URL, error) {
	rss, err := dry.FileGetBytes("http://" + ReadovkaBaseUrl + "/news-of-the-day/lenta?format=feed")
	if err != nil {
		return nil, err
	}
	items, err := prss.DecodeRss(rss)
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse("http://" + ReadovkaBaseUrl)
	var res []*url.URL
	for _, v := range items.Channel.Items {
		ur, _ := u.Parse(v.Link)
		res = append(res, ur)
	}
	return res, nil
}

func (p *ReadovkaParser) Parse(u string, bts []byte) (*articler.Article, error) {
	fmt.Print("here\n")
	a, _ := p.DefaultParser.Parse(bts)

	r := bytes.NewReader(bts)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	sBody := doc.Find(".itemFullText").Text()

	a.Text = strings.TrimSpace(sBody)
	//a.Description = doc.Find(".itemIntroText").Text()

	t, err := p.getDate(doc.Find("#main-content .item-date").Text())
	if err != nil {
		log.Println(err)
	}
	a.Published = t
	return a, err
}

func (p *ReadovkaParser) getDate(in string) (time.Time, error) {
	in = strings.ToLower(strings.TrimSpace(in))
	var months = map[string]string{
		"январь":   "01",
		"февраль":  "02",
		"март":     "03",
		"апрель":   "04",
		"май":      "05",
		"июнь":     "06",
		"июль":     "07",
		"август":   "08",
		"сентябрь": "09",
		"октябрь":  "10",
		"ноябрь":   "11",
		"декабрь":  "12",
	}
	var rep = []string{}
	for k, v := range months {
		rep = append(rep, k, v)
	}
	r := strings.NewReplacer(rep...)
	in = r.Replace(in)
	re := regexp.MustCompile("[0-9]{1,2} [0-9]{1,2} [0-9]{4} [0-9]{1,2}:[0-9]{1,2}")
	res := fmt.Sprintf("%s", re.FindString(in))
	return time.Parse("02 01 2006 15:04", res)
}

func init() {
	articler.RegisterArticleParser("readovka.ru", NewReadovkaParser())
}
