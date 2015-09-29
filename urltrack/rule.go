package urltrack

import (
	"time"
)

type CallbackFunc func(string, []byte)

type Rule struct {
	Url           string
	CheckInterval time.Duration
	CallBack      CallbackFunc
}
