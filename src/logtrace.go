package logtrace

import "fmt"

// GetLogErrorTrace returns a string of error stack depending on the configs
// if there are any error occurs, it will return the original error
func GetLogErrorTrace(errArg error) (errorTraceResult error) {
	return getLogErrorTraceSkip(errArg, 2)
}

func PrintLogErrorTrace(errArg error) {
	fmt.Println(getLogErrorTraceSkip(errArg, 2))
}
