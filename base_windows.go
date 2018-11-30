package gowindows

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/windows"
)

type GUID = windows.GUID
type SID = windows.SID
type Pointer = windows.Pointer
type Handle = syscall.Handle

func FormatMessage(errno uint32, msgSrc uintptr) (string, error) {
	const flags uint32 = windows.FORMAT_MESSAGE_ALLOCATE_BUFFER | windows.FORMAT_MESSAGE_FROM_HMODULE | windows.FORMAT_MESSAGE_FROM_SYSTEM | windows.FORMAT_MESSAGE_ARGUMENT_ARRAY | windows.FORMAT_MESSAGE_IGNORE_INSERTS
	buf := make([]uint16, 300)
	_, err := windows.FormatMessage(flags, uintptr(msgSrc), uint32(errno), 0, buf, nil)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return windows.UTF16ToString(buf), nil
}
