package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (s *rot13Reader) Read(b []byte) (int, error) {
	n, err := s.r.Read(b)
	for i, v := range b {
		switch {
		case v >= 'a' && v <= 'z':
			b[i] = (v-'a'+13)%26 + 'a'
		case v >= 'A' && v <= 'Z':
			b[i] = (v-'A'+13)%26 + 'A'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
