package main

import (
	"errors"
	"fmt"
	"github.com/logtrace/logtrace"
)

func main() {
	RunTrace(10)
}

func RunTrace(depth int) {
	logtrace.SetProjectName("logtrace")
	logtrace.SetIsPrintFromContentRoot(true)
	logtrace.SetIsPrintFunctionName(true)
	logtrace.SetIsPrintNewline(true)
	logtrace.SetIsUseTabSeparator(false)
	if depth > 0 {
		RunTrace(depth - 1)
	} else {
		fmt.Print(logtrace.GetLogErrorTrace(errors.New("test err")))
	}
}

func RunTdkErrorLog(depth int) {
	fmt.Errorf("shlh test error\n")
	if depth > 0 {
		RunTdkErrorLog(depth - 1)
	}
}
