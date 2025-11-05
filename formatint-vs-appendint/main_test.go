package formatintvsappendint

import (
	"bytes"
	"strconv"
	"testing"
)

const base = 10

func BenchmarkFormatInt(b *testing.B) {
	var s bytes.Buffer
	for i := 0; i < b.N; i++ {
		s.Write([]byte(strconv.FormatInt(int64(b.N), base)))
	}
}

func BenchmarkAppendInt(b *testing.B) {
	var s bytes.Buffer
	for i := 0; i < b.N; i++ {
		s.Write(strconv.AppendInt(nil, int64(b.N), base))
	}
}

func BenchmarkFormatBool(b *testing.B) {
	var s bytes.Buffer
	for i := 0; i < b.N; i++ {
		s.Write([]byte(strconv.FormatBool(b.N%2 == 0)))
	}
}

func BenchmarkAppendBool(b *testing.B) {
	d := make([]byte, 0)
	for i := 0; i < b.N; i++ {
		strconv.AppendBool(d, b.N%2 == 0)
	}
}

/*
go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkFormatInt-6    27980974                42.35 ns/op          27 B/op          1 allocs/op
BenchmarkAppendInt-6    22945772                52.49 ns/op          31 B/op          1 allocs/op
BenchmarkFormatBool-6   147978834               8.943 ns/op          14 B/op          0 allocs/op
BenchmarkAppendBool-6   41776646                26.83 ns/op          20 B/op          1 allocs/op
PASS
ok      ok      3.306s
*/
