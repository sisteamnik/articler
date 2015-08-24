package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/fetchbot"
	"github.com/Unknwon/com"
	"github.com/fatih/color"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func (cr *Crawler) Handler(ctx *fetchbot.Context, res *http.Response, e error) {
	if e != nil {
		log.Println(e)
		return
	}
	color.Cyan("[%d] %s,", res.StatusCode, ctx.Cmd.URL())
	if res.ContentLength > 10*1000*1000 {
		log.Println("File to large")
		return
	}

	crTime := time.Duration(0)

	durString := res.Header.Get(CrawlerRequestTimeHeader)
	if durString != "" {
		crTime = time.Duration(com.StrTo(durString).MustInt64())
		res.Header.Del(CrawlerRequestTimeHeader)
	}

	bts, e := httputil.DumpResponse(res, true)
	if e != nil {
		log.Println(e)
		return
	}

	l, has := cr.store.GetLink(ctx.Cmd.URL().RequestURI())
	if !has {
		l = NewLink(ctx.Cmd.URL())
	}

	l.CrawlTime = crTime
	l.Visited = time.Now()

	e = cr.store.SaveBlob(l.URL().String(), bts)
	if e != nil {
		log.Println(e)
		return
	}

	e = cr.store.SaveLink(l)
	if e != nil {
		log.Println(e)
		return
	}

	arr := strings.SplitN(string(bts), "\n\n", 2)
	if len(arr) < 2 {
		log.Println("body not found", ctx.Cmd.URL())
		return
	}

	urls := ExtractLinks(ctx.Cmd.URL(), []byte(arr[1]))
	for _, v := range urls {
		link, has := cr.store.GetLink(v.RequestURI())
		if has {
			continue
		}
		link = NewLink(v)
		e := cr.store.SaveLink(link)
		if e != nil {
			log.Println(e)
			return
		}

		clearUrl(v)
		if cr.NeedVisit(v) {
			cr.dupMu.Lock()
			cr.dup[v.String()] = struct{}{}
			cr.dupMu.Unlock()
			ctx.Q.SendStringGet(v.String())
		}

	}
}

func (cr *Crawler) NeedVisit(u *url.URL) bool {
	if _, ok := cr.dup[u.String()]; ok {
		return false
	}
	if _, ok := cr.allowedSites[u.Host]; !ok {
		return false
	}
	return true
}

func (cr *Crawler) ResponseMatch(res *http.Response) bool {
	return false
}

type ContextCmd struct {
}

func NewContextCmd(string) *ContextCmd {
	return nil
}

func (c ContextCmd) URL() *url.URL {
	return nil
}

func (c ContextCmd) Method() string {
	return ""
}

type Doer struct {
	*http.Client
}

var CrawlerRequestTimeHeader = "crawler-request-time"

func (d *Doer) Do(req *http.Request) (*http.Response, error) {
	if d.Client == nil {
		d.Client = http.DefaultClient
		d.Client.Timeout = 25 * time.Second
	}
	start := time.Now()
	resp, e := d.Client.Do(req)
	since := time.Since(start)
	if e != nil {
		return resp, e
	}
	resp.Header.Set(CrawlerRequestTimeHeader, fmt.Sprint(since.Nanoseconds()))
	return resp, e
}
