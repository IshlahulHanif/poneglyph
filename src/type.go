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
)
