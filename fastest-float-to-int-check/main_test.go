package main

import (
	"math"
	"testing"
)

func isIntegerUsingMod(f float64) bool {
	return math.Mod(f, 1) == 0
}

func isIntegerUsingCast(f float64) bool {
	return f == float64(int(f))
}

func BenchmarkIsIntegerUsingMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isIntegerUsingMod(12345.0)
	}
}

func BenchmarkIsIntegerUsingCast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isIntegerUsingCast(12345.0)
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkIsIntegerUsingMod-6            17660739                65.66 ns/op            0 B/op          0 allocs/op
BenchmarkIsIntegerUsingCast-6           1000000000               0.2554 ns/op          0 B/op          0 allocs/op
PASS
ok      ok      1.530s
*/
