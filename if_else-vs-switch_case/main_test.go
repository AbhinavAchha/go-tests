package main

import (
	"math/rand"
	"testing"
)

var a = rand.Int31() % 2147483647

func init() {
	println("a:", a)
}

func BenchmarkIfElese(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IfElese(a)
	}
}

func BenchmarkSwitchCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SwitchCase(a)
	}
}

func IfElese(a int32) {
	if a == 0 {
		return
	}

	if a < 0 {
		return
	}

	if a < 100 {
		return
	}

	if a < 1000 {
		return
	}

	if a < 10000 {
		return
	}

	if a < 100000 {
		return
	}

	if a < 1000000 {
		return
	}

	if a < 10000000 {
		return
	}

	if a < 100000000 {
		return
	}

	if a < 1000000000 {
		return
	}

	if a < 2147483647 {
		return
	}

}

func SwitchCase(a int32) {
	switch {
	case a == 0:
		return
	case a < 10:
		return
	case a < 100:
		return
	case a < 1000:
		return
	case a < 10000:
		return
	case a < 100000:
		return
	case a < 1000000:
		return
	case a < 10000000:
		return
	case a < 100000000:
		return
	case a < 1000000000:
		return
	case a < 2147483647:
		return
	}
}
