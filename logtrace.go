package logtrace

import (
	"fmt"
	"runtime/debug"
	"strings"
)

// Log configs
var (
	projectName            string
	isPrintFromContentRoot bool
	isPrintFunctionName    bool
	isPrintNewline         bool
	isUseTabSeparator      bool
)

// GetLogErrorTrace TODO: maybe also read at which context this occurs
func GetLogErrorTrace(err error) string {
	stackStr := string(debug.Stack()[:])
	stackSplitNewline := strings.Split(stackStr, "\n")

	var (
		locationLines   []string
		functionLines   []string
		currentFunction = strings.Split(stackSplitNewline[5], "(")[0] + "()"
		errorLogMessage string
	)
	for i, line := range stackSplitNewline {
		// bypass this function from stacktrace
		if i < 5 || len(line) < 1 {
			continue
		}

		// filter the lines where it only contains code location
		if line[0] == '\t' {
			// removing the memory address & \t
			location := strings.Split(line, " ")[0]

			// check if trim the content Root name
			if isPrintFromContentRoot {
				location = strings.Split(location, projectName)[1]
			}
			locationLines = append(locationLines, location[1:])
		}

		// filter the lines containing function names & detail
		if isPrintFunctionName {
			if line[0] != '\t' {
				// removing the memory address
				function := strings.Split(line, "(")[0]
				functionLines = append(functionLines, function)
			}
		}
	}

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

	// assembling error message
	errorLogMessage += fmt.Sprintf("Error: \"%s\" in %s%s", err.Error(), currentFunction, lineSeparator)

	// different option for printing style
	if isPrintFunctionName {
		for i := range locationLines {
			errorLogMessage += fmt.Sprintf("%s at %s( %s )%s", separator, functionLines[i], locationLines[i], lineSeparator)
		}
	} else {
		for _, location := range locationLines {
			errorLogMessage += fmt.Sprintf("%s at %s%s", separator, location, lineSeparator)
		}
	}

	return errorLogMessage
}

func SetProjectName(name string) {
	projectName = name
}

func SetIsPrintFromContentRoot(isPrint bool) {
	isPrintFromContentRoot = isPrint
}

func SetIsPrintFunctionName(isPrint bool) {
	isPrintFunctionName = isPrint
}

func SetIsPrintNewline(isPrint bool) {
	isPrintNewline = isPrint
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
