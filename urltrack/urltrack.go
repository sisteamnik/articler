package urltrack

import (
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"time"
)

type Service struct {
	rules  []Rule
	Client *http.Client

	jobs chan Rule
}

func New(rules []Rule) (*Service, error) {
	s := new(Service)
	s.rules = rules
	s.jobs = make(chan Rule)
	s.Client = http.DefaultClient

	return s, nil
}

func NewWithCommonInterval(urls []string, interval time.Duration,
	cb CallbackFunc) (*Service, error) {
	rules := []Rule{}
	for _, v := range urls {
		rules = append(rules, Rule{Url: v, CheckInterval: interval, CallBack: cb})
	}
	return New(rules)
}

func (s *Service) Run() {
	go func() {
		for {
			select {
			case rule := <-s.jobs:
				bts, e := s.Get(rule.Url)
				if e != nil {
					color.Red("Error %s with url %s", e, rule.Url)
				} else {
					rule.CallBack(rule.Url, bts)
					go func(jobs chan Rule, rule Rule) {
						time.Sleep(rule.CheckInterval)
						jobs <- rule
					}(s.jobs, rule)
				}
			}
		}
	}()
	for _, v := range s.rules {
		s.jobs <- v
	}
	for {

	}
}

func (s *Service) Get(url string) ([]byte, error) {
	resp, e := s.Client.Get(url)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	bts, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}
	return bts, e
}
