// Code generated by "stringer -type=OS"; DO NOT EDIT.

package drmaa2interface

import "fmt"

const _OS_name = "OtherOSAIXBSDLinuxHPUXIRIXMacOSSunOSTRU64UnixWareWinWinNT"

var _OS_index = [...]uint8{0, 7, 10, 13, 18, 22, 26, 31, 36, 41, 49, 52, 57}

func (i OS) String() string {
	if i < 0 || i >= OS(len(_OS_index)-1) {
		return fmt.Sprintf("OS(%d)", i)
	}
	return _OS_name[_OS_index[i]:_OS_index[i+1]]
}
