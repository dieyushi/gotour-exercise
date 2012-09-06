/*
Exercise: Rot13 Reader

A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of gzipped data) and returns 
a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream 
by applying the ROT13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = (b - 'A') + 'N'
	case 'N' <= b && b <= 'Z':
		b = (b - 'N') + 'A'
	case 'a' <= b && b <= 'm':
		b = (b - 'a') + 'n'
	case 'n' <= b && b <= 'z':
		b = (b - 'n') + 'a'
	}
	return b
}

func (b *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = b.r.Read(p)
	for i := range p[:n] {
		p[i] = rot13(p[i])
	}
	return
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
