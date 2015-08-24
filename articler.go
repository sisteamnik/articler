package articler

import (
	"io/ioutil"
)

type Articler struct {
	conf    *Config
	fetcher Fetcher

	defaultArticleParser *DefaultArticleParser
}

func New(configs ...*Config) (art *Articler, e error) {
	art = &Articler{}
	art.defaultArticleParser = NewDefaultArticleParser()
	RegisterArticleParser("default", art.defaultArticleParser)
	art.fetcher = &DefaultFetcher{}
	if len(configs) == 1 {
		art.conf = configs[0]
		if filepath := art.conf.DefaultArticleParserConf; filepath != "" {
			e = art.defaultArticleParser.LoadRules(filepath)
			if e != nil {
				return
			}
		}
	}
	return
}

/*func NewFromFile(filepath string) (art *Articler, e error) {

}*/

func (a *Articler) ParseArticleFromUrl(URL string) (*Article, error) {
	resp, e := a.fetcher.Get(URL)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	bts, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}

	return ParseArticle(URL, bts)
}
