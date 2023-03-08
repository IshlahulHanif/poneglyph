package logtrace

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// Log configs
var (
	projectName                    string
	isPrintRelativeFromContentRoot bool
	isPrintFunctionName            bool
	isUseSimpleFunctionName        bool
	isPrintNewline                 bool
	isUseTabSeparator              bool
)

// GetLogErrorTrace returns a string of error stack depending on the configs
// if there are any error occurs, it will return the original error
func GetLogErrorTrace(errArg error) (errorTraceResult error) {
	var (
		locationLines      []string
		functionLines      []string
		callerFunctionName string
	)

	defer func(errArg error) {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			errorTraceResult = errArg
		}
	}(errArg)

	for i := 1; i < 100; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		callerFunction := runtime.FuncForPC(pc)
		functionName := callerFunction.Name()

		if isUseSimpleFunctionName && isPrintFunctionName {
			functionSplit := strings.Split(functionName, "/")
			if len(functionSplit) > 0 {
				functionName = functionSplit[len(functionSplit)-1]
			}
		}

		functionLines = append(functionLines, functionName)

		// get working dir
		cwd, err := filepath.Abs(".")
		if err != nil {
			return errArg
		}

		relPath, err := filepath.Rel(cwd, file)
		if err != nil {
			return errArg
		}

		location := fmt.Sprintf("%s:%d", file, line)
		if isPrintRelativeFromContentRoot && !strings.Contains(relPath, "../") {
			location = fmt.Sprintf("%s:%d", relPath, line)
		}
		locationLines = append(locationLines, location)
	}

	// check length
	if len(locationLines) == 0 || len(functionLines) == 0 {
		return errArg
	}

	callerFunctionName = functionLines[0]

	// set lineSeparator between lines
	var (
		lineSeparator string
		separator     string
	)

	if isUseTabSeparator {
		separator = "\t"
	} else {
		separator = FourSpace
	}

	if isPrintNewline {
		lineSeparator = "\n"
	} else {
		lineSeparator = separator + "||"
	}

	var errorTraceMessage string
	// assembling error message
	errorTraceMessage += fmt.Sprintf("Error: \"%s\" in %s%s", errArg.Error(), callerFunctionName, lineSeparator)

	// different option for printing style
	if isPrintFunctionName {
		for i := range locationLines {
			errorTraceMessage += fmt.Sprintf("%s at %s( %s )%s", separator, functionLines[i], locationLines[i], lineSeparator)
		}
	} else {
		for _, location := range locationLines {
			errorTraceMessage += fmt.Sprintf("%s at %s%s", separator, location, lineSeparator)
		}
	}

	errorTraceResult = errors.New(errorTraceMessage)

	return errorTraceResult
}

func PrintLogErrorTrace(errArg error) {
	fmt.Println(GetLogErrorTrace(errArg))
}

func SetProjectName(name string) {
	projectName = name
}

func SetIsPrintFromContentRoot(isPrint bool) {
	isPrintRelativeFromContentRoot = isPrint
}

func SetIsPrintFunctionName(isPrint bool) {
	isPrintFunctionName = isPrint
}

func SetIsPrintNewline(isPrint bool) {
	isPrintNewline = isPrint
}

func SetIsUseSimpleFunctionName(isUse bool) {
	isUseSimpleFunctionName = isUse
}

func SetIsUseTabSeparator(isUse bool) {
	isUseTabSeparator = isUse
}

const (
	FourSpace = "   "
)

// JAVA example
// Exception in thread "main" java.lang.NullPointerException
//    at com.example.myproject.Book.getTitle(Book.java:16)
//    at com.example.myproject.Author.getBookTitles(Author.java:25)
//    at com.example.myproject.Bootstrap.main(Bootstrap.java:14)

// IDEAS:
// make function have extra opt param
// the opt param is a struct
// with the type attached to it, and what does the filter does
