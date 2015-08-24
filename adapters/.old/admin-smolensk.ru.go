package adapters

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/sisteamnik/articler"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	AdminBaseUrl = "admin-smolensk.ru"
)

type AdminParser struct {
	*articler.DefaultParser
}

func NewAdminParser() *AdminParser {
	dp := articler.NewDefaultParser(AdminBaseUrl)
	return &AdminParser{DefaultParser: dp}
}

func (s *AdminParser) Name() string {
	return "admin"
}

func (s *AdminParser) IsArticle(u string) bool {
	//todo log error
	matched, _ := regexp.MatchString("^/news/news_[0-9]*\\.html$", u)
	return matched
}

func (s *AdminParser) LastArticles() ([]*url.URL, error) {
	doc, err := goquery.NewDocument("http://" + AdminBaseUrl)
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse("http://" + AdminBaseUrl)

	var res []*url.URL
	sel := doc.Find("a")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		ur, _ := u.Parse(single.AttrOr("href", "/"))
		if s.IsArticle(ur.Path) {
			res = append(res, ur)
		}
	}
	return res, nil
}

func (p *AdminParser) Parse(bts []byte) (*articler.Article, error) {
	a, _ := p.DefaultParser.Parse(bts)

	r := bytes.NewReader(bts)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body := ""
	sel := doc.Find(".b-editor p")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		if i != 0 {
			body += "\n"
		}
		body += single.Text()
	}
	a.Body = []byte(strings.TrimSpace(body))

	t, err := p.getDate(doc.Text())
	if err != nil {
		log.Println(err)
	}
	a.Published = t
	log.Println(a.Published)
	return a, err
}

func (p *AdminParser) getDate(in string) (time.Time, error) {
	re := regexp.MustCompile("Дата последнего изменения ([0-9]{1,2}\\.[0-9]{1,2}\\.[0-9]{4} [0-9]{1,2}:[0-9]{1,2})")
	out := re.FindStringSubmatch(in)
	if len(out) < 2 {
		return time.Time{}, errors.New("So bad time")
	}
	return time.Parse("02.01.2006 15:04", out[1])
}
