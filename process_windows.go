package gowindows

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetProcessParameters(processHandle Handle) (string, string, error) {
	is64system, err := Is64System()
	if err != nil {
		return "", "", fmt.Errorf("Is64System,%v", err)
	}

	if is64system == true && ptrSize == 4 {
		// wow64 环境，复杂些
		pInfo := PROCESS_BASIC_INFORMATION64{}
		size := uint32(unsafe.Sizeof(pInfo))

		err = NtWow64QueryInformationProcess64(processHandle, ProcessBasicInformation,
			windows.Pointer(unsafe.Pointer(&pInfo)), size, nil)
		if err != nil {
			return "", "", fmt.Errorf("NtWow64QueryInformationProcess64, %v", err)
		}

		peb := PEB64{}
		err = NtWow64ReadVirtualMemory64(processHandle, pInfo.PebBaseAddress, windows.Pointer(unsafe.Pointer(&peb)), uint64(unsafe.Sizeof(peb)), nil)
		if err != nil {
			return "", "", fmt.Errorf("NtWow64ReadVirtualMemory64(peb), %v", err)
		}

		process_parameters := RTL_USER_PROCESS_PARAMETERS64{}
		err = NtWow64ReadVirtualMemory64(processHandle, peb.ProcessParameters, windows.Pointer(unsafe.Pointer(&process_parameters)), uint64(unsafe.Sizeof(process_parameters)), nil)
		if err != nil {
			return "", "", fmt.Errorf("NtWow64ReadVirtualMemory64(process_parameters), %v", err)
		}

		imagePathName := make([]uint16, process_parameters.ImagePathName.Length)
		commandLine := make([]uint16, process_parameters.CommandLine.Length)

		err = NtWow64ReadVirtualMemory64(processHandle, process_parameters.ImagePathName.Buffer, windows.Pointer(unsafe.Pointer(&imagePathName[0])), uint64(len(imagePathName)), nil)
		if err != nil {
			return "", "", fmt.Errorf("NtWow64ReadVirtualMemory64(imagePathName), %v", err)
		}
		fmt.Println("imagePathName:", windows.UTF16ToString(imagePathName))
		err = NtWow64ReadVirtualMemory64(processHandle, process_parameters.CommandLine.Buffer, windows.Pointer(unsafe.Pointer(&commandLine[0])), uint64(len(commandLine)), nil)
		if err != nil {
			return "", "", fmt.Errorf("NtWow64ReadVirtualMemory64(commandLine), %v", err)
		}

		return windows.UTF16ToString(imagePathName), windows.UTF16ToString(commandLine), nil

	} else {
		// 64程序@64位系统 或 32位程序@32位系统

		pInfo := PROCESS_BASIC_INFORMATION{}
		size := uint32(unsafe.Sizeof(pInfo))

		err = NtQueryInformationProcess(processHandle, ProcessBasicInformation,
			windows.Pointer(unsafe.Pointer(&pInfo)), size, nil)
		if err != nil {
			return "", "", fmt.Errorf("NtQueryInformationProcess, %v", err)
		}

		peb := PEB{}
		err = ReadProcessMemory(processHandle, pInfo.PebBaseAddress, windows.Pointer(unsafe.Pointer(&peb)), uint(unsafe.Sizeof(peb)), nil)
		if err != nil {
			return "", "", fmt.Errorf("ReadProcessMemory(peb), %v", err)
		}

		process_parameters := RTL_USER_PROCESS_PARAMETERS{}
		err = ReadProcessMemory(processHandle, peb.ProcessParameters, windows.Pointer(unsafe.Pointer(&process_parameters)), uint(unsafe.Sizeof(process_parameters)), nil)
		if err != nil {
			return "", "", fmt.Errorf("ReadProcessMemory(process_parameters), %v", err)
		}

		imagePathName := make([]uint16, process_parameters.ImagePathName.Length)
		commandLine := make([]uint16, process_parameters.CommandLine.Length)

		err = ReadProcessMemory(processHandle, process_parameters.ImagePathName.Buffer, windows.Pointer(unsafe.Pointer(&imagePathName[0])), uint(process_parameters.ImagePathName.Length), nil)
		if err != nil {
			return "", "", fmt.Errorf("ReadProcessMemory(imagePathName), %v", err)
		}
		err = ReadProcessMemory(processHandle, process_parameters.CommandLine.Buffer, windows.Pointer(unsafe.Pointer(&commandLine[0])), uint(process_parameters.CommandLine.Length), nil)
		if err != nil {
			return "", "", fmt.Errorf("ReadProcessMemory(commandLine), %v", err)
		}

		return windows.UTF16ToString(imagePathName), windows.UTF16ToString(commandLine), nil
	}
}

func GetProcessParametersWPid(pid uint32) (string, string, error) {
	processHandle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|PROCESS_VM_READ, false, uint32(pid))
	if err != nil {
		return "", "", fmt.Errorf("windows.OpenProcess, %v", err)
	}
	defer windows.CloseHandle(processHandle)

	return GetProcessParameters(Handle(processHandle))
}

func EscapeArg(s string) string {
	if len(s) == 0 {
		return "\"\""
	}
	n := len(s)
	hasSpace := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '"', '\\':
			n++
		case ' ', '\t':
			hasSpace = true
		}
	}
	if hasSpace {
		n += 2
	}
	if n == len(s) {
		return s
	}

	qs := make([]byte, n)
	j := 0
	if hasSpace {
		qs[j] = '"'
		j++
	}
	slashes := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		default:
			slashes = 0
			qs[j] = s[i]
		case '\\':
			slashes++
			qs[j] = s[i]
		case '"':
			for ; slashes > 0; slashes-- {
				qs[j] = '\\'
				j++
			}
			qs[j] = '\\'
			j++
			qs[j] = s[i]
		}
		j++
	}
	if hasSpace {
		for ; slashes > 0; slashes-- {
			qs[j] = '\\'
			j++
		}
		qs[j] = '"'
		j++
	}
	return string(qs[:j])
}

// makeCmdLine builds a command line out of args by escaping "special"
// characters and joining the arguments with spaces.
func makeCmdLine(args []string) string {
	var s string
	for _, v := range args {
		if s != "" {
			s += " "
		}
		s += EscapeArg(v)
	}
	return s
}

// 另一个windows 启动新进程的实现
// 创建原因是标准库的实现返回错误：
// 注意：这个函数未释放 ProcessInformation.Process ，需要调用者去释放它。
func MyCreateProcess(name string, hide bool, arg ...string) (*windows.ProcessInformation, error) {
	args := makeCmdLine(append([]string{name}, arg...))

	nameUtf16, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return nil, fmt.Errorf("windows.UTF16PtrFromString, %v", err)
	}

	argsUtf16, err := windows.UTF16PtrFromString(args)
	if err != nil {
		return nil, fmt.Errorf("windows.UTF16PtrFromString, %v", err)
	}

	si := new(windows.StartupInfo)
	si.Cb = uint32(unsafe.Sizeof(*si))
	//si.Flags = windows.STARTF_USESTDHANDLES  // 注释掉，不使用标准输入输出

	flags := uint32(windows.CREATE_UNICODE_ENVIRONMENT)

	if hide {
		si.Flags |= windows.STARTF_USESHOWWINDOW
		si.ShowWindow = windows.SW_HIDE
		flags = flags | 0x08000000
	}

	info := new(windows.ProcessInformation)

	err = windows.CreateProcess(nameUtf16, argsUtf16, nil, nil, false, flags, nil, nil, si, info)
	if err != nil {
		return nil, fmt.Errorf("windows.CreateProcess, %v", err)
	}

	// runtime.SetFinalizer(info, ProcessInformationRelease)

	return info, nil
}

func ProcessInformationRelease(info *windows.ProcessInformation) {
	if info != nil || info.Process != 0 {
		windows.CloseHandle(info.Process)
	}
}
