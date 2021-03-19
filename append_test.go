package main

import (
	"fmt"
	"testing"
)

func BenchmarkAppend_AllocateEveryTime(b *testing.B) {
	container := []string{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		container = append(container, fmt.Sprintf("No.%d", i))
	}
}

func BenchmarkAppend_AllocateOnce(b *testing.B) {
	container := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		container[i] = fmt.Sprintf("No.%d", i)
	}
}
