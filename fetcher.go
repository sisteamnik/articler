package articler

import (
	"errors"
	"io"
	"net/http"
)

var (
	MaxBodySize int64 = 1024 * 1024
)

type Fetcher interface {
	Get(string) (*http.Response, error)
	Head(string) (*http.Response, error)
}

type DefaultFetcher struct {
}

func (d *DefaultFetcher) Get(u string) (*http.Response, error) {
	r, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	r.Body = NewLimitedReadCloser(r.Body, MaxBodySize)
	return r, nil
}

func (d *DefaultFetcher) Head(u string) (*http.Response, error) {
	return http.Head(u)
}

type LimitedReadCloser struct {
	io.ReadCloser
	N int64
}

func NewLimitedReadCloser(rc io.ReadCloser, l int64) *LimitedReadCloser {
	return &LimitedReadCloser{rc, l}
}

func (l *LimitedReadCloser) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, errors.New("http: response body too large")
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.ReadCloser.Read(p)
	l.N -= int64(n)
	return
}
