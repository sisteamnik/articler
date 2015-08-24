package crawler

import (
	"bytes"
	"compress/gzip"
)

func gz(in []byte) (out []byte, e error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, e = w.Write(in)
	if e != nil {
		return
	}
	e = w.Close()
	if e != nil {
		return
	}
	return b.Bytes(), nil
}
