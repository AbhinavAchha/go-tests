package main

import (
	"fmt"
	"strings"
)

const s = "_"
const times = 100

func main() {
	var b strings.Builder
	b.Grow(times)
	b.WriteString(s)
	for b.Len() < times {
		fmt.Printf("b.Len(): %v\n", b.Len())
		chunk := times - b.Len()
		if chunk > b.Len() {
			chunk = b.Len()
		}
		b.WriteString(b.String()[:chunk])
	}

}
