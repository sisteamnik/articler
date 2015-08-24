package adapters

import (
	"encoding/xml"
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
	rss := Rss{}
	err := xml.Unmarshal(in, &rss)
	return rss, err
}
