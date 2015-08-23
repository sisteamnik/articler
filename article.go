package articler

import (
	"time"
)

type Article struct {
	Title       string
	Description string
	Body        []byte

	Images        []string
	Videos        []string
	Links         []string
	Authors       []string
	Tags          []string
	Catigories    []string
	Photographers []string

	Source string

	Published time.Time
}

type Articles []*Article

func (a Articles) Len() int           { return len(a) }
func (a Articles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Articles) Less(i, j int) bool { return a[i].Published.Sub(a[j].Published) < 0 }
