package articler

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html"
	"regexp"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

func DecodeRss(in []byte) (Rss, error) {
	//in = escapeDescription(in)
	rss := Rss{}
	err := xml.Unmarshal(in, &rss)
	return rss, err
}

func escapeDescription(in []byte) []byte {
	return escapeTag("description", in)
}

func escapeTag(tag string, in []byte) []byte {
	reg := fmt.Sprintf("<%s>(.*?)</%s>", tag, tag)
	re := regexp.MustCompile(reg)
	res := re.FindAllSubmatch(in, -1)
	for _, v := range res {
		if len(v) == 2 {
			in = bytes.Replace(in, v[1],
				[]byte(html.EscapeString(string(v[1]))), -1)
		}
	}
	return in
}
