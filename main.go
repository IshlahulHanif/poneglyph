package main

import (
	"errors"
	"github.com/logtrace/logtrace"
	"github.com/tokopedia/tdk/go/log"
)

func main() {
	RunTrace(10)
}

func RunTrace(depth int) {
	if depth > 0 {
		RunTrace(depth - 1)
	} else {
		logtrace.PrintLogErrorTrace(errors.New("shlh test err"))
	}
}

func RunTdkErrorLog(depth int) {
	log.Error("shlh test error")
	if depth > 0 {
		RunTdkErrorLog(depth - 1)
	}
}
