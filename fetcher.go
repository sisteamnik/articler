package articler

import (
	"net/http"
)

type Fetcher interface {
	Get(string) (*http.Response, error)
	Head(string) (*http.Response, error)
}

type DefaultFetcher struct {
}

func (d *DefaultFetcher) Get(u string) (*http.Response, error) {
	return http.Get(u)
}

func (d *DefaultFetcher) Head(u string) (*http.Response, error) {
	return http.Head(u)
}
