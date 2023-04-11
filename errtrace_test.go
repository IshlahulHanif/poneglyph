package poneglyph

import (
	"errors"
	"fmt"
	"testing"
)

func TestTrace(t *testing.T) {
	// Setup
	SetConfig(Config{
		projectName:                    "poneglyph",
		isPrintRelativeFromContentRoot: true,
		isPrintFunctionName:            true,
		isUseSimpleFunctionName:        true,
		isPrintNewline:                 true,
		isUseTabSeparator:              false,
		isSkipNonProject:               true, // not needed
		stackLimit:                     100,  // not needed
	})

	err := OneOfTheFunctionForTesting()
	fmt.Println(GetErrorLogMessage(Trace(err)))
}

func OneOfTheFunctionForTesting() error {
	err := MessageIncludedTestingFunction()
	return Trace(err)
}

func MessageIncludedTestingFunction() error {
	err := AnotherTestingFunction()
	return Trace(err, fmt.Sprintf("parameter 1 is: %v", 1234))
}

func AnotherTestingFunction() error {
	err := TheCoolerTestingFunction()
	return Trace(err)
}

func TheCoolerTestingFunction() error {
	err := GoroutineTestingFunction()
	return Trace(err)
}

func GoroutineTestingFunction() error {
	var errChan = make(chan error)

	go func() {
		var err error
		defer func() {
			if err != nil {
				err = Trace(err)
			}
			errChan <- err
		}()
		err = Trace(LastTestFunction())
	}()

	err := <-errChan
	return Trace(err)
}

func LastTestFunction() error {
	err := errors.New("testing the complex function")
	return Trace(err)
}
