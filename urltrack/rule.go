package urltrack

import (
	"github.com/fatih/color"
	"time"
)

type CallbackFunc func(string, []byte)

type Rule struct {
	Url           string
	CheckInterval time.Duration
	CallBack      CallbackFunc

	srv *Service
}

func (s *Service) NewRule(rule Rule) *Rule {
	r := &rule
	r.srv = s
	return r
}

func (rule *Rule) Run() {
	color.Cyan("Start running %s", rule.Url)
	ticker := time.Tick(rule.CheckInterval)
	for _ = range ticker {
		color.Cyan("Tick %s", rule.Url)
		rule.Do()
	}
}

func (rule *Rule) Do() {
	bts, e := rule.srv.Get(rule.Url)
	if e != nil {
		color.Red("Error %s with url %s", e, rule.Url)
	} else {
		color.Cyan("[callback] %s", rule.Url)
		rule.CallBack(rule.Url, bts)
	}
}
