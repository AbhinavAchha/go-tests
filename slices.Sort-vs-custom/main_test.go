package main

import (
	"math/rand"
	"sort"
	"testing"

	"slices"
)

type Test struct {
	Num int
}

var arr []Test

func init() {
	arr = make([]Test, 100)
	for i := 0; i < 100; i++ {
		arr[i].Num = rand.Intn(100)
	}
}

func BenchmarkSlicesSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var cpy = make([]Test, len(arr))
		copy(cpy, arr)
		slices.SortFunc(cpy, func(a, b Test) int {
			return a.Num - b.Num
		})
	}
}

func BenchmarkSlicesSortCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var cpy = make([]Test, len(arr))
		l := copy(cpy, arr)
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				if cpy[i].Num > cpy[j].Num {
					cpy[i], cpy[j] = cpy[j], cpy[i]
				}
			}
		}
	}
}

func BenchmarkSortSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var cpy = make([]Test, len(arr))
		copy(cpy, arr)
		sort.Slice(cpy, func(i, j int) bool {
			return cpy[i].Num < cpy[j].Num
		})
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkSlicesSort-6             216426              5430 ns/op             896 B/op          1 allocs/op
BenchmarkSlicesSortCustom-6       163777              8758 ns/op             896 B/op          1 allocs/op
BenchmarkSortSlice-6              213036              5926 ns/op             952 B/op          3 allocs/op
PASS
ok      ok      5.076s
*/
