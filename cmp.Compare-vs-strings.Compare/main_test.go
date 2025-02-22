package main

import (
	"cmp"
	"math/rand"
	"strings"
	"testing"
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const n = 1000

func getRandom(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}

func BenchmarkCmp(b *testing.B) {
	s := getRandom(rand.Intn(n))
	s2 := getRandom(rand.Intn(n))
	for i := 0; i < b.N; i++ {
		cmp.Compare(s, s2)
	}
}

func BenchmarkStrings(b *testing.B) {
	s := getRandom(rand.Intn(n))
	s2 := getRandom(rand.Intn(n))
	for i := 0; i < b.N; i++ {
		strings.Compare(s, s2)
	}
}

func BenchmarkEqual(b *testing.B) {
	s := getRandom(rand.Intn(n))
	s2 := getRandom(rand.Intn(n))
	for i := 0; i < b.N; i++ {
		_ = s == s2
	}
}
