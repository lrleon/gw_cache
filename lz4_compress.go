package gw_cache

import (
	"bytes"
	"io"

	"github.com/pierrec/lz4"
)

type lz4Compressor struct{}

func (_ lz4Compressor) Compress(in []byte) ([]byte, error) {
	r := bytes.NewReader(in)
	w := &bytes.Buffer{}
	zw := lz4.NewWriter(w)
	_, err := io.Copy(zw, r)
	if err != nil {
		return nil, err
	}
	// Closing is *very* important
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (_ lz4Compressor) Decompress(in []byte) ([]byte, error) {
	r := bytes.NewReader(in)
	w := &bytes.Buffer{}
	zr := lz4.NewReader(r)
	_, err := io.Copy(w, zr)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
