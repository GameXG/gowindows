package gowindows

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	rpcrt4     = windows.NewLazyDLL("Rpcrt4.dll")
	uuidCreate = rpcrt4.NewProc("UuidCreate")
)

//RPCRTAPI RPC_STATUS UuidCreate(
//UUID *Uuid
//);
func UuidCreate(v *GUID) error {
	r1, _, e1 := uuidCreate.Call(uintptr(unsafe.Pointer(v)))
	if r1 != RPC_S_OK {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetNewUuid() (*GUID, error) {
	v := GUID{}

	err := UuidCreate(&v)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
