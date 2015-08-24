package crawler

import (
	"sync"
)

type Store interface {
	Connect(string) error

	SaveLink(*Link) error
	GetLink(string) (*Link, bool)
	CountLink() int
	Visited() ([]*Link, error)
	NotVisited() ([]*Link, error)

	SaveSite(*Site) error
	GetSite(string) (*Site, bool)
	CountSite() int
	AllowedSites() ([]*Site, error)

	SaveBlob(string, []byte) error
	GetBlob(string) ([]byte, error)
}

var (
	stores   = map[string]Store{}
	storesMu sync.Mutex
)

func Register(name string, store Store) {
	storesMu.Lock()
	defer storesMu.Unlock()
	stores[name] = store
}

func Open(name string, conf string) (Store, error) {
	if store, ok := stores[name]; !ok {
		panic("store not found")
	} else {
		return store, store.Connect(conf)
	}
}
