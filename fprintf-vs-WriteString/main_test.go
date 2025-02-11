package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

var longString = strings.Repeat("0123456789", 100)

var bytes = []byte(longString)

func BenchmarkFpintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		buf.Grow(len(longString))
		fmt.Fprintf(&buf, longString)
	}
}

func BenchmarkIOWriteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		buf.Grow(len(longString))
		io.WriteString(&buf, longString)
	}
}

func BenchmarkWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		buf.Grow(len(longString))
		buf.Write(bytes)
	}
}

func BenchmarkWriteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		buf.Grow(len(longString))
		buf.WriteString(longString)
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkFpintf-6                1000000              1622 ns/op            1056 B/op          2 allocs/op
BenchmarkIOWriteString-6         2551422               471.5 ns/op          1056 B/op          2 allocs/op
BenchmarkWrite-6                 3160246               382.7 ns/op          1024 B/op          1 allocs/op
BenchmarkWriteString-6           3075967               383.8 ns/op          1024 B/op          1 allocs/op
PASS
ok      ok      6.500s
*/
