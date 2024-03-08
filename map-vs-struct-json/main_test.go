package mapvsstructjson

import (
	"encoding/json"
	"testing"
)

var n struct {
	A int
	B int
}

var N = map[string]int{
	"A": 0,
	"B": 0,
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(N); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(n); err != nil {
			b.Error(err)
		}
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkMap-6            899119              1285 ns/op             144 B/op          6 allocs/op
BenchmarkStruct-6        2672896               461.6 ns/op            32 B/op          2 allocs/op
PASS
ok      ok      3.861s
*/
