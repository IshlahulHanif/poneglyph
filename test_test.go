package logtrace

import (
	"errors"
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

func RunTrace(depth int) {
	SetProjectName("test")
	SetIsPrintFromContentRoot(true)
	SetIsPrintFunctionName(true)
	SetIsPrintNewline(true)
	SetIsUseTabSeparator(false)
	if depth > 0 {
		RunTrace(depth - 1)
	} else {
		fmt.Print(GetLogErrorTrace(errors.New("test err")))
	}
}

func RunTdkErrorLog(depth int) {
	fmt.Errorf("shlh test error\n")
	if depth > 0 {
		RunTdkErrorLog(depth - 1)
	}
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
