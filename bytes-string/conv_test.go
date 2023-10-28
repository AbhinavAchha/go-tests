package main

import (
	"bytes"
	"testing"
)

const h = "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."

var by = []byte(h)

func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if string(by) != h {
			b.Fatal("not equal")
		}
	}
}

func BenchmarkBytesCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if bytes.Compare(by, []byte(h)) != 0 {
			b.Fatal("not equal")
		}
	}
}

func BenchmarkWithEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !bytes.Equal(by, []byte(h)) {
			b.Fatal("not equal")
		}
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkStrconv-6              205660273                5.817 ns/op           0 B/op          0 allocs/op
BenchmarkBytesCompare-6         92631925                11.87 ns/op            0 B/op          0 allocs/op
BenchmarkWithEqual-6            100000000               11.93 ns/op            0 B/op          0 allocs/op
PASS
ok      ok      4.112s
*/
