package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 74382},
	{input: 382399},
}

func BenchmarkRunTrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunTrace(10)
	}
}

func BenchmarkRunTraceWithTable(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("function depth: %d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RunTrace(v.input)
			}
		})
	}
}

func BenchmarkRunErrorLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunTdkErrorLog(10)
	}
}

func BenchmarkRunErrorLogWithTable(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("function depth: %d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RunTdkErrorLog(v.input)
			}
		})
	}
}
