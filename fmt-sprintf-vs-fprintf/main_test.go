package main

import (
	"fmt"
	"strings"
	"testing"
)

var longString = strings.Repeat("%s", 1000)
var c = make([]any, 1000)

func init() {
	for i := 0; i < 1000; i++ {
		c[i] = strings.Repeat("a", 1000)
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b strings.Builder
		b.WriteString(fmt.Sprintf(longString, c...))
	}
}

func BenchmarkFprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := new(strings.Builder)
		fmt.Fprintf(b, longString, c...)
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkSprintf-6        344295              3786 ns/op            4097 B/op          2 allocs/op
BenchmarkFprintf-6        344166              3228 ns/op            2080 B/op          2 allocs/op
PASS
ok      ok      2.494s
*/
