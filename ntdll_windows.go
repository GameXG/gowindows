package gowindows

import (
	"fmt"

	"unsafe"

	"syscall"

	"golang.org/x/sys/windows"
)

const STATUS_SUCCESS = 0

var (
	ntdll                            *windows.DLL
	ntWow64QueryInformationProcess64 *windows.Proc
	ntWow64ReadVirtualMemory64       *windows.Proc
	ntQueryInformationProcess        *windows.Proc
	ntReadVirtualMemory              *windows.Proc
	ntRtlGetVersion                  *windows.Proc
)

func init() {
	var err error
	ntdll, err = windows.LoadDLL("ntdll.dll")
	if err == nil {
		// 这几个需要判断是否存在，所以无法使用延迟加载
		ntWow64QueryInformationProcess64, _ = ntdll.FindProc("NtWow64QueryInformationProcess64")
		ntWow64ReadVirtualMemory64, _ = ntdll.FindProc("NtWow64ReadVirtualMemory64")
		ntQueryInformationProcess, _ = ntdll.FindProc("NtQueryInformationProcess")
		ntReadVirtualMemory, _ = ntdll.FindProc("NtReadVirtualMemory")

		ntRtlGetVersion, _ = ntdll.FindProc("RtlGetVersion")
	}
}

// 只能用在 32位程序64位系统下
func NtWow64QueryInformationProcess64(processHandle Handle, processInformationClass int32,
	processInformation windows.Pointer, processInformationLength uint32, returnLength *uint32) error {

	if ntWow64QueryInformationProcess64 == nil {
		return fmt.Errorf("ntWow64QueryInformationProcess64==nil")
	}

	r1, _, err := ntWow64QueryInformationProcess64.Call(uintptr(processHandle), uintptr(processInformationClass),
		uintptr(unsafe.Pointer(processInformation)), uintptr(processInformationLength),
		uintptr(unsafe.Pointer(returnLength)))

	if int(r1) < 0 {
		// 函数内部如果发现空间不匹配 r1 会等于 STATUS_INFO_LENGTH_MISMATCH(0xC0000004),
		// 指示空间不匹配。
		if err != ERROR_SUCCESS {
			return err
		} else {
			return syscall.EINVAL
		}
	}

	return nil
}

//__kernel_entry NTSTATUS NtQueryInformationProcess(
//IN HANDLE               ProcessHandle,
//IN PROCESSINFOCLASS     ProcessInformationClass,
//OUT PVOID               ProcessInformation,
//IN ULONG                ProcessInformationLength,
//OUT PULONG ReturnLength OPTIONAL
//);
func NtQueryInformationProcess(processHandle Handle, processInformationClass int32,
	processInformation windows.Pointer, processInformationLength uint32, returnLength *uint32) error {

	if ntQueryInformationProcess == nil {
		return fmt.Errorf("ntQueryInformationProcess==nil")
	}

	r1, _, err := ntQueryInformationProcess.Call(uintptr(processHandle), uintptr(processInformationClass),
		uintptr(unsafe.Pointer(processInformation)), uintptr(processInformationLength),
		uintptr(unsafe.Pointer(returnLength)))

	if int(r1) < 0 {
		// 函数内部如果发现空间不匹配 r1 会等于 STATUS_INFO_LENGTH_MISMATCH(0xC0000004),
		// 指示空间不匹配。

		if err != ERROR_SUCCESS {
			return err
		} else {
			return syscall.EINVAL
		}
	}

	return nil
}

//NTSTATUS (__stdcall *NtWow64ReadVirtualMemory64)(
//HANDLE ProcessHandle,
//PVOID64 BaseAddress,
//PVOID Buffer,
//ULONGLONG BufferSize,
//PULONGLONG NumberOfBytesRead
//);
func NtWow64ReadVirtualMemory64(processHandle Handle, baseAddress uint64,
	bufferData windows.Pointer, bufferSize uint64, returnSize *uint64) error {

	if ntWow64ReadVirtualMemory64 == nil {
		return fmt.Errorf("ntWow64ReadVirtualMemory64==nil")
	}

	var r1 uintptr
	var err error

	if ptrSize == 8 {
		// 64位程序，虽然理论不应该有这种情况
		r1, _, err = ntWow64ReadVirtualMemory64.Call(uintptr(processHandle), uintptr(baseAddress),
			uintptr(unsafe.Pointer(bufferData)), uintptr(bufferSize), uintptr(unsafe.Pointer(returnSize)))
	} else {
		// 32 位程序
		r1, _, err = ntWow64ReadVirtualMemory64.Call(uintptr(processHandle),
			uintptr(baseAddress), uintptr(baseAddress>>32), uintptr(unsafe.Pointer(bufferData)),
			uintptr(bufferSize), uintptr(bufferSize>>32),
			uintptr(unsafe.Pointer(returnSize)))
	}

	if int(r1) < 0 {
		// 如果读操作跨越不可访问的区域将失败。

		if err != ERROR_SUCCESS {
			return err
		} else {
			return syscall.EINVAL
		}
	}

	return nil
}

func ReadVirtualMemory(processHandle Handle, baseAddress uint,
	bufferData windows.Pointer, bufferSize uint, returnSize *uint) error {

	if ntReadVirtualMemory == nil {
		return fmt.Errorf("ntReadVirtualMemory==nil")
	}

	var r1 uintptr
	var err error

	// 64位程序，虽然理论不应该有这种情况
	r1, _, err = ntReadVirtualMemory.Call(uintptr(processHandle), uintptr(baseAddress),
		uintptr(unsafe.Pointer(bufferData)), uintptr(bufferSize), uintptr(unsafe.Pointer(returnSize)))

	if int(r1) < 0 {
		// 如果读操作跨越不可访问的区域将失败。

		if err != ERROR_SUCCESS {
			return err
		} else {
			return syscall.EINVAL
		}
	}

	return nil
}

// https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/content/wdm/ns-wdm-_osversioninfow
// 版本号参考: https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/content/wdm/ns-wdm-_osversioninfoexw
//typedef struct _OSVERSIONINFOW {
//  ULONG dwOSVersionInfoSize;
//  ULONG dwMajorVersion;
//  ULONG dwMinorVersion;
//  ULONG dwBuildNumber;
//  ULONG dwPlatformId;
//  WCHAR szCSDVersion[128];
//} OSVERSIONINFOW, *POSVERSIONINFOW, *LPOSVERSIONINFOW, RTL_OSVERSIONINFOW, *PRTL_OSVERSIONINFOW;
type OsVsersionInfow struct {
	OSVersionInfoSize ULong
	MajorVersion      ULong
	MinorVersion      ULong
	BuildNumber       ULong
	PlatformId        ULong
	CSDVersion        [128]WCHAR
}

func (v *OsVsersionInfow) GetCSDVersion() string {
	return windows.UTF16ToString(v.CSDVersion[:])
}

func (v *OsVsersionInfow) IsWindows2000OrGreater() bool {
	// win2k
	// 5.0
	if v.MajorVersion >= 5 {
		return true
	}
	return false
}
func (v *OsVsersionInfow) IsWindowsXPOrGreater() bool {
	// winxp
	// 5.1
	if v.MajorVersion > 5 || (v.MajorVersion == 5 && v.MinorVersion >= 1) {
		return true
	}
	return false
}
func (v *OsVsersionInfow) IsWindowsVistaOrGreater() bool {
	// win vista
	// 6.0
	if v.MajorVersion >= 6 {
		return true
	}
	return false
}
func (v *OsVsersionInfow) IsWindows7OrGreater() bool {
	// win 7
	// 6.1
	if v.MajorVersion > 6 || (v.MajorVersion == 6 && v.MinorVersion >= 1) {
		return true
	}
	return false
}
func (v *OsVsersionInfow) IsWindows8OrGreater() bool {
	// win 8
	// 6.2
	if v.MajorVersion > 6 || (v.MajorVersion == 6 && v.MinorVersion >= 2) {
		return true
	}
	return false
}
func (v *OsVsersionInfow) IsWindows8Point1OrGreater() bool {
	// win 8.1
	// 6.3
	if v.MajorVersion > 6 || (v.MajorVersion == 6 && v.MinorVersion >= 3) {
		return true
	}
	return false
}
func (v *OsVsersionInfow) IsWindows10OrGreater() bool {
	// win 10
	// 10
	if v.MajorVersion >= 10 {
		return true
	}
	return false
}

func (v *OsVsersionInfow) GetString() string {
	return fmt.Sprintf("%v.%v", v.MajorVersion, v.MinorVersion)
}

// 这个函数可以获得 win8 之后正确的版本号
// 但是 Windows2003之前 RtlGetVersion 的行为在兼容模式下和 GetVersionExW 不一致，
// Vista 之后在兼容模式下它的行为和 GetVersionExW 一致返回的是兼容的目标版本的系统版本。
// Windows系统版本判定那些事儿 https://blog.csdn.net/magictong/article/details/40753519
// https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/content/wdm/nf-wdm-rtlgetversion
func RtlGetVersion() (*OsVsersionInfow, error) {
	if ntRtlGetVersion == nil {
		return nil, fmt.Errorf("ntRtlGetVersion==nil")
	}

	version := new(OsVsersionInfow)
	version.OSVersionInfoSize = ULong(unsafe.Sizeof(OsVsersionInfow{}))

	r1, _, e1 := ntRtlGetVersion.Call(uintptr(unsafe.Pointer(version)))
	if r1 != STATUS_SUCCESS {
		if e1 != ERROR_SUCCESS {
			return nil, e1
		} else {
			return nil, syscall.EINVAL
		}
	}
	return version, nil
}
