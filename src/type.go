package poneglyph

type (
	Config struct {
		projectName                    string
		isPrintRelativeFromContentRoot bool
		isPrintFunctionName            bool
		isUseSimpleFunctionName        bool
		isPrintNewline                 bool
		isUseTabSeparator              bool
		isSkipNonProject               bool
		stackLimit                     int
	}

	ErrorTrace struct {
		// base error
		err error

		// traces
		locationLines []string
		functionLines []string
		messages      []string
	}
)
