package gowindows

import (
	"testing"
	"unsafe"
)

func TestTcpmilStruct(t *testing.T) {
	if ptrSize == 8 {
		if unsafe.Sizeof(MibTcpTable2{}) != 32 {
			t.Errorf("MibTcpTable2 %v!=32", unsafe.Sizeof(MibTcpTable2{}))
		}
		if unsafe.Sizeof(MibTcpRow2{}) != 28 {
			t.Errorf("MibTcpRow2 %v!=28", unsafe.Sizeof(MibTcpRow2{}))
		}
	} else {
		if unsafe.Sizeof(MibTcpTable2{}) != 32 {
			t.Errorf("MibTcpTable2 %v!=32", unsafe.Sizeof(MibTcpTable2{}))
		}
		if unsafe.Sizeof(MibTcpRow2{}) != 28 {
			t.Errorf("MibTcpRow2 %v!=28", unsafe.Sizeof(MibTcpRow2{}))
		}
	}
}
