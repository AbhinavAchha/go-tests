package main

import "testing"

var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func BenchmarkVariadic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Variadic(s...)
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slice(s)
	}
}

func Variadic(args ...int) {
	for i := range args {
		_ = args[i]
	}
}

func Slice(args []int) {
	for i := range args {
		_ = args[i]
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkVariadic-6     100000000               10.11 ns/op            0 B/op          0 allocs/op
BenchmarkSlice-6        100000000               11.16 ns/op            0 B/op          0 allocs/op
PASS
ok      ok      2.159s
*/
