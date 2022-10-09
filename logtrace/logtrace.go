package logtrace

import (
	"fmt"
	"runtime/debug"
	"strings"
)

var (
	projectName string
)

// PrintLogErrorTrace TODO: maybe also read at which context this occurs
func PrintLogErrorTrace(err error) {
	projectName = "sauron"

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
			locationLines = append(locationLines, location[1:])
		}
		// filter the lines containing function names & detail
		if line[0] != '\t' {
			// removing the memory address
			function := strings.Split(line, "(")[0]
			functionLines = append(functionLines, function+"()")
		}
	}

	// assembling error message
	errorLogMessage += "\nError: \"" + err.Error() + "\" in " + currentFunction + "\n"
	for _, location := range locationLines {
		errorLogMessage += "\t at " + location + "\n"
	}

	fmt.Println(errorLogMessage)
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
