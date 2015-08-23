package adapters

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/sisteamnik/articler"
	"github.com/ungerik/go-dry"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	MvdBaseUrl = "67.mvd.ru"
	MvdScheme  = "https://"
)

type MvdParser struct {
	*articler.DefaultParser
}

func NewMvdParser() *MvdParser {
	dp := articler.NewDefaultParser(MvdBaseUrl)
	return &MvdParser{DefaultParser: dp}
}

func (s *MvdParser) Name() string {
	return "mvd"
}

func (s *MvdParser) IsArticle(u string) bool {
	//todo log error
	///news/item/3599309/
	matched, _ := regexp.MatchString("^/news/item/[0-9]*/$", u)
	return matched
}

func (s *MvdParser) LastArticles() ([]*url.URL, error) {
	rss, err := dry.FileGetBytes(MvdScheme + MvdBaseUrl + "/news/rss")
	if err != nil {
		return nil, err
	}
	items, err := DecodeRss(rss)
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse(MvdScheme + MvdBaseUrl)
	var res []*url.URL
	for _, v := range items.Channel.Items {
		ur, _ := u.Parse(v.Link)
		res = append(res, ur)
	}
	return res, nil
}

func (p *MvdParser) Parse(bts []byte) (*articler.Article, error) {
	a, _ := p.DefaultParser.Parse(bts)

	r := bytes.NewReader(bts)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	title := doc.Find("title").Text()
	a.Title = title

	body := ""
	sel := doc.Find(".article p")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		if i != 0 {
			body += "\n"
		}
		body += single.Text()
	}
	a.Body = []byte(body)

	t, err := p.getDate(doc.Find(".article-date-item").Text())
	if err != nil {
		log.Println(err)
	}
	a.Published = t
	return a, err
}

func (p *MvdParser) getDate(in string) (time.Time, error) {
	now := time.Now()
	arr := strings.Split(in, " ")
	switch len(arr) {
	case 2:
		t, err := time.Parse("Сегодня 15:04", in)
		if err != nil {
			return t, err
		}
		t.Date()
		rt := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, now.Location())
		return rt, nil
	case 3:
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
		t, err := time.Parse("02 01 15:04", in)
		if err != nil {
			return time.Time{}, err
		}
		rt := time.Date(now.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, now.Location())
		return rt, nil
	}
	return time.Time{}, errors.New("Unknow time error")
}
