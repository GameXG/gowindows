package gowindows

import "C"
import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// 创建无文件命名内存共享
func CreateMmap(name string, size uint32, write bool) (*Mmap, error) {
	return CreateMmapWithSecurityDescriptor(name, size, write, "")
}

// 按指定权限创建
// 主要目的是服务创建，普通用户有权访问，UAC 也没问题。
// securityDescriptor    "D:P(A;OICI;GWGR;;;SY)(A;OICI;GWGR;;;BA)(A;OICI;GWGR;;;IU)(A;OICI;GWGR;;;RC)"
//“D：P（A; OICI; GA ;;; SY）（A; OICI; GA ;;; BA）（A; OICI; GR ;;; IU）”指定DACL。D：P表示这是一个DACL（而不是SACL ......你很少使用SACL），后面是几个控制谁可以访问的ACE字符串。每一个都是A（允许）并允许对象并包含继承（OICI）。第一个授予系统（SY）和管理员（BA，内置管理员）的所有访问权限（GA - 全部授予）。最后一次授予读取（GR）给交互式用户（IU），这些用户实际登录到会话。
// https://blog.csdn.net/qinlicang/article/details/5538307
// https://stackoverflow.com/questions/898683/how-to-share-memory-between-services-and-user-processes
func CreateMmapWithSecurityDescriptor(name string, size uint32, write bool, securityDescriptor string) (*Mmap, error) {
	if size == 0 {
		return nil, fmt.Errorf("size==0")
	}

	cancel := false

	namePtr, _ := windows.UTF16PtrFromString(name)

	prot := uint32(syscall.PAGE_READONLY)
	if write {
		prot = syscall.PAGE_READWRITE
	}

	var security *windows.SecurityAttributes
	if len(securityDescriptor) != 0 {
		security = new(windows.SecurityAttributes)
		err := ConvertStringSecurityDescriptorToSecurityDescriptor(securityDescriptor, SDDL_REVISION_1, (*SecurityDescriptor)(unsafe.Pointer(&security.SecurityDescriptor)), nil)
		if err != nil {
			return nil, fmt.Errorf("ConvertStringSecurityDescriptorToSecurityDescriptor, %v", err)
		}
		defer func() {
			LocalFree(windows.Pointer(unsafe.Pointer(security.SecurityDescriptor)))
		}()
	}

	fileHandle, err := windows.CreateFileMapping(windows.Handle(INVALID_HANDLE_VALUE),
		security, prot, 0, uint32(size), namePtr)
	if err != nil {
		return nil, fmt.Errorf("CreateFileMapping, %v", err)
	}
	defer func() {
		if cancel {
			windows.CloseHandle(fileHandle)
		}
	}()

	access := uint32(windows.FILE_MAP_READ)
	if write {
		access = windows.FILE_MAP_WRITE
	}
	addr, err := windows.MapViewOfFile(fileHandle, access, 0,
		0, uintptr(size))
	if err != nil {
		cancel = true
		return nil, fmt.Errorf("MapViewOfFile, %v", err)
	}

	m := &Mmap{fileHandle: Handle(fileHandle), addr: addr, size: int(size)}

	runtime.SetFinalizer(m, (*Mmap).Close)

	return m, nil
}

// 打开无文件命名内存共享
// 虽然 MapViewOfFile 函数允许为0 ，但是当 size==0 时无法转换为 []byte ，所以不支持 size == 0。
func OpenMmap(name string, size uint32, write bool) (*Mmap, error) {
	if size == 0 {
		return nil, fmt.Errorf("size==0")
	}

	cancel := false

	access := uint32(windows.FILE_MAP_READ)
	if write {
		access = windows.FILE_MAP_WRITE
	}
	fileHandle, err := OpenFileMapping(DWord(access), false, name)
	if err != nil {
		return nil, fmt.Errorf("OpenFileMapping, %v", err)
	}
	defer func() {
		if cancel {
			windows.CloseHandle(windows.Handle(fileHandle))
		}
	}()

	addr, err := windows.MapViewOfFile(windows.Handle(fileHandle), uint32(access), 0,
		0, uintptr(size))
	if err != nil {
		cancel = true
		return nil, fmt.Errorf("MapViewOfFile, %v", err)
	}

	m := &Mmap{fileHandle: fileHandle, addr: addr, size: int(size)}

	runtime.SetFinalizer(m, (*Mmap).Close)

	return m, nil
}

// 关闭、释放
func (m *Mmap) Close() error {
	if m.addr != uintptr(0) {
		err := windows.UnmapViewOfFile(uintptr(m.addr))
		if err != nil {
			return fmt.Errorf("UnmapViewOfFile, %v", err)
		}
		m.size = 0
		m.addr = uintptr(0)
	}

	if m.fileHandle != 0 {
		err := windows.CloseHandle(windows.Handle(m.fileHandle))
		if err != nil {
			return fmt.Errorf("CloseHandle, %v", err)
		}
		m.size = 0
		m.fileHandle = 0
	}

	runtime.KeepAlive(m)

	// 测试确保 runtime.SetFinalizer 正常工作
	// fmt.Println("mmap.close()")
	return nil
}
