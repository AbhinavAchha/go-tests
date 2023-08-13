package main

import "testing"

const length = 10000000

var slice = make([]int, length)

var array = [length]int{}

func BenchmarkSliceWrite(b *testing.B) {
	for i := 0; i < length; i++ {
		slice[i] = i
	}
}

func BenchmarkArrayWrite(b *testing.B) {
	for i := 0; i < length; i++ {
		array[i] = i
	}
}

func BenchmarkSliceRead(b *testing.B) {
	for i := 0; i < length; i++ {
		_ = slice[i]
	}
}

func BenchmarkArrayRead(b *testing.B) {
	for i := 0; i < length; i++ {
		_ = array[i]
	}
}

func BenchmarkSliceIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range slice {
			_ = v
		}
	}
}

func BenchmarkArrayIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range array {
			_ = v
		}
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkSliceWrite-6           1000000000               0.1547 ns/op          0 B/op          0 allocs/op
BenchmarkArrayWrite-6           1000000000               0.1317 ns/op          0 B/op          0 allocs/op
BenchmarkSliceRead-6            1000000000               0.03823 ns/op         0 B/op          0 allocs/op
BenchmarkArrayRead-6            1000000000               0.06427 ns/op         0 B/op          0 allocs/op
BenchmarkSliceIterate-6               46          26351257 ns/op               0 B/op          0 allocs/op
BenchmarkArrayIterate-6               45          25634027 ns/op               0 B/op          0 allocs/op
PASS
ok      ok      4.243s

*/
