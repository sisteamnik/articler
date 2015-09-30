package urltrack

import (
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"time"
)

type Service struct {
	rules         []*Rule
	Client        *http.Client
	CheckInterval time.Duration

	jobs chan Rule
}

func New(rules []Rule) (*Service, error) {
	s := new(Service)
	rulesPoint := []*Rule{}
	for _, v := range rules {
		rulesPoint = append(rulesPoint, s.NewRule(v))
	}
	s.rules = rulesPoint
	s.CheckInterval = 5 * time.Second

	s.jobs = make(chan Rule, 10)
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
