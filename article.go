package articler

import (
	"fmt"
	"net/url"
	"sync"
	"time"
)

type Article struct {
	Title     string
	Text      string
	Published time.Time
	Source    string

	Parsed string

	/*	Images        []string
		Videos        []string
		Links         []string
		Authors       []string
		Tags          []string
		Catigories    []string
		Photographers []string*/
}

type Articles []*Article

func (a Articles) Len() int           { return len(a) }
func (a Articles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Articles) Less(i, j int) bool { return a[i].Published.Sub(a[j].Published) < 0 }

type ArticleParserAdapter interface {
	Parse(string, []byte) (*Article, error)
	IsArticle(string) bool
}

var (
	articleAdaptersMu sync.Mutex
	articleAdapters   = map[string]ArticleParserAdapter{}
)

func RegisterArticleParser(name string, adapter ArticleParserAdapter) {
	if adapter == nil {
		panic("adapter is nil")
	}
	articleAdaptersMu.Lock()
	defer articleAdaptersMu.Unlock()
	if _, dup := articleAdapters[name]; dup {
		panic("sql: Register called twice for adapter " + name)
	}
	articleAdapters[name] = adapter
}

var dp = &DefaultArticleParser{}

func ParseArticle(URL string, in []byte) (art *Article, e error) {
	var (
		u      *url.URL
		parser ArticleParserAdapter
	)
	u, e = url.Parse(URL)
	if e != nil {
		return
	}

	for k := range articleAdapters {
		if k == u.Host {
			parser = articleAdapters[k]
		}
	}

	if parser != nil && parser.IsArticle(u.RequestURI()) {
		return parser.Parse(URL, in)
	} else {
		if parser, ok := articleAdapters["default"]; ok {
			return parser.Parse(URL, in)
		}
		return dp.Parse(URL, in)
	}
	return nil, fmt.Errorf("Adapter not found")
}
