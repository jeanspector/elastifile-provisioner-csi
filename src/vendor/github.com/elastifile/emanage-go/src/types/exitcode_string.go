// Code generated by "stringer -type=ExitCode"; DO NOT EDIT

package types

import "fmt"

const (
	_ExitCode_name_0 = "ExitCodeSuccessExitCodeTeslaFailedExitCodeToolFailedExitCodeVerifyFailedExitCodeFailed"
	_ExitCode_name_1 = "ExitCodeTesterFocus"
)

var (
	_ExitCode_index_0 = [...]uint8{0, 15, 34, 52, 72, 86}
	_ExitCode_index_1 = [...]uint8{0, 19}
)

func (i ExitCode) String() string {
	switch {
	case 0 <= i && i <= 4:
		return _ExitCode_name_0[_ExitCode_index_0[i]:_ExitCode_index_0[i+1]]
	case i == 197:
		return _ExitCode_name_1
	default:
		return fmt.Sprintf("ExitCode(%d)", i)
	}
}