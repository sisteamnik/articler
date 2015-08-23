package articler

import (
	"errors"
	"github.com/PuerkitoBio/fetchbot"
	"github.com/fatih/color"
	"github.com/ungerik/go-dry"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

type Parser struct {
	articles chan *Article

	adapters map[string]Adapter

	interval time.Duration

	fetcher *fetchbot.Fetcher
	queue   *fetchbot.Queue

	db DB
}

func NewParser(db DB) (*Parser, error) {
	p := &Parser{
		interval: 5 * time.Minute,
		adapters: map[string]Adapter{},
		articles: make(chan *Article, 10),
		db:       db,
	}
	fb := fetchbot.New(fetchbot.HandlerFunc(p.handler))
	fb.UserAgent = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"
	queue := fb.Start()
	p.fetcher = fb
	p.queue = queue
	return p, nil
}

func (p *Parser) RegistreAdapter(adapters ...Adapter) {
	for _, v := range adapters {
		p.adapters[v.Domain()] = v
	}
}

func (p *Parser) RegistreAdapters(adapters map[string]Adapter) {
	for _, v := range adapters {
		p.adapters[v.Domain()] = v
	}
}

func (p *Parser) Adapter(dom string) Adapter {
	return p.adapters[dom]
}

func (p *Parser) GetAll() []*Article {
	return p.db.GetAll()
}

func (p *Parser) Run() {
	go func() {
		for {
			for _, v := range p.adapters {
				links, err := v.LastArticles()
				if err != nil {
					//todo handle error
					_, fl, ln, _ := runtime.Caller(0)
					logErr(err, fl, ln)
					continue
				}
				for _, v := range links {
					if !p.visited(v) {
						p.queue.SendStringGet(v.String())
					}
				}
			}
			time.Sleep(p.interval)
		}
	}()
}

func (p *Parser) Wait() {
	p.queue.Block()
}

func (p *Parser) ParseUrl(uri string) (*Article, error) {
	u, err := url.Parse(uri)
	if err != nil {
		_, fl, ln, _ := runtime.Caller(0)
		logErr(err, fl, ln)
		return nil, err
	}
	var (
		adapter Adapter
		ok      bool
	)
	if adapter, ok = p.adapters[u.Host]; ok {

	} else {
		adapter = NewDefaultParser(u.Host)
	}
	log.Println(adapter.Name())
	log.Println(adapter.Domain())
	if adapter.IsArticle(u.Path) {
		bts, err := dry.FileGetBytes(u.String())
		if err != nil {
			_, fl, ln, _ := runtime.Caller(0)
			logErr(err, fl, ln)
			return nil, err
		}
		return adapter.Parse(bts)
	}

	//todo const error
	_, fl, ln, _ := runtime.Caller(0)
	logErr(errors.New("Some error"), fl, ln)
	return nil, errors.New("Some error")
}

func (p *Parser) visited(u *url.URL) bool {
	return p.db.Visited(u.String())
}

func (p *Parser) handler(ctx *fetchbot.Context, res *http.Response, err error) {
	if p.db.Visited(ctx.Cmd.URL().String()) {
		return
	}
	if err != nil && res == nil {
		//todo handle error
		_, fl, ln, _ := runtime.Caller(0)
		logErr(err, fl, ln)
		return
	}
	err = p.db.Visit(ctx.Cmd.URL().String())
	adapter, ok := p.adapters[ctx.Cmd.URL().Host]
	if !ok {
		//todo handle error
		_, fl, ln, _ := runtime.Caller(0)
		logErr(err, fl, ln)
		return
	}
	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//todo handle error
		_, fl, ln, _ := runtime.Caller(0)
		logErr(err, fl, ln)
		return
	}
	res.Body.Close()
	art, err := adapter.Parse(bts)
	if err != nil {
		//todo handle error
		_, fl, ln, _ := runtime.Caller(0)
		logErr(err, fl, ln)
		return
	}
	art.Source = ctx.Cmd.URL().String()
	p.db.Save(art)
	greenAdapterName := color.GreenString("[%s]", adapter.Name())
	log.Println(greenAdapterName, ctx.Cmd.URL().String())
	//todo check normal article
	//p.articles <- art
}

func logErr(e error, file string, line int) {
	log.Println(color.RedString("[%s] %s:%d", e, file, line))
}
