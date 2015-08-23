package articler

import (
	"gopkg.in/yaml.v2"
	"net/url"
	"time"
)

type AdapterConfig struct {
	Host   string `yaml:"host,omitempty"`
	Scheme string
	Name   string

	FeedUri         string
	FeedType        string //html or rss
	FeedUriGenerate FeedUriGenerateFunc
	FeedExtract     FeedExtractFunc

	//FeedType must be html
	//this is link's selector, like "main a"
	FeedSelector string

	ArticleUriRegex string
	ParseFunc       ExtractArticleFunc

	TitleSelector    string
	TitleExtractFunc ExtractFunc

	BodySelector    string
	BodyExtractFunc ExtractFunc

	DateSelector    string
	DateFormat      string
	DateRegex       string
	DateExtractFunc ExtractTimeFunc
}

func ParseAdapterConfig(in []byte) (*AdapterConfig, error) {
	ac := &AdapterConfig{}
	err := yaml.Unmarshal(in, &ac)
	return ac, err
}

type ExtractFunc func([]byte) ([]byte, error)

type ExtractTimeFunc func([]byte) (time.Time, error)

type ExtractArticleFunc func([]byte) (*Article, error)

type FeedUriGenerateFunc func() *url.URL

type FeedExtractFunc func([]byte) ([]*url.URL, error)
