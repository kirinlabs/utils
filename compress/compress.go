package compress

import (
	"bytes"
	"compress/zlib"
	"io"
)

func ZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

func ZlibUncompress(src []byte) []byte {
	var out bytes.Buffer
	b := bytes.NewReader(src)
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}