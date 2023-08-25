package main

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkErrors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = errors.New("test")
	}
}

func BenchmarkFmtErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Errorf("test")
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkErrors-6       1000000000               0.4427 ns/op          0 B/op          0 allocs/op
BenchmarkFmtErrorf-6     5839056               199.2 ns/op            20 B/op          2 allocs/op
PASS
ok      ok      1.869s
*/
