package gowindows

import (
	"strings"

	"strconv"

	"fmt"

	"encoding/binary"
)

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

func GUIDFormString(s string) (guid GUID, err error) {
	s = strings.Trim(s, "{}")
	ss := strings.Split(s, "-")

	if len(ss) != 5 {
		err = fmt.Errorf("len(-)%v!=5", len(ss))
		return
	}

	ss[3] = fmt.Sprint(ss[3], ss[4])

	var i uint64

	i, err = strconv.ParseUint(ss[0], 16, 32)
	if err != nil {
		return GUID{}, err
	}
	guid.Data1 = uint32(i)

	i, err = strconv.ParseUint(ss[1], 16, 16)
	if err != nil {
		return GUID{}, err
	}
	guid.Data2 = uint16(i)

	i, err = strconv.ParseUint(ss[2], 16, 16)
	if err != nil {
		return GUID{}, err
	}
	guid.Data3 = uint16(i)

	i, err = strconv.ParseUint(ss[3], 16, 64)
	if err != nil {
		return GUID{}, err
	}
	binary.BigEndian.PutUint64(guid.Data4[:], i)
	return
}
