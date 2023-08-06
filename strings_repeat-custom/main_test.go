package main

import (
	"strings"
	"testing"
)

// 1st
func BenchmarkStringRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat(s, times)
	}
}

// 6th
func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(make([]string, times), s)
	}
}

// 5th
func BenchmarkStringBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str strings.Builder
		for j := 0; j < times; j++ {
			str.WriteString(s)
		}
	}
}

// 4th
func BenchmarkStringBufferGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str strings.Builder
		str.Grow(times)
		for j := 0; j < times; j++ {
			str.WriteString(s)
		}
	}
}

// 3rd
func BenchmarkStringBufferGrowByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str strings.Builder
		str.Grow(times)
		for j := 0; j < times; j++ {
			str.WriteByte('_')
		}
	}
}

// 4th
func BenchmarkStringBufferGrowRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str strings.Builder
		str.Grow(times)
		for j := 0; j < times; j++ {
			str.WriteRune('_')
		}
	}
}

/*
results
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkStringRepeat-6          7258807               151.8 ns/op           112 B/op          1 allocs/op
BenchmarkStringJoin-6            1000000              1077 ns/op             112 B/op          1 allocs/op
BenchmarkStringBuffer-6          1263193               947.0 ns/op           248 B/op          5 allocs/op
BenchmarkStringBufferGrow-6      1829146               599.7 ns/op           112 B/op          1 allocs/op
BenchmarkStringBufferGrow2-6     4562769               315.4 ns/op           112 B/op          1 allocs/op
BenchmarkStringBufferGrow3-6     1647830               707.9 ns/op           112 B/op          1 allocs/op
PASS
ok      ok      9.846s
*/
