// +build windows

package gowindows

import (
	"syscall"
)

func abort(funcname string, err int) {
	panic(funcname + " failed: " + syscall.Errno(err).Error())
}

const (
	MB_OK                uint = 0x00000000
	MB_OKCANCEL          uint = 0x00000001
	MB_ABORTRETRYIGNORE  uint = 0x00000002
	MB_YESNOCANCEL       uint = 0x00000003
	MB_YESNO             uint = 0x00000004
	MB_RETRYCANCEL       uint = 0x00000005
	MB_CANCELTRYCONTINUE uint = 0x00000006
	MB_ICONHAND          uint = 0x00000010
	MB_ICONQUESTION      uint = 0x00000020
	MB_ICONEXCLAMATION   uint = 0x00000030
	MB_ICONASTERISK      uint = 0x00000040
	MB_USERICON          uint = 0x00000080
	MB_ICONWARNING       uint = MB_ICONEXCLAMATION
	MB_ICONERROR         uint = MB_ICONHAND
	MB_ICONINFORMATION   uint = MB_ICONASTERISK
	MB_ICONSTOP          uint = MB_ICONHAND
	MB_DEFBUTTON1        uint = 0x00000000
	MB_DEFBUTTON2        uint = 0x00000100
	MB_DEFBUTTON3        uint = 0x00000200
	MB_DEFBUTTON4        uint = 0x00000300

	WM_COPYDATA = 74
)

type COPYDATASTRUCT struct {
	dwData uintptr
	cbData uint32
	lpData uintptr
}
