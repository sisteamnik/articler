package crawler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Crawler struct {
	store    Store
	fetchbot *fetchbot.Fetcher

	client fetchbot.Doer

	allowedSites map[string]struct{}
	dupMu        sync.Mutex
	dup          map[string]struct{}

	q *fetchbot.Queue
}

func New(cnf Config) (*Crawler, error) {
	store, e := Open(cnf.Dbdriver, cnf.Dbcnf)
	if e != nil {
		return nil, e
	}

	cr := &Crawler{
		store:        store,
		client:       &Doer{},
		dup:          map[string]struct{}{},
		allowedSites: map[string]struct{}{},
	}

	mux := fetchbot.NewMux()

	mux.HandleErrors(fetchbot.HandlerFunc(func(ctx *fetchbot.Context, res *http.Response, err error) {
		color.Red("[ERR] %s %s - %s\n", ctx.Cmd.Method(), ctx.Cmd.URL(), err)
	}))

	mux.Response().Status(200).ContentType("text/html").Handler(fetchbot.HandlerFunc(cr.Handler))

	cr.fetchbot = fetchbot.New(mux)
	cr.fetchbot.HttpClient = cr.client
	cr.fetchbot.DisablePoliteness = true
	cr.fetchbot.CrawlDelay = time.Second * 1
	cr.fetchbot.UserAgent = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"
	return cr, nil
}

func (cr *Crawler) Run(links ...string) {
	color.Green("Run crawler")
	vi, _ := cr.store.Visited()
	for _, v := range vi {
		cr.dup[v.URL().String()] = struct{}{}
	}
	color.Green("Current link count: %d(visited: %d)", cr.store.CountLink(), len(vi))
	for _, v := range links {
		u, e := url.Parse(v)
		if e != nil {
			continue
		}
		site, has := cr.store.GetSite(u.Host)
		if !has {
			site = NewSite(u.Host, true)
		} else {
			site.Allowed = true
		}
		e = cr.store.SaveSite(site)
		if e != nil {
			color.Red(e.Error())
		}
	}
	as, _ := cr.store.AllowedSites()
	color.Green("Current site count: %d, allowed %d", cr.store.CountSite(), len(as))
	for _, v := range as {
		if cr.allowedSites == nil {
			cr.allowedSites = map[string]struct{}{}
		}
		cr.allowedSites[v.Host] = struct{}{}
	}
	q := cr.fetchbot.Start()
	for _, v := range links {
		cr.dup[v] = struct{}{}
	}
	_, e := q.SendStringGet(links...)
	if e != nil {
		color.Red(e.Error())
	}
	cr.q = q
	lnks, e := cr.store.NotVisited()
	if e != nil {
		color.Red(e.Error())
	}

	for _, v := range lnks {
		l := v.URL()
		clearUrl(l)
		if cr.NeedVisit(l) {
			q.SendStringGet(l.String())
		}
	}
}
func (cr *Crawler) Stop() {
	cr.q.Cancel()
}

func (cr *Crawler) Block() {
	if cr.q != nil {
		cr.q.Block()
	}
}

func ExtractLinks(url *url.URL, in []byte) (res []*url.URL) {
	doc, e := goquery.NewDocumentFromReader(bytes.NewReader(in))
	if e != nil {
		return
	}

	dup := map[string]struct{}{}
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("href")
		// Resolve address
		u, err := url.Parse(val)
		if err != nil {
			return
		}
		clearUrl(u)
		if err != nil {
			fmt.Printf("error: resolve URL %s - %s\n", val, err)
			return
		}
		if _, ok := dup[u.String()]; !ok {
			res = append(res, u)
		} else {
			dup[u.String()] = struct{}{}
		}
	})
	return
}

func clearUrl(u *url.URL) {
	u.RawQuery = ""
	u.Fragment = ""
}

func (cr *Crawler) Get(rawurl string, forceFetch bool) (link *Link, blob []byte, e error) {
	var (
		u    *url.URL
		has  bool
		req  *http.Request
		resp *http.Response
		bts  []byte
	)

	u, e = url.Parse(rawurl)
	if e != nil {
		return
	}

	link, has = cr.store.GetLink(u.RequestURI())
	if has && !forceFetch && !link.Visited.IsZero() {
		blob, e = link.GetBlob(cr.store)
		color.Cyan("From db")
		return
	}
	req, e = http.NewRequest("GET", rawurl, nil)
	if e != nil {
		return
	}
	resp, e = cr.fetchbot.HttpClient.Do(req)
	if e != nil {
		return
	}
	defer resp.Body.Close()

	blob, e = httputil.DumpResponse(resp, true)
	if e != nil {
		return
	}

	color.Yellow("%d", len(bts))

	link = NewLink(u)
	cr.store.SaveBlob(link.URL().String(), blob)

	link.Visited = time.Now()
	e = cr.store.SaveLink(link)
	return
}

var ErrNotFound = fmt.Errorf("not found")
