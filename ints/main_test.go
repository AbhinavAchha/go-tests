package main

import (
	"math/big"
	"testing"
)

var n uint = 127

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := int(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := int8(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := int16(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := int32(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := int64(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := uint(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := uint8(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := uint16(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := uint32(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := uint64(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := float32(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := float64(n)
		for i > 0 {
			i--
		}
	}
}

func BenchmarkBigInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := big.NewInt(int64(n))
		for i.Sign() > 0 {
			i.Sub(i, big.NewInt(1))
		}
	}
}

func BenchmarkBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := big.NewFloat(float64(n))
		for i.Sign() > 0 {
			i.Sub(i, big.NewFloat(1))
		}
	}
}

func BenchmarkBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := big.NewRat(int64(n), 1)
		for i.Sign() > 0 {
			i.Sub(i, big.NewRat(1, 1))
		}
	}
}

func BenchmarkBigFloatRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := big.NewFloat(0)
		i.SetInt(big.NewInt(int64(n)))
		for i.Sign() > 0 {
			i.Sub(i, big.NewFloat(1))
		}
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkInt-6                  16139017                70.45 ns/op            0 B/op          0 allocs/op
BenchmarkInt8-6                 16772132                72.87 ns/op            0 B/op          0 allocs/op
BenchmarkInt16-6                16097419                73.37 ns/op            0 B/op          0 allocs/op
BenchmarkInt32-6                17052178                69.89 ns/op            0 B/op          0 allocs/op
BenchmarkInt64-6                17172228                72.74 ns/op            0 B/op          0 allocs/op
BenchmarkUint-6                 17016302                69.93 ns/op            0 B/op          0 allocs/op
BenchmarkUint8-6                17184645                72.83 ns/op            0 B/op          0 allocs/op
BenchmarkUint16-6               15984944                73.25 ns/op            0 B/op          0 allocs/op
BenchmarkUint32-6               17050960                70.30 ns/op            0 B/op          0 allocs/op
BenchmarkUint64-6               17071629                72.73 ns/op            0 B/op          0 allocs/op
BenchmarkFloat32-6              10922221               111.1 ns/op             0 B/op          0 allocs/op
BenchmarkFloat64-6              10000432               110.5 ns/op             0 B/op          0 allocs/op
BenchmarkBigInt-6                 716134              1746 ns/op               8 B/op          1 allocs/op
BenchmarkBigFloat-6                43551             31809 ns/op            7072 B/op        254 allocs/op
BenchmarkBigRat-6                  14925             83678 ns/op           18352 B/op       1019 allocs/op
BenchmarkBigFloatRat-6             38220             32716 ns/op            7072 B/op        254 allocs/op
PASS
ok      ok      25.122s
*/
