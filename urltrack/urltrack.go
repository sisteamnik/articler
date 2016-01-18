package urltrack

import (
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Service struct {
	rules         []*Rule
	Client        *http.Client
	CheckInterval time.Duration

	cb CallbackFunc

	jobs chan Rule
}

func New(rules []Rule, interval time.Duration, cb CallbackFunc) (*Service, error) {
	s := new(Service)
	rulesPoint := []*Rule{}
	for _, v := range rules {
		rulesPoint = append(rulesPoint, s.NewRule(v))
	}
	s.cb = cb
	s.rules = rulesPoint
	s.CheckInterval = interval

	s.jobs = make(chan Rule, 10)
	s.Client = http.DefaultClient

	return s, nil
}

func NewWithCommonInterval(urls []string, interval time.Duration,
	cb CallbackFunc) (*Service, error) {
	rules := []Rule{}
	for _, v := range urls {
		if strings.TrimSpace(v) == "" {
			continue
		}
		rules = append(rules, Rule{Url: v, CheckInterval: interval, CallBack: cb})
	}
	return New(rules, interval, cb)
}

func (s *Service) AddUrl(u string) error {
	s.rules = append(s.rules, &Rule{Url: u, CheckInterval: s.CheckInterval, CallBack: s.cb})
	return nil
}

func (s *Service) Run() {
	for {
		for _, v := range s.rules {
			v.Do()
		}
		time.Sleep(s.CheckInterval)
	}
}

func (s *Service) Send(rule Rule) {
	color.Cyan("Send %s", rule)
	s.jobs <- rule
}

func (s *Service) Get(url string) ([]byte, error) {
	color.Cyan("[get] %s", url)
	resp, e := s.Client.Get(url)
	color.Cyan("[here] %s", url)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	bts, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}
	color.Cyan("[getted] %s", url)
	return bts, e
}
