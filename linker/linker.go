package linker

type Linker interface {
	Visited(string) bool
}
