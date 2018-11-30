package gowindows

import (
	"unsafe"

	"syscall"

	"golang.org/x/sys/windows"
)

var (
	ole32           = windows.NewLazyDLL("ole32.dll")
	stringFromGUID2 = ole32.NewProc("StringFromGUID2")
	//cLSIDFromString = ole32.NewProc("CLSIDFromString")
)

func StringFromGUID2(guid *GUID) (string, error) {
	size := 40
	b := make([]uint16, size)

	r1, _, e1 := stringFromGUID2.Call(uintptr(unsafe.Pointer(guid)), uintptr(unsafe.Pointer(&b[0])), uintptr(size))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return "", e1
		} else {
			return "", syscall.EINVAL
		}
	}

	return windows.UTF16ToString(b[:r1]), nil
}

/*
// CLSID == GUID
// windows 的要求必须系统注册了，否则会报错无法转换
func GUIDFromString(s string) (guid GUID, err error) {
	var a *uint16
	a, err = windows.UTF16PtrFromString(s)
	if err != nil {
		return
	}
	r1, _, e1 := cLSIDFromString.Call(uintptr(unsafe.Pointer(a)), uintptr(unsafe.Pointer(&guid)))

}
*/
