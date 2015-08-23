package adapters

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/sisteamnik/articler"
	"github.com/sisteamnik/goose"
	"github.com/ungerik/go-dry"
	"log"
	"net/url"
	"regexp"
	"runtime"
	"time"
)

type SuperParser struct {
	cnf *articler.AdapterConfig
}

func NewParser(cnf *articler.AdapterConfig) (*SuperParser, error) {
	sp := &SuperParser{cnf: cnf}
	if cnf.Scheme == "" {
		sp.cnf.Scheme = "http"
	}
	if cnf.TitleSelector == "" && cnf.TitleExtractFunc == nil {
		cnf.TitleSelector = "title"
	}
	if cnf.FeedSelector == "" {
		cnf.FeedSelector = "a"
	}
	return sp, nil
}

func (p *SuperParser) BaseUrl() *url.URL {
	return &url.URL{Scheme: p.cnf.Scheme, Host: p.cnf.Host}
}

func (p *SuperParser) Name() string {
	return p.cnf.Name
}

func (p *SuperParser) Domain() string {
	return p.cnf.Host
}

func (p *SuperParser) LastArticles() ([]*url.URL, error) {
	u := p.BaseUrl()
	if p.cnf.FeedUriGenerate != nil {
		u = p.cnf.FeedUriGenerate()
	}
	//todo fix if query
	if p.cnf.FeedUri != "" {
		u.Path = p.cnf.FeedUri
	}
	bts, err := dry.FileGetBytes(u.String())
	if err != nil {
		return nil, err
	}
	return p.ExtractFeedLinks(bts)
}

func (p *SuperParser) IsArticle(u string) bool {
	re := regexp.MustCompile(p.cnf.ArticleUriRegex)
	return re.MatchString(u)
}

func (p *SuperParser) Parse(in []byte) (art *articler.Article, err error) {
	if p.cnf.ParseFunc != nil {
		return p.cnf.ParseFunc(in)
	}
	art = &articler.Article{}

	title, err := extractBySelectorOrFunc(p.cnf.TitleSelector,
		p.cnf.TitleExtractFunc, in)

	body, err := extractBySelectorOrFunc(p.cnf.BodySelector,
		p.cnf.BodyExtractFunc, in)

	date, err := extractTime(p.cnf.DateSelector, p.cnf.DateFormat, p.cnf.DateRegex,
		p.cnf.DateExtractFunc, in)

	art.Title = string(title)
	art.Body = body
	art.Published = date
	if err != nil {
		return art, err
	}

	return art, nil
}

func (p *SuperParser) ExtractFeedLinks(in []byte) (res []*url.URL, err error) {
	switch p.cnf.FeedType {
	case "rss":
		feed, err := DecodeRss(in)
		if err != nil {
			return nil, err
		}
		for _, v := range feed.Channel.Items {
			u, err := p.BaseUrl().Parse(v.Link)
			if err != nil {
				log.Println(err)
				continue
			}
			if p.IsArticle(u.RequestURI()) {
				res = append(res, u)
			}
		}
		return res, nil
	default:
		if p.cnf.FeedSelector != "" {
			return p.extractLinksBySelector(p.BaseUrl(), p.cnf.FeedSelector, in)
		}
		if p.cnf.FeedExtract != nil {
			return p.cnf.FeedExtract(in)
		}
	}
	return nil, errors.New("Unknown error")
}

//utils

func (p *SuperParser) extractLinksBySelector(base *url.URL, selector string, in []byte) ([]*url.URL, error) {
	r := bytes.NewReader(in)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var res []*url.URL
	sel := doc.Find(selector)
	for i := range sel.Nodes {
		single := sel.Eq(i)
		ur, _ := base.Parse(single.AttrOr("href", "/"))
		if p.IsArticle(ur.RequestURI()) {
			res = append(res, ur)
		}
	}
	return res, nil
}

func extractBySelectorOrFunc(selector string, eF articler.ExtractFunc,
	in []byte) ([]byte, error) {

	if selector != "" {
		r := bytes.NewReader(in)
		doc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			return nil, err
		}
		htm, err := doc.Find(selector).Html()
		if err != nil {
			return nil, err
		}
		out := goose.Clean([]byte(htm))
		return out, nil
	}

	if eF != nil {
		return eF(in)
	}

	return nil, errors.New("Selector or func not setted")
}

//todo fix bug with zero minutes or seconds parsing
func extractTime(selector, format, reg string, eF articler.ExtractTimeFunc,
	in []byte) (time.Time, error) {

	in = fixRussianDate(in)

	if IsRelativeTime(string(in)) {
		in = []byte(FixRelativeTime(string(in)))
	}

	if selector != "" {
		r := bytes.NewReader(in)
		doc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			_, fl, ln, _ := runtime.Caller(0)
			logErr(err, fl, ln)
			return time.Time{}, err
		}

		zone := doc.Find(selector)
		dateTimeAttr := zone.AttrOr("datetime", "")
		if dateTimeAttr == "" {
			zoneStr, err := zone.Html()
			if err != nil {
				_, fl, ln, _ := runtime.Caller(0)
				logErr(err, fl, ln)
				return time.Time{}, err
			}
			in = []byte((zoneStr))

		} else {
			in = []byte(dateTimeAttr)
		}
	}
	out := []string{}
	//clean input
	if reg != "" {
		re, err := regexp.Compile(reg)
		if err != nil {
			_, fl, ln, _ := runtime.Caller(0)
			logErr(err, fl, ln)
			return time.Time{}, err
		}
		out = re.FindStringSubmatch(string(in))
		if len(out) < 2 {
			_, fl, ln, _ := runtime.Caller(0)
			logErr(errors.New("So bad time"), fl, ln)
			return time.Time{}, errors.New("So bad time")
		}
		in = []byte(out[1])
	}

	if format != "" {
		TimeFormats = append(TimeFormats, format)
		return time.Parse(format, string(in))
	}

	if eF != nil {
		return eF(in)
	}

	return ParseTime(string(in))
}

func extractAttrBySelector(selector, attrName string, in []byte) (string, error) {
	r := bytes.NewReader(in)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}

	return doc.Find(selector).AttrOr(attrName, ""), nil

}

//code from https://github.com/codelucas/newspaper/blob/master/newspaper/extractors.py
//3 strategies for publishing date extraction. The strategies
//        are descending in accuracy and the next strategy is only
//        attempted if a preferred one fails.
//        1. Pubdate from URL
//        2. Pubdate from metadata

var dateRegexp = `([\./\-_]{0,1}(19|20)\d{2})[\./\-_]{0,1}(([0-3]{0,1}[0-9][\./\-_])|(\w{3,5}[\./\-_]))([0-3]{0,1}[0-9][\./\-]{0,1})?`
var DateRegexp = regexp.MustCompile(dateRegexp)
var usualyDateRegexp = `(\d{1,2}(\.| )[A-zА-я\d]{2,6}(\.| )\d{4} \d{2}:\d{2})`
var UsualyDateRegexp = regexp.MustCompile(usualyDateRegexp)
var relDateRegexp = `[А-я]+(,| в)? \d{2}:\d{2}`
var RelDateRegexp = regexp.MustCompile(relDateRegexp)

//todo compare fulled times beetwen diferent extraction methods
func getPublishingDate(ctx *url.URL, in []byte) (time.Time, error) {
	res := DateRegexp.FindString(ctx.String())
	if res != "" {
		date, err := ParseTime(res)
		if err == nil {
			return date, err
		}
	}

	reader := bytes.NewReader(in)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return time.Now(), err
	}

	var selectors = map[string]string{
		"time": "datetime",

		//smoldaily.ru
		".meta-post .date": "text",

		//meta
		"[property=\"rnews:datePublished\"]": "content",
		//"[property=\"article:published_time\"]": "content",
		"[name=OriginalPublicationDate]":   "content",
		"[itemprop=datePublished]":         "datetime",
		"[property=\"og:published_time\"]": "content",
		"[name=article_date_original]":     "content",
		"[name=publication_date]":          "content",
		"[name=\"sailthru.date\"]":         "content",
		"[name=PublishDate]":               "content",
	}

	for sel, attr := range selectors {
		if s := doc.Find(sel); s.Length() > 0 {
			if attr == "text" {
				return ParseTime(s.Text())
			} else if strTime, ok := s.Attr(attr); ok {
				return ParseTime(strTime)
			}
		}
	}

	strTime := RelDateRegexp.FindString(string(in))
	if strTime != "" {
		return ParseTime(strTime)
	}

	strTime = UsualyDateRegexp.FindString(string(in))
	if strTime != "" {
		return ParseTime(strTime)
	}
	return time.Now(), errors.New("Unknown error")
}

/*

def get_publishing_date(self, url, doc):
        """
        """

        def parse_date_str(date_str):
            try:
                datetime_obj = date_parser(date_str)
                return datetime_obj
            except:
                # near all parse failures are due to URL dates without a day
                # specifier, e.g. /2014/04/
                return None

        date_match = re.search(urls.DATE_REGEX, url)
        if date_match:
            date_str = date_match.group(0)
            datetime_obj = parse_date_str(date_str)
            if datetime_obj:
                return datetime_obj

        PUBLISH_DATE_TAGS = [
            {'attribute': 'property', 'value': 'rnews:datePublished', 'content': 'content'},
            {'attribute': 'property', 'value': 'article:published_time', 'content': 'content'},
            {'attribute': 'name', 'value': 'OriginalPublicationDate', 'content': 'content'},
            {'attribute': 'itemprop', 'value': 'datePublished', 'content': 'datetime'},
            {'attribute': 'property', 'value': 'og:published_time', 'content': 'content'},
            {'attribute': 'name', 'value': 'article_date_original', 'content': 'content'},
            {'attribute': 'name', 'value': 'publication_date', 'content': 'content'},
            {'attribute': 'name', 'value': 'sailthru.date', 'content': 'content'},
            {'attribute': 'name', 'value': 'PublishDate', 'content': 'content'},
        ]
        for known_meta_tag in PUBLISH_DATE_TAGS:
            meta_tags = self.parser.getElementsByTag(
                doc,
                attr=known_meta_tag['attribute'],
                value=known_meta_tag['value'])
            if meta_tags:
                date_str = self.parser.getAttribute(
                    meta_tags[0],
                    known_meta_tag['content'])
                datetime_obj = parse_date_str(date_str)
                if datetime_obj:
                    return datetime_obj

        return None

*/
