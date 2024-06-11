package main

import (
	"bytes"
	"testing"
)

var randomStr string

var randomBytes = []byte(randomStr)

func BenchmarkStrringConv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randomStr == string(randomBytes)
	}
}

func BenchmarkBytesEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bytes.Equal(randomBytes, []byte(randomStr))
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = loop(randomBytes, randomStr)
	}
}

func loop(b []byte, s string) bool {
	for i := range b {
		if b[i] != s[i] {
			return false
		}
	}
	return true
}

func init() {
	ramdomStr := make([]byte, 1000)
	for i := range ramdomStr {
		ramdomStr[i] = byte(i)
	}
	randomStr = string(ramdomStr)
}
