package main

import (
	"fmt"
	"strings"
	"testing"
)

const alias = "a"

var fields = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u"}

func BenchmarkJoin_Strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf(" %s.`", alias) + strings.Join(fields, fmt.Sprintf("`,%s.`", alias))
	}
}

func BenchmarkJoin_Iterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf(" %s.`%s", alias, fields[0])
		for _, f := range fields[1:] {
			_ = fmt.Sprintf("`,%s.`%s", alias, f)
		}
	}
}
