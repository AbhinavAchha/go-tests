package main

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
)

func BenchmarkBigRatAdd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Rat).SetFrac(big.NewInt(1), big.NewInt(3))
		b := new(big.Rat).SetFrac(big.NewInt(2), big.NewInt(5))
		a.Add(a, b)
	}
}

func BenchmarkBigFloatAdd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Float).SetFloat64(1.0 / 3.0)
		b := new(big.Float).SetFloat64(2.0 / 5.0)
		a.Add(a, b)
	}
}

func BenchmarkDecimalAdd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := decimal.NewFromFloat(1.0 / 3.0)
		b := decimal.NewFromFloat(2.0 / 5.0)
		a = a.Add(b)
	}
}

func BenchmarkBigRatMul(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Rat).SetFrac(big.NewInt(1), big.NewInt(3))
		b := new(big.Rat).SetFrac(big.NewInt(2), big.NewInt(5))
		a.Mul(a, b)
	}
}

func BenchmarkBigFloatMul(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Float).SetFloat64(1.0 / 3.0)
		b := new(big.Float).SetFloat64(2.0 / 5.0)
		a.Mul(a, b)
	}
}

func BenchmarkDecimalMul(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := decimal.NewFromFloat(1.0 / 3.0)
		b := decimal.NewFromFloat(2.0 / 5.0)
		a = a.Mul(b)
	}
}

func BenchmarkBigRatSub(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Rat).SetFrac(big.NewInt(1), big.NewInt(3))
		b := new(big.Rat).SetFrac(big.NewInt(2), big.NewInt(5))
		a.Sub(a, b)
	}
}

func BenchmarkBigFloatSub(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Float).SetFloat64(1.0 / 3.0)
		b := new(big.Float).SetFloat64(2.0 / 5.0)
		a.Sub(a, b)
	}
}

func BenchmarkDecimalSub(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := decimal.NewFromFloat(1.0 / 3.0)
		b := decimal.NewFromFloat(2.0 / 5.0)
		a = a.Sub(b)
	}
}

func BenchmarkBigRatDiv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Rat).SetFrac(big.NewInt(1), big.NewInt(3))
		b := new(big.Rat).SetFrac(big.NewInt(2), big.NewInt(5))
		a.Quo(a, b)
	}
}

func BenchmarkBigFloatDiv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Float).SetFloat64(1.0 / 3.0)
		b := new(big.Float).SetFloat64(2.0 / 5.0)
		a.Quo(a, b)
	}
}

func BenchmarkDecimalDiv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := decimal.NewFromFloat(1.0 / 3.0)
		b := decimal.NewFromFloat(2.0 / 5.0)
		a = a.Div(b)
	}
}

func BenchmarkBigRatString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Rat).SetFrac(big.NewInt(1), big.NewInt(3))
		a.String()
	}
}

func BenchmarkBigFloatString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := new(big.Float).SetFloat64(1.0 / 3.0)
		a.Text('f', 10)
	}
}

func BenchmarkDecimalString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := decimal.NewFromFloat(1.0 / 3.0)
		a.String()
	}
}

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkBigRatAdd-6             2314854               529.9 ns/op           272 B/op         14 allocs/op
BenchmarkBigFloatAdd-6           8381310               136.1 ns/op            64 B/op          3 allocs/op
BenchmarkDecimalAdd-6             790454              1468 ns/op             352 B/op         12 allocs/op
BenchmarkBigRatMul-6             2882684               421.7 ns/op           176 B/op         12 allocs/op
BenchmarkBigFloatMul-6           9600714               125.2 ns/op            64 B/op          3 allocs/op
BenchmarkDecimalMul-6             944631              1151 ns/op             160 B/op          6 allocs/op
BenchmarkBigRatSub-6             2402082               494.0 ns/op           224 B/op         13 allocs/op
BenchmarkBigFloatSub-6          13155297                88.23 ns/op           16 B/op          2 allocs/op
BenchmarkDecimalSub-6             758347              1469 ns/op             312 B/op         12 allocs/op
BenchmarkBigRatDiv-6             2785820               429.8 ns/op           176 B/op         12 allocs/op
BenchmarkBigFloatDiv-6           6034087               199.2 ns/op            88 B/op          5 allocs/op
BenchmarkDecimalDiv-6             742180              1578 ns/op             360 B/op         15 allocs/op
BenchmarkBigRatString-6          3703168               324.9 ns/op            56 B/op          7 allocs/op
BenchmarkBigFloatString-6        2504122               479.9 ns/op           168 B/op          7 allocs/op
BenchmarkDecimalString-6         1437510               827.7 ns/op           104 B/op          6 allocs/op
PASS
ok      ok      21.871s
*/
