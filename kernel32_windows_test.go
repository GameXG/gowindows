package gowindows

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"
	"unsafe"

	"runtime"

	"syscall"

	"golang.org/x/sys/windows"
)

func TestReadProcessMemory_self(t *testing.T) {
	b := make([]byte, 100)
	rand.Read(b)

	addr := &b[0]

	processHandle, err := windows.GetCurrentProcess()
	if err != nil {
		t.Fatal(err)
	}

	b2 := make([]byte, 100)
	size := uint(0)
	err = ReadProcessMemory(Handle(processHandle), uint(uintptr(unsafe.Pointer(addr))),
		windows.Pointer(unsafe.Pointer(&b2[0])), uint(len(b2)), &size)
	if err != nil {
		t.Fatal(err)
	}

	if size != uint(len(b2)) {
		t.Errorf("%v!=%v", size, len(b2))
	}

	if bytes.Equal(b, b2) == false {
		t.Errorf("%#v!=%#v", b, b2)
	}
}

func TestIsWow64Process(t *testing.T) {
	h, err := windows.GetCurrentProcess()
	if err != nil {
		t.Fatal(err)
	}

	iswow, err := IsWow64Process(Handle(h))
	if err != nil {
		t.Fatal(err)
	}

	// 这个测试前提是操作系统是64位，否则这个测试结果会不正确。
	switch runtime.GOARCH {
	case "amd64":
		if iswow == true {
			t.Fatal("请确认操作系统是否是64位？")
		}
	case "386":
		if iswow == false {
			t.Fatal("请确认操作系统是否是64位？")
		}
	default:
		t.Fatalf("未知的 runtime.GOARCH %v", runtime.GOARCH)
	}
}

func TestIs64System(t *testing.T) {
	is64, err := Is64System()
	if err != nil {
		t.Fatal(err)
	}

	// 这个测试前提是操作系统是64位，否则这个测试结果会不正确。
	if is64 == false {
		t.Fatal("请确认操作系统是否是64位？")
	}

}

func TestOpenFileMapping(t *testing.T) {
	name := "name456742014252432"
	namePtr, _ := windows.UTF16PtrFromString(name)

	h, err := windows.CreateFileMapping(windows.Handle(INVALID_HANDLE_VALUE),
		nil, uint32(syscall.PAGE_READWRITE), 0, uint32(4096), namePtr)
	if err != nil {
		t.Fatal(err)
	}

	addr, err := windows.MapViewOfFile(h, uint32(syscall.FILE_MAP_READ), 0,
		0, uintptr(4096))
	if err != nil {
		t.Fatal(err)
	}

	b := *((*[4096]byte)(unsafe.Pointer(addr)))
	_ = b

	h2, err := OpenFileMapping(syscall.PAGE_READWRITE, false, name)
	if err != nil {
		t.Fatal(err)
	}

	addr2, err := windows.MapViewOfFile(windows.Handle(h2), uint32(syscall.FILE_MAP_READ), 0,
		0, uintptr(4096))
	if err != nil {
		t.Fatal(err)
	}

	_ = addr2
}

func TestGetModuleFileName(t *testing.T) {
	path, err := GetModuleFileName(0)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("path:%v\r\n", path)
}

func TestStr2MultiByteAndMultiByte2Str(t *testing.T) {
	data := []string{
		"abc",
		"abc中文测试",
	}

	for _, v := range data {
		d, err := Str2MultiByte(CP_ACP, WC_NO_BEST_FIT_CHARS, v, nil, nil)
		if err != nil {
			t.Errorf("Str2MultiByte,%v", err)
			continue
		}

		s, err := MultiByte2Str(CP_ACP, 0, d)
		if err != nil {
			t.Errorf("MultiByte2Str, %v", err)
			continue
		}

		if v != s {
			t.Errorf("%v != %v", v, s)
		}
	}
}
