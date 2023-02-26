# Logtrace
Java-style logtrace printer for Golang

Do you find it tiring to always print your error on every function?

Using this lib, you can prompt to print your error just in one line, the first time you find the error at instantly print the error stack in a more readable way than the original stacktrace

User can also modify which options on how the error should be printed

Example log with error = "test err"
```azure
Error: "test err" in main.RunTrace()
    at main.RunTrace( main.go:22 )
    at main.RunTrace( main.go:20 )
    at main.RunTrace( main.go:20 )
    at main.RunTrace( main.go:20 )
    at main.RunTrace( main.go:20 )
    at main.RunTrace( main.go:20 )
    at main.RunTrace( main.go:20 )
    at main.RunTrace( main.go:20 )
    at main.RunTrace( main.go:20 )
    at main.RunTrace( main.go:20 )
    at main.RunTrace( main.go:20 )
    at main.main( main.go:10 )
```

This lib does not actually print the error, only giving users the string which they can print it themselves using their log print lib of choice

```go
logtrace.SetProjectName("logtrace")
logtrace.SetIsPrintFromContentRoot(true)
logtrace.SetIsPrintFunctionName(true)
logtrace.SetIsPrintNewline(true)
logtrace.SetIsUseTabSeparator(false)

logErr := logtrace.GetLogErrorTrace(errors.New("test err"))
fmt.Print(logErr)
```
