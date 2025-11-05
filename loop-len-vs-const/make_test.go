package main

import "testing"

var items = make([]string, 1000)

func BenchmarkRange(b *testing.B) {
	for b.Loop() {
		l := len(items)
		for j := range l {
			_ = items[j]
		}
	}
}

func BenchmarkConst(b *testing.B) {
	for b.Loop() {
		l := len(items)
		for j := 0; j < l; j++ {
			_ = items[j]
		}
	}
}

func BenchmarkLenRuntime(b *testing.B) {
	for b.Loop() {
		for j := 0; j < len(items); j++ {
			_ = items[j]
		}
	}
}

func init() {
	for i := range items {
		items[i] = "a"
	}
}
