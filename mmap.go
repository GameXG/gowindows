package gowindows

import "C"
import (
	"unsafe"
)

type Mmap struct {
	fileHandle Handle
	//TODO: 这里建议更换类型， 内存 gc 时会统计 unsafe.Pointer 引用的内存，换成uintptr来避开内存 gc 的问题。
	addr unsafe.Pointer
	size int
}

func (m *Mmap) GetHandle() Handle {
	return m.fileHandle
}

func (m *Mmap) GetBytes() []byte {
	if m.size == 0 || m.addr == unsafe.Pointer(uintptr(0)) {
		return nil
	}

	return ToBytes(uintptr(m.addr), m.size, m.size)
}
