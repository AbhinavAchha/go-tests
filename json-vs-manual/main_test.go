package main

import (
	"encoding/json"
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
