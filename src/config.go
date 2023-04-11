package poneglyph

// Log configs
var (
	projectName                    string
	isPrintRelativeFromContentRoot bool
	isPrintFunctionName            bool
	isUseSimpleFunctionName        bool
	isPrintNewline                 bool
	isUseTabSeparator              bool
	isSkipNonProject               bool
	stackLimit                     int
)

// TODO: set config should only be called once
func SetConfig(config Config) {
	projectName = config.projectName
	isPrintRelativeFromContentRoot = config.isPrintRelativeFromContentRoot
	isPrintFunctionName = config.isPrintFunctionName
	isUseSimpleFunctionName = config.isUseSimpleFunctionName
	isPrintNewline = config.isPrintNewline
	isUseTabSeparator = config.isUseTabSeparator
	isSkipNonProject = config.isSkipNonProject
	stackLimit = config.stackLimit
}

func init() {
	stackLimit = 100
}

func SetProjectName(name string) {
	projectName = name
}

func SetIsPrintFromContentRoot(isPrint bool) {
	isPrintRelativeFromContentRoot = isPrint
}

func SetIsPrintFunctionName(isPrint bool) {
	isPrintFunctionName = isPrint
}

func SetIsPrintNewline(isPrint bool) {
	isPrintNewline = isPrint
}

func SetIsUseSimpleFunctionName(isUse bool) {
	isUseSimpleFunctionName = isUse
}

func SetIsUseTabSeparator(isUse bool) {
	isUseTabSeparator = isUse
}

func SetIsSkipNonProject(isSkip bool) {
	isSkipNonProject = isSkip
}

func SetStackLimit(limit int) {
	stackLimit = limit
}
