package gowindows

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// https://support.microsoft.com/en-us/help/131065/how-to-obtain-a-handle-to-any-process-with-sedebugprivilege
func SetPrivilege(hToken windows.Token, privilege Privilege, bEnablePrivilege bool) error {
	tp := TOKEN_PRIVILEGES{}
	luid := LUID{}
	tpPrevious := TOKEN_PRIVILEGES{}
	cbPrevious := uint32(unsafe.Sizeof(TOKEN_PRIVILEGES{}))

	err := LookupPrivilegeValue(nil, windows.StringToUTF16Ptr(string(privilege)), &luid)
	if err != nil {
		return fmt.Errorf("LookupPrivilegeValue, %v", err)
	}

	// 第一步，获取当前权限设置
	tp.PrivilegeCount = 1
	tp.Privileges[0].Luid = luid
	tp.Privileges[0].Attributes = 0

	err = AdjustTokenPrivileges(hToken, false, &tp, uint32(unsafe.Sizeof(tp)), &tpPrevious, &cbPrevious)
	if err != nil {
		return fmt.Errorf("AdjustTokenPrivileges1, %v", err)
	}

	// 第二步，根据原来的权限设置权限
	tpPrevious.PrivilegeCount = 1
	tpPrevious.Privileges[0].Luid = luid

	if bEnablePrivilege {
		tpPrevious.Privileges[0].Attributes |= SE_PRIVILEGE_ENABLED
	} else {
		tpPrevious.Privileges[0].Attributes ^= (SE_PRIVILEGE_ENABLED &
			tpPrevious.Privileges[0].Attributes)
	}

	err = AdjustTokenPrivileges(hToken, false, &tpPrevious, cbPrevious, nil, nil)
	if err != nil {
		return fmt.Errorf("AdjustTokenPrivileges2 ,%v", err)
	}
	return nil
}

// https://support.microsoft.com/en-us/help/131065/how-to-obtain-a-handle-to-any-process-with-sedebugprivilege
func SetSelfProcessPrivilege(privilege Privilege, bEnablePrivilege bool) error {
	pHandle, err := windows.GetCurrentProcess()
	if err != nil {
		return fmt.Errorf("windows.GetCurrentProcess, %v", err)
	}
	defer windows.CloseHandle(pHandle)

	hToken := windows.Token(0)

	err = windows.OpenProcessToken(pHandle, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &hToken)
	if err != nil {
		return fmt.Errorf("windows.OpenProcessToken, %v", err)
	}
	defer windows.CloseHandle(windows.Handle(hToken))

	err = SetPrivilege(hToken, privilege, bEnablePrivilege)
	if err != nil {
		return fmt.Errorf("SetPrivilege, %v", err)
	}

	return nil
}

//#define RTN_OK 0
//#define RTN_USAGE 1
//#define RTN_ERROR 13
//
//#include <windows.h>
//#include <stdio.h>
//
//BOOL SetPrivilege(
//HANDLE hToken,          // token handle
//LPCTSTR Privilege,      // Privilege to enable/disable
//BOOL bEnablePrivilege   // TRUE to enable.  FALSE to disable
//);
//
//void DisplayError(LPTSTR szAPI);
//
//int main(int argc, char *argv[])
//{
//HANDLE hProcess;
//HANDLE hToken;
//int dwRetVal=RTN_OK; // assume success from main()
//
//// show correct usage for kill
//if (argc != 2)
//{
//fprintf(stderr,"Usage: %s [ProcessId]\n", argv[0]);
//return RTN_USAGE;
//}
//
//if(!OpenThreadToken(GetCurrentThread(), TOKEN_ADJUST_PRIVILEGES | TOKEN_QUERY, FALSE, &hToken))
//{
//if (GetLastError() == ERROR_NO_TOKEN)
//{
//if (!ImpersonateSelf(SecurityImpersonation))
//return RTN_ERROR;
//
//if(!OpenThreadToken(GetCurrentThread(), TOKEN_ADJUST_PRIVILEGES | TOKEN_QUERY, FALSE, &hToken)){
//DisplayError("OpenThreadToken");
//return RTN_ERROR;
//}
//}
//else
//return RTN_ERROR;
//}
//
//// enable SeDebugPrivilege
//if(!SetPrivilege(hToken, SE_DEBUG_NAME, TRUE))
//{
//DisplayError("SetPrivilege");
//
//// close token handle
//CloseHandle(hToken);
//
//// indicate failure
//return RTN_ERROR;
//}
//
//// open the process
//if((hProcess = OpenProcess(
//PROCESS_ALL_ACCESS,
//FALSE,
//atoi(argv[1]) // PID from commandline
//)) == NULL)
//{
//DisplayError("OpenProcess");
//return RTN_ERROR;
//}
//
//// disable SeDebugPrivilege
//SetPrivilege(hToken, SE_DEBUG_NAME, FALSE);
//
//if(!TerminateProcess(hProcess, 0xffffffff))
//{
//DisplayError("TerminateProcess");
//dwRetVal=RTN_ERROR;
//}
//
//// close handles
//CloseHandle(hToken);
//CloseHandle(hProcess);
//
//return dwRetVal;
//}
//BOOL SetPrivilege(
//HANDLE hToken,          // token handle
//LPCTSTR Privilege,      // Privilege to enable/disable
//BOOL bEnablePrivilege   // TRUE to enable.  FALSE to disable
//)
//{
//TOKEN_PRIVILEGES tp;
//LUID luid;
//TOKEN_PRIVILEGES tpPrevious;
//DWORD cbPrevious=sizeof(TOKEN_PRIVILEGES);
//
//if(!LookupPrivilegeValue( NULL, Privilege, &luid )) return FALSE;
//
////
//// first pass.  get current privilege setting
////
//tp.PrivilegeCount           = 1;
//tp.Privileges[0].Luid       = luid;
//tp.Privileges[0].Attributes = 0;
//
//AdjustTokenPrivileges(
//hToken,
//FALSE,
//&tp,
//sizeof(TOKEN_PRIVILEGES),
//&tpPrevious,
//&cbPrevious
//);
//
//if (GetLastError() != ERROR_SUCCESS) return FALSE;
//
////
//// second pass.  set privilege based on previous setting
////
//tpPrevious.PrivilegeCount       = 1;
//tpPrevious.Privileges[0].Luid   = luid;
//
//if(bEnablePrivilege) {
//tpPrevious.Privileges[0].Attributes |= (SE_PRIVILEGE_ENABLED);
//}
//else {
//tpPrevious.Privileges[0].Attributes ^= (SE_PRIVILEGE_ENABLED &
//tpPrevious.Privileges[0].Attributes);
//}
//
//AdjustTokenPrivileges(
//hToken,
//FALSE,
//&tpPrevious,
//cbPrevious,
//NULL,
//NULL
//);
//
//if (GetLastError() != ERROR_SUCCESS) return FALSE;
//
//return TRUE;
//}
//BOOL SetPrivilege(
//HANDLE hToken,  // token handle
//LPCTSTR Privilege,  // Privilege to enable/disable
//BOOL bEnablePrivilege  // TRUE to enable. FALSE to disable
//)
//{
//TOKEN_PRIVILEGES tp = { 0 };
//// Initialize everything to zero
//LUID luid;
//DWORD cb=sizeof(TOKEN_PRIVILEGES);
//if(!LookupPrivilegeValue( NULL, Privilege, &luid ))
//return FALSE;
//tp.PrivilegeCount = 1;
//tp.Privileges[0].Luid = luid;
//if(bEnablePrivilege) {
//tp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED;
//} else {
//tp.Privileges[0].Attributes = 0;
//}
//AdjustTokenPrivileges( hToken, FALSE, &tp, cb, NULL, NULL );
//if (GetLastError() != ERROR_SUCCESS)
//return FALSE;
//
//return TRUE;
//}
//void DisplayError(
//LPTSTR szAPI    // pointer to failed API name
//)
//{
//LPTSTR MessageBuffer;
//DWORD dwBufferLength;
//
//fprintf(stderr,"%s() error!\n", szAPI);
//
//if(dwBufferLength=FormatMessage(
//FORMAT_MESSAGE_ALLOCATE_BUFFER |
//FORMAT_MESSAGE_FROM_SYSTEM,
//NULL,
//GetLastError(),
//GetSystemDefaultLangID(),
//(LPTSTR) &MessageBuffer,
//0,
//NULL
//))
//{
//DWORD dwBytesWritten;
//
////
//// Output message string on stderr
////
//WriteFile(
//GetStdHandle(STD_ERROR_HANDLE),
//MessageBuffer,
//dwBufferLength,
//&dwBytesWritten,
//NULL
//);
//
////
//// free the buffer allocated by the system
////
//LocalFree(MessageBuffer);
//}
//}
