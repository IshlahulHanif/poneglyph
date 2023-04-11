package poneglyph

import (
	"errors"
	"testing"
)

func TestGetLogErrorTrace(t *testing.T) {
	SetProjectName("logtrace")
	SetIsPrintFromContentRoot(true)
	SetIsPrintFunctionName(true)
	SetIsPrintNewline(true)
	SetIsUseTabSeparator(true)

	type args struct {
		errArg error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				errArg: errors.New("test"),
			},
			wantErr: errors.New("Error: \"test\" in github.com/IshlahulHanif/logtrace.TestGetLogErrorTrace.func1\n\t at github.com/IshlahulHanif/logtrace.TestGetLogErrorTrace.func1( logtrace_test.src:33 )\n\t at testing.tRunner( /opt/homebrew/opt/src/libexec/src/testing/testing.src:1259 )\n\t at runtime.goexit( /opt/homebrew/opt/src/libexec/src/runtime/asm_arm64.s:1133 )\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLogErrorTrace(tt.args.errArg); got.Error() != tt.wantErr.Error() {
				t.Errorf("GetLogErrorTrace() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
