package gowindows

import "C"

type Mmap struct {
	fileHandle Handle
	addr       uintptr
	size       int
}

func (m *Mmap) GetHandle() Handle {
	return m.fileHandle
}

func (m *Mmap) GetBytes() []byte {
	if m.size == 0 || m.addr == uintptr(0) {
		return nil
	}

	return ToBytes(m.addr, m.size, m.size)
}
