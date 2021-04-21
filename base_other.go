// +build darwin,386 darwin,amd64 dragonfly freebsd linux nacl netbsd openbsd solaris

package gowindows

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}
type SID struct{}
type Pointer *struct{}
type Handle uintptr
