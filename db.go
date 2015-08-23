package articler

type DB interface {
	//for crawlers
	Visited(string) bool
	Visit(string) error

	//for article
	Save(*Article) error
	Get(url string) (*Article, error)
	GetAll() []*Article
}

type MemoryDb struct {
	urls map[string]struct{}
}

func (db *MemoryDb) Visited(u string) bool {
	if _, ok := db.urls[u]; ok {
		return true
	}
	return false
}
func (db *MemoryDb) Visit(u string) error {
	if db.urls == nil {
		db.urls = map[string]struct{}{}
	}
	db.urls[u] = struct{}{}
	return nil
}
