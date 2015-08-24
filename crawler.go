package articler

import (
	"net/http"
)

type Crawler interface {
	Fetch(string) (*http.Response, error)
}
