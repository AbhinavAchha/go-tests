package main

import (
	"math/big"
	"testing"
)

type MutableQuantity struct {
	bigInt     *big.Int
	int32Field int32
}

func (q *MutableQuantity) Add(q2 MutableQuantity) *MutableQuantity {
	q.bigInt = new(big.Int).Add(q.bigInt, q2.bigInt)
	q.int32Field += q2.int32Field
	return q
}

func (q *MutableQuantity) Sub(q2 MutableQuantity) *MutableQuantity {
	q.bigInt = new(big.Int).Sub(q.bigInt, q2.bigInt)
	q.int32Field -= q2.int32Field
	return q
}

func (q *MutableQuantity) Mul(q2 MutableQuantity) *MutableQuantity {
	q.bigInt = new(big.Int).Mul(q.bigInt, q2.bigInt)
	q.int32Field *= q2.int32Field
	return q
}

func (q *MutableQuantity) Div(q2 MutableQuantity) *MutableQuantity {
	q.bigInt = new(big.Int).Div(q.bigInt, q2.bigInt)
	q.int32Field /= q2.int32Field
	return q
}

type ImmutableQuantity struct {
	bigInt     *big.Int
	int32Field int32
}

func NewImmutableQuantity(bigIntValue *big.Int, int32Value int32) ImmutableQuantity {
	bigIntCopy := new(big.Int).Set(bigIntValue)
	return ImmutableQuantity{bigInt: bigIntCopy, int32Field: int32Value}
}

func (q ImmutableQuantity) Add(q2 ImmutableQuantity) ImmutableQuantity {
	newBigInt := new(big.Int).Add(q.bigInt, q2.bigInt)
	return NewImmutableQuantity(newBigInt, q.int32Field+q2.int32Field)
}

func (q ImmutableQuantity) Sub(q2 ImmutableQuantity) ImmutableQuantity {
	newBigInt := new(big.Int).Sub(q.bigInt, q2.bigInt)
	return NewImmutableQuantity(newBigInt, q.int32Field-q2.int32Field)
}

func (q ImmutableQuantity) Mul(q2 ImmutableQuantity) ImmutableQuantity {
	newBigInt := new(big.Int).Mul(q.bigInt, q2.bigInt)
	return NewImmutableQuantity(newBigInt, q.int32Field*q2.int32Field)
}

func (q ImmutableQuantity) Div(q2 ImmutableQuantity) ImmutableQuantity {
	newBigInt := new(big.Int).Div(q.bigInt, q2.bigInt)
	return NewImmutableQuantity(newBigInt, q.int32Field/q2.int32Field)
}

func BenchmarkMutableAdd(b *testing.B) {
	q1 := MutableQuantity{bigInt: big.NewInt(1000000000), int32Field: 5} // larger big.Int
	q2 := MutableQuantity{bigInt: big.NewInt(2000000000), int32Field: 10}

	for i := 0; i < b.N; i++ {
		q1.Add(q2)                                                          // Modify in place
		q1 = MutableQuantity{bigInt: big.NewInt(1000000000), int32Field: 5} // Reset for next iteration
	}
}

func BenchmarkImmutableAdd(b *testing.B) {
	q1 := NewImmutableQuantity(big.NewInt(1000000000), 5) // larger big.Int
	q2 := NewImmutableQuantity(big.NewInt(2000000000), 10)

	for i := 0; i < b.N; i++ {
		q1.Add(q2)                                           // Create a new value
		q1 = NewImmutableQuantity(big.NewInt(1000000000), 5) // Reset for next iteration
	}
}

func BenchmarkMutableSub(b *testing.B) {
	q1 := MutableQuantity{bigInt: big.NewInt(1000000000), int32Field: 5} // larger big.Int
	q2 := MutableQuantity{bigInt: big.NewInt(2000000000), int32Field: 10}

	for i := 0; i < b.N; i++ {
		q1.Sub(q2)                                                          // Modify in place
		q1 = MutableQuantity{bigInt: big.NewInt(1000000000), int32Field: 5} // Reset for next iteration
	}
}

func BenchmarkImmutableSub(b *testing.B) {
	q1 := NewImmutableQuantity(big.NewInt(1000000000), 5) // larger big.Int
	q2 := NewImmutableQuantity(big.NewInt(2000000000), 10)

	for i := 0; i < b.N; i++ {
		q1.Sub(q2)                                           // Create a new value
		q1 = NewImmutableQuantity(big.NewInt(1000000000), 5) // Reset for next iteration
	}
}

func BenchmarkMutableMul(b *testing.B) {
	q1 := MutableQuantity{bigInt: big.NewInt(1000000000), int32Field: 5} // larger big.Int
	q2 := MutableQuantity{bigInt: big.NewInt(2000000000), int32Field: 10}

	for i := 0; i < b.N; i++ {
		q1.Mul(q2)                                                          // Modify in place
		q1 = MutableQuantity{bigInt: big.NewInt(1000000000), int32Field: 5} // Reset for next iteration
	}
}

func BenchmarkImmutableMul(b *testing.B) {
	q1 := NewImmutableQuantity(big.NewInt(1000000000), 5) // larger big.Int
	q2 := NewImmutableQuantity(big.NewInt(2000000000), 10)

	for i := 0; i < b.N; i++ {
		q1.Mul(q2)                                           // Create a new value
		q1 = NewImmutableQuantity(big.NewInt(1000000000), 5) // Reset for next iteration
	}
}

func BenchmarkMutableDiv(b *testing.B) {
	q1 := MutableQuantity{bigInt: big.NewInt(1000000000), int32Field: 5} // larger big.Int
	q2 := MutableQuantity{bigInt: big.NewInt(2000000000), int32Field: 10}

	for i := 0; i < b.N; i++ {
		q1.Div(q2)                                                          // Modify in place
		q1 = MutableQuantity{bigInt: big.NewInt(1000000000), int32Field: 5} // Reset for next iteration
	}
}

func BenchmarkImmutableDiv(b *testing.B) {
	q1 := NewImmutableQuantity(big.NewInt(1000000000), 5) // larger big.Int
	q2 := NewImmutableQuantity(big.NewInt(2000000000), 10)

	for i := 0; i < b.N; i++ {
		q1.Div(q2)                                           // Create a new value
		q1 = NewImmutableQuantity(big.NewInt(1000000000), 5) // Reset for next iteration
	}
}
