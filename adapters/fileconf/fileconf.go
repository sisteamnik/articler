package fileconf

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/sisteamnik/articler"
	"github.com/sisteamnik/articler/date"
	"github.com/sisteamnik/goose"
	"net/url"
)

type Parser struct {
	rules map[string]*Rule
}

func New(rules []*Rule) *Parser {
	p := new(Parser)
	p.rules = map[string]*Rule{}

	for _, v := range rules {
		p.rules[v.Host] = v
	}
	return p
}

func NewFromFile(filepath string) (*Parser, error) {
	rules, e := getRulesFromFile(filepath)
	if e != nil {
		return nil, e
	}
	return New(rules), nil
}

func (p *Parser) IsArticle(_ string) bool {
	return true
}

func (p *Parser) Parse(rawurl string, in []byte) (*articler.Article, error) {
	u, e := url.Parse(rawurl)
	if e != nil {
		return nil, e
	}
	for host := range p.rules {
		if host == u.Host {
			doc, e := goquery.NewDocumentFromReader(bytes.NewReader(in))
			if e != nil {
				return nil, e
			}
			return p.parse(u, p.rules[host], doc)
		}
	}
	color.Yellow("Run with goose, %s", u)
	g := goose.New()
	article := g.ExtractFromRawHtml(rawurl, string(in))
	return &articler.Article{Title: article.Title, Text: article.CleanedText}, nil
}

func (p *Parser) parse(u *url.URL, rule *Rule, doc *goquery.Document) (*articler.Article, error) {
	color.Green("Run from conf, %s", u)
	title := doc.Find(rule.TitleSelector).Text()

	publ, e := date.ExtractFromSelection(doc.Find(rule.PublishedSelector))
	if e != nil {
		return nil, e
	}

	return &articler.Article{
		Title:     title,
		Published: publ,
	}, nil
}
