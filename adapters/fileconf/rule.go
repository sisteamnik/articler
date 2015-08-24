package fileconf

import (
	"fmt"
	"github.com/ungerik/go-dry"
	"reflect"
	"strings"
)

type Rule struct {
	Host string

	TitleSelector     string
	TextSelector      string
	PublishedSelector string
}

func getRulesFromFile(filename string) (res []*Rule, e error) {
	var (
		lines []string
	)

	lines, e = dry.FileGetLines(filename)
	if e != nil {
		return
	}
	for _, v := range lines {
		rule, e := parseRule(v)
		if e != nil {
			continue
		}
		res = append(res, rule)
	}
	return
}

var (
	ErrCommentLine error = fmt.Errorf("line is comment")
	ErrEmptyLine         = fmt.Errorf("emty line")
	ErrEmptyRule         = fmt.Errorf("emty rule")
)

func parseRule(line string) (*Rule, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, ErrEmptyLine
	}
	if line[0] == "#"[0] {
		return nil, ErrCommentLine
	}

	arr := strings.SplitN(line, " ", 2)
	if len(arr) != 2 {
		return nil, ErrEmptyRule
	}

	rule := Rule{
		Host: arr[0],
	}

	st := reflect.StructTag(arr[1])
	rule.TextSelector = st.Get("text")
	rule.TitleSelector = st.Get("title")
	rule.PublishedSelector = st.Get("publ")
	return &rule, nil
}
