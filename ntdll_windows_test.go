package gowindows

import (
	"bytes"
	"crypto/rand"
	"runtime"
	"testing"
	"unsafe"

	"os/exec"

	"golang.org/x/sys/windows"
)

// set GOARCH=386
func TestNtWow64QueryInformationProcess64(t *testing.T) {
	if runtime.GOARCH != "386" {
		t.Skip("不是32位程序，跳过测试。")
	}

	//实际还应该判断是否是 64位系统，32位系统也不支持这个函数。

	p := PROCESS_BASIC_INFORMATION64{}
	size := uint32(unsafe.Sizeof(p))
	outSize := uint32(0)

	cmd := exec.Command(`C:\WINDOWS\system32\notepad.exe`)
	err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer cmd.Process.Kill()

	//	time.Sleep(1 * time.Second)

	processHandle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|PROCESS_VM_READ, false, uint32(cmd.Process.Pid))
	if err != nil {
		t.Fatal(err)
	}

	err = NtWow64QueryInformationProcess64(Handle(processHandle), ProcessBasicInformation, windows.Pointer(unsafe.Pointer(&p)), size, &outSize)
	if err != nil {
		t.Fatal(err)
	}

	if outSize > size {
		t.Fatalf("%v > %v", outSize, size)
	}

	//fmt.Printf("%#v\r\n", p)

}

func TestNtQueryInformationProcess(t *testing.T) {
	p := PROCESS_BASIC_INFORMATION{}
	size := uint32(unsafe.Sizeof(p))
	outSize := uint32(0)

	cmd := exec.Command(`C:\WINDOWS\system32\notepad.exe`)
	err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer cmd.Process.Kill()

	//	time.Sleep(1 * time.Second)

	processHandle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|PROCESS_VM_READ, false, uint32(cmd.Process.Pid))
	if err != nil {
		t.Fatal(err)
	}

	err = NtQueryInformationProcess(Handle(processHandle), ProcessBasicInformation, windows.Pointer(unsafe.Pointer(&p)), size, &outSize)
	if err != nil {
		t.Fatal(err)
	}

	if outSize > size {
		t.Fatalf("%v > %v", outSize, size)
	}

	//fmt.Printf("%#v\r\n", p)
}

func TestNtWow64ReadVirtualMemory64_self(t *testing.T) {
	if runtime.GOARCH != "386" {
		t.Skip("不是32位程序，跳过测试。")
	}
	b := make([]byte, 100)
	rand.Read(b)

	addr := &b[0]

	// 直接使用 windows.GetCurrentProcess() 获得的无法使用，返回错误的句柄。
	processHandle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|PROCESS_VM_READ,
		false, uint32(windows.Getpid()))
	if err != nil {
		t.Fatal(err)
	}

	b2 := make([]byte, 100)
	size := uint64(0)
	err = NtWow64ReadVirtualMemory64(Handle(processHandle), uint64(uintptr(unsafe.Pointer(addr))),
		windows.Pointer(unsafe.Pointer(&b2[0])), uint64(len(b2)), &size)
	if err != nil {
		t.Fatal(err)
	}

	if size != uint64(len(b2)) {
		t.Errorf("%v!=%v", size, len(b2))
	}

	if bytes.Equal(b, b2) == false {
		t.Errorf("%#v!=%#v", b, b2)
	}
}

func TestReadVirtualMemory_self(t *testing.T) {
	b := make([]byte, 100)
	rand.Read(b)

	addr := &b[0]

	processHandle, err := windows.GetCurrentProcess()
	if err != nil {
		t.Fatal(err)
	}

	b2 := make([]byte, 100)
	size := uint(0)
	err = ReadVirtualMemory(Handle(processHandle), uint(uintptr(unsafe.Pointer(addr))),
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

func TestRtlGetVersion(t *testing.T) {
	version, err := RtlGetVersion()
	if err != nil {
		t.Fatal(err)
	}

	if version == nil {
		t.Fatal(err)
	}

	//fmt.Printf("windows version :%v", version.GetString())
}
