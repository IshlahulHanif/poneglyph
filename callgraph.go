package poneglyph

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

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

		functionName, err := getFunctionName(pc)
		if err != nil {
			return
		}

		locationLine, err := getLocationLines(file, line)
		if err != nil {
			return
		}

		functionLines = append(functionLines, functionName)
		locationLines = append(locationLines, locationLine)
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

// JAVA example
// Exception in thread "main" java.lang.NullPointerException
//    at com.example.myproject.Book.getTitle(Book.java:16)
//    at com.example.myproject.Author.getBookTitles(Author.java:25)
//    at com.example.myproject.Bootstrap.main(Bootstrap.java:14)

// IDEAS:
// make function have extra opt param
// the opt param is a struct
// with the type attached to it, and what does the filter does