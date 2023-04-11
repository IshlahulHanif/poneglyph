# Poneglyph
Java-style logtrace printer for Golang

Do you find it tiring to always print your error on every function?

Using this lib, you can prompt to print your error just in one line, the first time you find the error at instantly print the error stack in a more readable way than the original stacktrace

User can also modify which options on how the error should be printed

Example log with error = "test err"
```azure
Error: "testing the complex function" in src.TestTrace
    at src.LastTestFunction ( errtrace_test.go:66 )
    at src.GoroutineTestingFunction.func1 ( errtrace_test.go:57 )
    at src.GoroutineTestingFunction.func1.1 ( errtrace_test.go:53 )
    at src.GoroutineTestingFunction ( errtrace_test.go:61 )
    at src.TheCoolerTestingFunction ( errtrace_test.go:43 )
    at src.AnotherTestingFunction ( errtrace_test.go:38 )
    at src.MessageIncludedTestingFunction ( errtrace_test.go:33 ) [Messages] [parameter 1 is: 1234]
    at src.OneOfTheFunctionForTesting ( errtrace_test.go:28 )
    at src.TestTrace ( errtrace_test.go:23 )
```

This lib does not actually print the error, only giving users the string which they can print it themselves using their log print lib of choice

```go
// setup the lib
poneglyph.SetConfig(Config{
    projectName:                    "poneglyph",
    isPrintRelativeFromContentRoot: true,
    isPrintFunctionName:            true,
    isUseSimpleFunctionName:        true,
    isPrintNewline:                 true,
    isUseTabSeparator:              false,
})

func NotLatestFunction() (Resp, error) {
    //...
    resp, err := Somefunction()
    if err != nil {
        return resp, poneglyph.Trace(err)
    }
    //...
}

func LatestFunctionYouWantToTrace() (Resp, error) {
    //...
    resp, err := NotLatestFunction()
    if err != nil {
        logErr := poneglyph.GetErrorLogMessage(Trace(err))
        fmt.Println(logErr)
        return resp, err
    }
    //...
}
```
