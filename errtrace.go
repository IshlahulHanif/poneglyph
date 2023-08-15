package poneglyph

import (
	"errors"
	"fmt"
	"runtime"
)

func (l *ErrorTrace) Error() string {
	if l == nil || l.err == nil {
		return ""
	}

	return l.err.Error()
}

func (l *ErrorTrace) Trace(messages ...string) {
	if l == nil || l.err == nil {
		return
	}

	l.doTrace(2)

	var message string
	if len(messages) > 0 {
		message = fmt.Sprint(messages)
	}
	l.messages = append(l.messages, message)
}

func (l *ErrorTrace) doTrace(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	functionName, err := getFunctionName(pc)
	if err != nil {
		return
	}

	locationLine, err := getLocationLines(file, line)
	if err != nil {
		return
	}

	l.functionLines = append(l.functionLines, functionName)
	l.locationLines = append(l.locationLines, locationLine)
}

func (l *ErrorTrace) GetErrorTraceMessage() string {
	// set lineSeparator between lines
	var (
		lineSeparator, separator, callerFunctionName, errorTraceMessage string
	)

	// check length
	if len(l.locationLines) == 0 || len(l.functionLines) == 0 {
		return l.Error()
	}

	callerFunctionName = l.functionLines[len(l.functionLines)-1]

	// creates separator
	if isUseTabSeparator {
		separator = "\t"
	} else {
		separator = FourSpace
	}

	// creates line separator
	if isPrintNewline {
		lineSeparator = "\n"
	} else {
		lineSeparator = separator + "||"
	}

	// assembling error message
	errorTraceMessage += fmt.Sprintf("Error: \"%s\" in %s%s", l.Error(), callerFunctionName, lineSeparator)

	for i := range l.locationLines {
		// add separator
		errorTraceMessage += separator

		// add space and 'at' message
		errorTraceMessage += " at "

		// different option for printing function & location
		if isPrintFunctionName {
			errorTraceMessage += l.functionLines[i]
			errorTraceMessage += " ( "
		}

		errorTraceMessage += l.locationLines[i]

		if isPrintFunctionName {
			errorTraceMessage += " )"
		}

		// print extra message if any
		if len(l.messages[i]) > 0 {
			errorTraceMessage += fmt.Sprintf(" [Messages] %s", l.messages[i])
		}

		errorTraceMessage += lineSeparator
	}

	return errorTraceMessage
}

func Trace(err error, messages ...string) *ErrorTrace {
	if err == nil {
		return nil
	}

	var errorTrace *ErrorTrace
	ok := errors.As(err, &errorTrace)
	if !ok {
		errorTrace = &ErrorTrace{err: err}
	}

	errorTrace.doTrace(2)

	var message string
	if len(messages) > 0 {
		message = fmt.Sprint(messages)
	}
	errorTrace.messages = append(errorTrace.messages, message)

	return errorTrace
}

func TraceStr(errStr string, messages ...string) *ErrorTrace {
	if len(errStr) == 0 {
		return nil
	}

	err := errors.New(errStr)

	var errorTrace *ErrorTrace
	ok := errors.As(err, &errorTrace)
	if !ok {
		errorTrace = &ErrorTrace{err: err}
	}

	errorTrace.doTrace(2)

	var message string
	if len(messages) > 0 {
		message = fmt.Sprint(messages)
	}
	errorTrace.messages = append(errorTrace.messages, message)

	return errorTrace
}

func GetErrorLogMessage(err error) string {
	if err == nil {
		return ""
	}

	var errorTrace *ErrorTrace
	ok := errors.As(err, &errorTrace)
	if !ok {
		return err.Error()
	}

	return errorTrace.GetErrorTraceMessage()
}
