package main

import (
	"maps"
	"math/rand"
	"slices"
	"testing"
)

const size = 100000

var data = make(map[int]struct{}, size)

func init() {
	for range size {
		data[rand.Intn(size)] = struct{}{}
	}
}

func BenchmarkPrimitive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		values := make([]int, len(data))
		i := 0
		for k := range data {
			values[i] = k
			i++
		}
	}
}

func BenchmarkIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		values := maps.Keys(data)
		v := make([]int, len(data))
		slices.AppendSeq(v, values)
	}
}
