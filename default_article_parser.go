package articler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/sisteamnik/articler/date"
	"github.com/sisteamnik/goose"
	"github.com/ungerik/go-dry"
	"net/url"
	"reflect"
	"strings"
)

type DefaultArticleParser struct {
	rules map[string]*Rule
}

func NewDefaultArticleParser() *DefaultArticleParser {
	p := &DefaultArticleParser{}
	p.rules = map[string]*Rule{}

	return p
}

func NewFromFile(filepath string) (*DefaultArticleParser, error) {
	p := NewDefaultArticleParser()
	return p, p.LoadRules(filepath)
}

func (p *DefaultArticleParser) IsArticle(_ string) bool {
	return true
}

func (p *DefaultArticleParser) Parse(rawurl string, in []byte) (*Article, error) {
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
	g := goose.New()
	article := g.ExtractFromRawHtml(rawurl, string(in))
	return &Article{Title: article.Title, Text: article.CleanedText,
		Parsed: "default"}, nil
}

func (p *DefaultArticleParser) parse(u *url.URL, rule *Rule, doc *goquery.Document) (*Article, error) {
	color.Green("Run from conf, %s", u)
	title := doc.Find(rule.TitleSelector).Text()

	publ, e := date.ExtractFromSelection(doc.Find(rule.PublishedSelector))
	if e != nil {
		return nil, e
	}

	text := doc.Find(rule.TextSelector).Text()

	return &Article{
		Title:     title,
		Published: publ,
		Text:      text,
		Parsed:    "default",
	}, nil
}

func (p *DefaultArticleParser) LoadRules(filepath string) error {
	if p.rules == nil {
		p.rules = map[string]*Rule{}
	}
	rules, e := getRulesFromFile(filepath)
	if e != nil {
		return e
	}
	for _, v := range rules {
		p.rules[v.Host] = v
	}

	return nil
}

type Rule struct {
	Host string

	TitleSelector     string
	TextSelector      string
	PublishedSelector string
}

func getRulesFromFile(filename string) (res []*Rule, e error) {
	var (
		lines []string
	)

	lines, e = dry.FileGetLines(filename)
	if e != nil {
		return
	}
	for _, v := range lines {
		rule, e := parseRule(v)
		if e != nil {
			continue
		}
		res = append(res, rule)
	}
	return
}

var (
	ErrCommentLine error = fmt.Errorf("line is comment")
	ErrEmptyLine         = fmt.Errorf("emty line")
	ErrEmptyRule         = fmt.Errorf("emty rule")
)

func parseRule(line string) (*Rule, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, ErrEmptyLine
	}
	if line[0] == "#"[0] {
		return nil, ErrCommentLine
	}

	arr := strings.SplitN(line, " ", 2)
	if len(arr) != 2 {
		return nil, ErrEmptyRule
	}

	rule := Rule{
		Host: arr[0],
	}

	st := reflect.StructTag(arr[1])
	rule.TextSelector = st.Get("text")
	rule.TitleSelector = st.Get("title")
	rule.PublishedSelector = st.Get("publ")
	return &rule, nil
}
