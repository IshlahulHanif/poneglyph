package poneglyph

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func getFunctionName(pc uintptr) (string, error) {
	callerFunction := runtime.FuncForPC(pc)
	if callerFunction == nil {
		return "", errors.New(fmt.Sprintf("nil pointer for pc: %v", pc))
	}
	functionName := callerFunction.Name()

	if isUseSimpleFunctionName && isPrintFunctionName {
		functionSplit := strings.Split(functionName, "/")
		if len(functionSplit) > 0 {
			functionName = functionSplit[len(functionSplit)-1]
		}
	}

	return functionName, nil
}

func getLocationLines(file string, line int) (string, error) {
	// get working dir
	cwd, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}

	location := fmt.Sprintf("%s:%d", file, line)

	relPath, err := filepath.Rel(cwd, file)
	if err != nil {
		return "", err
	}

	if isPrintRelativeFromContentRoot && !strings.Contains(relPath, "../") {
		location = fmt.Sprintf("%s:%d", relPath, line)
	}
	return location, nil
}
