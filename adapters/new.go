package adapters

import (
	"github.com/sisteamnik/articler"
)

var adapters = map[string]articler.Adapter{
	"smoldaily": &SmoldailyParser{},
	"smolenski": NewSmolenskIParser(),
	"readovka":  NewReadovkaParser(),
	"fas":       NewFasParser(),
	"admin":     NewAdminParser(),
	"mvd":       NewMvdParser(),
}

func New(name string) articler.Adapter {
	if a, ok := adapters[name]; ok {
		return a
	}
	return &articler.DefaultParser{}
}

func AllAdapters() map[string]articler.Adapter {
	return adapters
}
