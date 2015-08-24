package main

import (
	"github.com/sisteamnik/articler"
	"github.com/sisteamnik/articler/adapters"
	"github.com/ungerik/go-dry"
	"path/filepath"
)

var ArticlerAdapters = map[string]articler.Adapter{}

func init() {
	files, err := dry.ListDirFiles("config")
	if err != nil {
		panic(err)
	}
	for _, v := range files {
		path := filepath.Join("config", v)
		bts, err := dry.FileGetBytes(path)
		if err != nil {
			panic(err)
		}
		ac, err := articler.ParseAdapterConfig(bts)
		if err != nil {
			panic(err)
		}
		sp, err := adapters.NewParser(ac)
		if err != nil {
			panic(err)
		}
		ArticlerAdapters[sp.Domain()] = sp
	}
}
