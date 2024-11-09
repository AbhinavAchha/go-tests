package main_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"
)

var s = strings.Repeat("\"Hello, World!\"", 1)

func BenchmarkIoRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(s)
		d, err := io.ReadAll(r)
		if err != nil {
			b.Fatal(err)
		}

		_ = string(d[1 : len(d)-1])
	}
}

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str string
		if err := json.NewDecoder(strings.NewReader(s)).Decode(&str); err != nil {
			b.Fatal(err)
		}
	}
}
