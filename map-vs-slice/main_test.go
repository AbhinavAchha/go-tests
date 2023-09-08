package main

import (
	"math/rand"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maps()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}

func maps() {
	m := make(map[int]struct{}, 100000)
	for i := 0; i < 100000; i++ {
		m[i] = struct{}{}
	}

	for i := 0; i < 100; i++ {
		r := rand.Intn(100000)
		_ = m[r]
	}

	for i := 0; i < 100; i++ {
		r := rand.Intn(100000)
		m[r] = struct{}{}
	}
}

func slice() {
	s := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		s[i] = i
	}

	for i := 0; i < 100; i++ {
		r := rand.Intn(100000)
		_ = s[r]
	}

	for i := 0; i < 100; i++ {
		r := rand.Intn(100000)
		s[r] = r
	}
}
