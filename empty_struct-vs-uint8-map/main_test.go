package main

import (
	"testing"
)

func BenchmarkMapStruct(b *testing.B) {
	m := make(map[int]struct{}, 4000)
	for i := 0; i < 4000; i++ {
		m[i] = struct{}{}
	}
}

func BenchmarkMapUint8(b *testing.B) {
	m := make(map[int]uint8, 4000)
	for i := 0; i < 4000; i++ {
		m[i] = 1
	}
}

func BenchmarkMapBoolean(b *testing.B) {
	m := make(map[int]bool, 4000)
	for i := 0; i < 4000; i++ {
		m[i] = true
	}
}
