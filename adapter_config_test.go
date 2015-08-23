package articler

import (
	"github.com/jinzhu/now"
	"log"
	"testing"
	"time"
)

func TestParseAdapterConfig(t *testing.T) {
	ac, err := ParseAdapterConfig([]byte(aconf))
	if err != nil {
		t.Error(err)
	}
	log.Println(ac)
}

func TestNowParse(t *testing.T) {
	now.TimeFormats = append(now.TimeFormats, time.RFC3339)
	s := time.Now().Format(time.RFC3339)
	d, err := now.Parse(s)
	if err != nil {
		t.Error(err)
	}
	log.Println(d)
}

var aconf = `
host: smolensk-i.ru
name: smolenski

feedtype: html
articleuriregex: "^/[a-z-_0-9]*/[a-z-_0-9]*$"
feedselector: ".contentColumn h1 a"

bodyselector: .entry-content
`
