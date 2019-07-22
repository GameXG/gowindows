package gowindows

import (
	"syscall"
	"unicode/utf16"
	"unsafe"

	"fmt"

	"golang.org/x/sys/windows"
)

var (
	kernel32            = windows.NewLazyDLL("kernel32.dll")
	readProcessMemory   = kernel32.NewProc("ReadProcessMemory")
	isWow64Process      = kernel32.NewProc("IsWow64Process")
	createMutexW        = kernel32.NewProc("CreateMutexW")
	openMutexW          = kernel32.NewProc("OpenMutexW")
	releaseMutex        = kernel32.NewProc("ReleaseMutex")
	localFree           = kernel32.NewProc("LocalFree")
	openFileMappingW    = kernel32.NewProc("OpenFileMappingW")
	getModuleFileName   = kernel32.NewProc("GetModuleFileNameW")
	wideCharToMultiByte = kernel32.NewProc("WideCharToMultiByte")
)

//BOOL WINAPI ReadProcessMemory(
//_In_  HANDLE  hProcess,
//_In_  LPCVOID lpBaseAddress,
//_Out_ LPVOID  lpBuffer,
//_In_  SIZE_T  nSize,
//_Out_ SIZE_T  *lpNumberOfBytesRead
//);
func ReadProcessMemory(processHandle Handle, baseAddress uint,
	bufferData windows.Pointer, bufferSize uint, returnSize *uint) error {

	r1, _, err := readProcessMemory.Call(uintptr(processHandle), uintptr(baseAddress),
		uintptr(unsafe.Pointer(bufferData)), uintptr(bufferSize), uintptr(unsafe.Pointer(returnSize)))
	if r1 == 0 {
		if err != ERROR_SUCCESS {
			return err
		} else {
			return syscall.EINVAL
		}
		return err
	}
	return nil
}

// BOOL IsWow64Process(
// HANDLE hProcess,
// PBOOL  Wow64Process
// );
func IsWow64Process(hProcess Handle) (bool, error) {
	// 不存在 IsWow64Process 函数的操作系统绝对不是64位系统，那么不会是 wow64进程
	if isWow64Process.Find() != nil {
		return false, nil
	}

	is64 := 0

	r1, _, err := isWow64Process.Call(uintptr(hProcess), uintptr(unsafe.Pointer(&is64)))
	if r1 == 0 {
		return false, err
	}

	switch is64 {
	case 0:
		return false, nil
	default:
		return true, nil
	}
}

// 是否是64位系统
func Is64System() (bool, error) {
	if ptrSize == 8 {
		// 当前程序是 64 位，系统就是64位
		return true, nil
	}

	h, err := windows.GetCurrentProcess()
	if err != nil {
		return false, fmt.Errorf("windows.GetCurrentProcess, %v", err)
	}
	defer windows.CloseHandle(h)

	is64, err := IsWow64Process(Handle(h))
	if err != nil {
		return false, fmt.Errorf("IsWow64Process, %v", err)
	}

	return is64, nil
}

// 释放指定的本地内存对象并使其句柄无效。
// https://docs.microsoft.com/zh-cn/windows/desktop/api/winbase/nf-winbase-localfree
// 参数：
// h     本地内存对象的句柄。此句柄由 LocalAlloc或 LocalReAlloc函数返回。释放使用GlobalAlloc分配的内存是不安全的。
//HLOCAL
//WINAPI
//LocalFree(
//    _Frees_ptr_opt_ HLOCAL hMem
//    );
// typedef HANDLE              HLOCAL;
func LocalFree(h windows.Pointer) error {
	r1, _, e1 := localFree.Call(uintptr(unsafe.Pointer(h)))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

//https://docs.microsoft.com/en-us/windows/desktop/api/winbase/nf-winbase-openfilemappinga
//HANDLE OpenFileMappingA(
//  DWORD  dwDesiredAccess,
//  BOOL   bInheritHandle,
//  LPCSTR lpName
//);
// windows.FILE_MAP_COPY    = 0x01
// windows.FILE_MAP_WRITE   = 0x02
// windows.FILE_MAP_READ    = 0x04
// windows.FILE_MAP_EXECUTE = 0x20
func OpenFileMapping(desiredAccess DWord, inheritHandle bool, name string) (Handle, error) {
	var _inheritHandle Bool
	switch inheritHandle {
	case true:
		_inheritHandle = 1
	case false:
		_inheritHandle = 0
	}

	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}

	r1, _, e1 := openFileMappingW.Call(uintptr(desiredAccess), uintptr(_inheritHandle), uintptr(unsafe.Pointer(_name)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}

	return Handle(r1), nil
}

func GetModuleFileName(module HMODULE) (string, error) {
	size := DWord(syscall.MAX_PATH * 2)

	for i := 0; i < 10; i++ {
		buf := make([]uint16, size)
		r1, _, e1 := getModuleFileName.Call(uintptr(module),
			uintptr(unsafe.Pointer(&buf[0])), uintptr(size))

		// 返回0表示函数失败
		if r1 == 0 {
			if e1 != ERROR_SUCCESS {
				return "", e1
			} else {
				return "", syscall.EINVAL
			}
		}

		// 可能空间不够，再次执行
		if DWord(r1) >= size {
			size = size * 2
			continue
		}

		return syscall.UTF16ToString(buf), nil
	}
	return "", fmt.Errorf("多次扩展空间还是空间不足。")
}

//TODO:下面的需要仔细检查

//第二个参数，TRUE表示建立占有Mutex使用权，FALSE表示建立者不占有Mutex使用权
func CreateMutex(lpSecurityAttributes uintptr, bInitalOwner bool, name string) (handle Handle, err error) {
	mutexName, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}
	var _p0 uint32
	if bInitalOwner {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall(createMutexW.Addr(), 3, uintptr(lpSecurityAttributes), uintptr(_p0), uintptr(unsafe.Pointer(mutexName)))
	handle = Handle(r0)
	if handle == InvalidHandle {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// 打开现有的已命名互斥对象。
// dwDesiredAccess	对互斥锁对象的访问。使用互斥锁只需要SYNCHRONIZE访问权限; 要更改互斥锁的安全性，请指定MUTEX_ALL_ACCESS。如果指定对象的安全描述符不允许对调用进程请求访问，则该函数将失败。有关访问权限的列表，请参阅 [同步对象安全性和访问权限](https://docs.microsoft.com/zh-cn/windows/win32/sync/synchronization-object-security-and-access-rights)。
// bInitalOwner		如果此值为TRUE，则此进程创建的进程将继承该句柄。否则，进程不会继承此句柄。
// name				要打开的互斥锁的名称。名称比较区分大小写。
// https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-openmutexw
// https://msdn.microsoft.com/en-us/windows/desktop/ms684315
func OpenMutex(dwDesiredAccess DWord, bInitalOwner bool, name string) (handle Handle, err error) {
	mutexName, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}
	var _p0 uint32
	if bInitalOwner {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall(openMutexW.Addr(), 3, uintptr(dwDesiredAccess), uintptr(_p0), uintptr(unsafe.Pointer(mutexName)))
	handle = Handle(r0)
	if handle == InvalidHandle {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ReleaseMutex(handle Handle) (err error) {
	releaseMutex.Call(uintptr(handle))
	return nil
}

//     syscall.WaitForSingleObject(mu,syscall.INFINITE)

//只运行一次返回 nil ，如果已经存在运行中的返回 err
func One(name string) (handle Handle, rerr error) {
	h, err := CreateMutex(0, false, name)
	if err != nil {
		return 0, fmt.Errorf("创建信号量失败，%v", err)
	}
	lasterr := syscall.GetLastError()
	if lasterr == syscall.ERROR_ALREADY_EXISTS {
		syscall.CloseHandle(Handle(h))
		return 0, fmt.Errorf("已经运行了一个实例。")
	}

	event, err := syscall.WaitForSingleObject(Handle(h), 0)
	// WAIT_OBJECT_0   	mutex 未被持有或安全释放
	// WAIT_ABANDONED	上一持有 mutex 的线程或进程异常终止
	if err == nil && (event == syscall.WAIT_OBJECT_0 || event == syscall.WAIT_ABANDONED) {
		// 持有信号量成功
		return h, nil
	}

	syscall.CloseHandle(Handle(h))
	return 0, fmt.Errorf("信号量已经被其他程序持有。")
}

func OneDefer(h Handle) {
	ReleaseMutex(h)
	syscall.CloseHandle(Handle(h))
}

// https://docs.microsoft.com/en-us/windows/desktop/api/stringapiset/nf-stringapiset-widechartomultibyte
// int WideCharToMultiByte(
//  UINT                               CodePage,
//  DWORD                              dwFlags,
//  _In_NLS_string_(cchWideChar)LPCWCH lpWideCharStr,
//  int                                cchWideChar,
//  LPSTR                              lpMultiByteStr,
//  int                                cbMultiByte,
//  LPCCH                              lpDefaultChar,
//  LPBOOL                             lpUsedDefaultChar
//);
func WideCharToMultiByte(codePage uint32, dwFlags DWord, wchar *uint16, nwchar int32, str *byte, nstr int32, lpDefaultChar *byte, pfUsedDefaultChar *bool) (nwrite int32, err error) {
	r1, _, e1 := wideCharToMultiByte.Call(
		uintptr(codePage),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(wchar)),
		uintptr(nwchar),
		uintptr(unsafe.Pointer(str)),
		uintptr(nstr),
		uintptr(unsafe.Pointer(lpDefaultChar)),
		uintptr(unsafe.Pointer(pfUsedDefaultChar)),
	)

	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}

	return int32(r1), nil
}

type CodePage uint32

const (
	//  Code Page Default Values.
	//
	CP_ACP        CodePage = 0  // default to ANSI code page
	CP_OEMCP      CodePage = 1  // default to OEM  code page
	CP_MACCP      CodePage = 2  // default to MAC  code page
	CP_THREAD_ACP CodePage = 3  // current thread's ANSI code page
	CP_SYMBOL     CodePage = 42 // SYMBOL translations

	CP_UTF7 CodePage = 65000 // UTF-7 translation
	CP_UTF8 CodePage = 65001 // UTF-8 translation
)

type WideChar2MultiByteFlags DWord

const (
	WC_COMPOSITECHECK WideChar2MultiByteFlags = 0x00000200 // convert composite to precomposed
	WC_DISCARDNS      WideChar2MultiByteFlags = 0x00000010 // discard non-spacing chars
	WC_SEPCHARS       WideChar2MultiByteFlags = 0x00000020 // generate separate chars
	WC_DEFAULTCHAR    WideChar2MultiByteFlags = 0x00000040 // replace w/ default char
	//#if (WINVER >= 0x0600)
	// 	Windows Vista及更高版本：如果遇到无效的输入字符，则失败（通过返回0并将最后错误代码设置为ERROR_NO_UNICODE_TRANSLATION）。您可以通过调用GetLastError来检索最后一个错误代码。如果未设置此标志，则该函数用U + FFFD替换非法序列（根据指定的代码页编码），并通过返回转换后的字符串的长度来成功。请注意，这个标志时才适用代码页被指定为CP_UTF8或54936.它不能与其他代码页值来使用。
	WC_ERR_INVALID_CHARS WideChar2MultiByteFlags = 0x00000080 // error for invalid chars
	//#endif

	//#if(WINVER >= 0x0500)
	// 把不能直接转换成相应多字节字符的Unicode字符转换成lpDefaultChar指定的默认字符。也就是说，如果把Unicode转换成多字节字符，然后再转换回来，你并不一定得到相同的Unicode字符，因为这期间可能使用了默认字符。此选项可以单独使用，也可以和其他选项一起使用。
	// 对于需要验证的字符串，例如文件，资源和用户名，应用程序应始终使用WC_NO_BEST_FIT_CHARS标志。此标志阻止函数将字符映射到看似相似但具有非常不同语义的字符。在某些情况下，语义变化可能是极端的。例如，“∞”（无穷大）的符号在某些代码页中映射到8（8）。
	WC_NO_BEST_FIT_CHARS WideChar2MultiByteFlags = 0x00000400 // do not use best fit chars
//#endif /* WINVER >= 0x0500 */
)

// wchar(UTF-16LE) 转换到 char
// 输入如果不包含 \0 ，输出也将不包含。
func WideChar2MultiByte(codePage CodePage, dwFlags WideChar2MultiByteFlags, wchar []uint16, lpDefaultChar *byte, pfUsedDefaultChar *bool) ([]byte, error) {
	if len(wchar) == 0 {
		return nil, nil
	}

	n, err := WideCharToMultiByte(uint32(codePage), DWord(dwFlags), &wchar[0], int32(len(wchar)), nil, 0, lpDefaultChar, pfUsedDefaultChar)
	if err != nil {
		return nil, err
	}

	out := make([]byte, n)

	n, err = WideCharToMultiByte(uint32(codePage), DWord(dwFlags), &wchar[0], int32(len(wchar)), &out[0], int32(len(out)), lpDefaultChar, pfUsedDefaultChar)
	if err != nil {
		return nil, err
	}

	return out[:n], nil
}

type MultiByteToWideCharFlags DWord

const (
	//
	//  MBCS and Unicode Translation Flags.
	//

	// 默认; 不要与MB_COMPOSITE一起使用。始终使用预组合字符，即具有单个字符值的字符作为基本字符或非间距字符组合。例如，在角色è中，e是基本角色，重音符号是非滑动角色。如果为字符定义了单个Unicode代码点，则应用程序应使用它而不是单独的基本字符和非间距字符。例如，Ä由单个Unicode代码点LATIN CAPITAL LETTER A WITH DIAERESIS（U + 00C4）表示。
	MB_PRECOMPOSED MultiByteToWideCharFlags = 0x00000001 // use precomposed chars

	// 始终使用分解的字符，即基本字符和一个或多个非间距字符各自具有不同代码点值的字符。例如，Ä由A +¨表示：LATIN CAPITAL LETTER A（U + 0041）+ COMBINING DIAERESIS（U + 0308）。请注意，此标志不能与MB_PRECOMPOSED一起使用。
	MB_COMPOSITE MultiByteToWideCharFlags = 0x00000002 // use composite chars

	// 使用字形字符而不是控制字符。
	MB_USEGLYPHCHARS MultiByteToWideCharFlags = 0x00000004 // use glyph chars, not ctrl chars

	// 如果遇到无效的输入字符，则失败。
	//从Windows Vista开始，如果应用程序未设置此标志，则该函数不会删除非法代码点，而是使用U + FFFD替换非法序列（根据指定的代码页编码）。
	//
	//带有SP4及更高版本的Windows 2000，Windows XP：  如果未设置此标志，则该功能将以静默方式删除非法代码点。对GetLastError的调用返回ERROR_NO_UNICODE_TRANSLATION。
	MB_ERR_INVALID_CHARS MultiByteToWideCharFlags = 0x00000008 // error for invalid chars
)

//  char 转换到 wchar(UTF-16LE)
// 输入如果不包含 \0 ，输出也将不包含。
// https://docs.microsoft.com/en-us/windows/desktop/api/stringapiset/nf-stringapiset-multibytetowidechar
func MultiByte2WideChar(codePage CodePage, dwFlags MultiByteToWideCharFlags, str []byte) ([]uint16, error) {
	if len(str) == 0 {
		return nil, nil
	}

	n, err := windows.MultiByteToWideChar(uint32(codePage), uint32(dwFlags), &str[0], int32(len(str)), nil, 0)
	if err != nil {
		return nil, err
	}

	out := make([]uint16, n)
	n, err = windows.MultiByteToWideChar(uint32(codePage), uint32(dwFlags), &str[0], int32(len(str)), &out[0], int32(len(out)))
	if err != nil {
		return nil, err
	}

	return out[:n], nil
}

// 内部使用 windows api 来实现
// 主要目的其实是 go string 转换为 windows 本地编码
func Str2MultiByte(codePage CodePage, dwFlags WideChar2MultiByteFlags, str string, lpDefaultChar *byte, pfUsedDefaultChar *bool) ([]byte, error) {
	utf16Data := utf16.Encode([]rune(str))

	return WideChar2MultiByte(codePage, dwFlags, utf16Data, lpDefaultChar, pfUsedDefaultChar)
}

// 内部使用 windows api 来实现
// 主要目的其实是 windows 本地编码 转换为 go string
func MultiByte2Str(codePage CodePage, dwFlags MultiByteToWideCharFlags, str []byte) (string, error) {
	utf16Data, err := MultiByte2WideChar(codePage, dwFlags, str)
	if err != nil {
		return "", err
	}

	return string(utf16.Decode(utf16Data)), nil
}

var WaitForSingleObject = windows.WaitForSingleObject
