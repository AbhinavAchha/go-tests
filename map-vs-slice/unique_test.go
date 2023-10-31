package main_test

import (
	"math/rand"
	"slices"
	"testing"
)

const l = 410

var nonunique = make([]int, l)

func init() {
	for i := range nonunique {
		nonunique[i] = rand.Intn(l)
	}
}

func BenchmarkSliceUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unique := make([]int, 0, l)
		for _, v := range nonunique {
			if !slices.Contains(unique, v) {
				unique = append(unique, v)
			}
		}

	}
}

func BenchmarkMapUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unique := make(map[int]struct{}, l)
		for _, v := range nonunique {
			unique[v] = struct{}{}
		}
	}
}

/*
go test -bench=Unique -benchmem
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkSliceUnique-6             46687             26130 ns/op               0 B/op          0 allocs/op
BenchmarkMapUnique-6               44708             27295 ns/op           10904 B/op          2 allocs/op
PASS
ok      ok      2.982s


Any value before ~400 is faster to use a slice, after that a map is faster.
*/
