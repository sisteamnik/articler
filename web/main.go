package main

import (
	"github.com/sisteamnik/articler"
	"github.com/sisteamnik/articler/adapters"
)

func main() {
	db, err := NewDb("db")
	if err != nil {
		panic(err)
	}

	api, _ := articler.NewParser(db)
	api.RegistreAdapters(adapters.AllAdapters())
	api.RegistreAdapters(ArticlerAdapters)
	api.Run()

	s := NewServer(api)
	s.Start()
}
