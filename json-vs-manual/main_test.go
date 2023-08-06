package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

var data = `"abcd@gmail.com"`

func BenchmarkUnmarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var email string
		if err := json.Unmarshal([]byte(data), &email); err != nil {
			panic(err)
		}
	}
}

func BenchmarkManualUnmarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if data[0] != '"' || data[len(data)-1] != '"' {
			panic("invalid email")
		}
		email := data[1 : len(data)-1]
		_ = email
	}
}

var email = "abcd@gmail.com"

func BenchmarkMarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(email); err != nil {
			panic(err)
		}
	}
}

func BenchmarkManualMarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := []byte(`"` + email + `"`)
		_ = b
	}
}

func BenchmarkManualMarshalJSONFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := []byte(fmt.Sprintf("%q", email))
		_ = b
	}
}

// go test -bench=. -benchmem
/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkUnmarshalJSON-6                 1634362               715.5 ns/op           192 B/op          4 allocs/op
BenchmarkManualUnmarshalJSON-6          930185868                1.291 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSON-6                   2815305               390.4 ns/op            32 B/op          2 allocs/op
BenchmarkManualMarshalJSON-6            32687058                36.40 ns/op            0 B/op          0 allocs/op
BenchmarkManualMarshalJSONFmt-6          2455448               439.2 ns/op            32 B/op          2 allocs/op
PASS
ok      ok      8.673s
*/
