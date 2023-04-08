package logtrace

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// GetLogErrorTrace returns a string of error stack depending on the configs
// if there are any error occurs, it will return the original error
func GetLogErrorTrace(errArg error) (errorTraceResult error) {
	return getLogErrorTraceSkip(errArg, 2)
}

func getLogErrorTraceSkip(errArg error, skip int) (errorTraceResult error) {
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

	for i := skip; i < stackLimit; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		if isSkipNonProject {
			if !strings.Contains(file, projectName) {
				continue
			}
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
	fmt.Println(getLogErrorTraceSkip(errArg, 2))
}

// JAVA example
// Exception in thread "main" java.lang.NullPointerException
//    at com.example.myproject.Book.getTitle(Book.java:16)
//    at com.example.myproject.Author.getBookTitles(Author.java:25)
//    at com.example.myproject.Bootstrap.main(Bootstrap.java:14)

// IDEAS:
// make function have extra opt param
// the opt param is a struct
// with the type attached to it, and what does the filter does
