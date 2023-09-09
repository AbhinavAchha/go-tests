package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("items.%d.", i)
	}
}

func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "items." + strconv.Itoa(i) + "."
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkSprintf-6       4645328               325.3 ns/op            24 B/op          1 allocs/op
BenchmarkStrconv-6       8856484               130.9 ns/op             7 B/op          0 allocs/op
PASS
ok      ok      3.078s
*/
