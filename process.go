package gowindows

// 标准 PEB 结构
// 32位程序@32位系统 或 64位程序@64位系统
// https://docs.microsoft.com/zh-cn/windows/desktop/api/winternl/ns-winternl-_peb
// http://terminus.rewolf.pl/terminus/structures/ntdll/_PEB_combined.html
//typedef struct _PEB {
//BYTE                          Reserved1[2];
//BYTE                          BeingDebugged;
//BYTE                          Reserved2[1];
//PVOID                         Reserved3[2];
//PPEB_LDR_DATA                 Ldr;
//PRTL_USER_PROCESS_PARAMETERS  ProcessParameters;
//PVOID                         Reserved4[3];
//PVOID                         AtlThunkSListPtr;
//PVOID                         Reserved5;
//ULONG                         Reserved6;
//PVOID                         Reserved7;
//ULONG                         Reserved8;
//ULONG                         AtlThunkSListPtr32;
//PVOID                         Reserved9[45];
//BYTE                          Reserved10[96];
//PPS_POST_PROCESS_INIT_ROUTINE PostProcessInitRoutine;
//BYTE                          Reserved11[128];
//PVOID                         Reserved12[1];
//ULONG                         SessionId;
//} PEB, *PPEB;
type PEB struct {
	Reserved1              [2]byte
	BeingDebugged          byte
	Reserved2              [1]byte
	Reserved3              [2]uint
	Ldr                    uint
	ProcessParameters      uint
	Reserved4              [3]uint
	AtlThunkSListPtr       uint
	Reserved5              uint
	Reserved6              uint32
	Reserved7              uint
	Reserved8              uint32
	AtlThunkSListPtr32     uint32
	Reserved9              [45]uint
	Reserved10             [96]byte
	PostProcessInitRoutine uint
	Reserved11             [128]byte
	Reserved12             [1]uint
	SessionId              uint32
}

// 64 位 PEB 结构
// 在32位程序下也能够和 64位PEB结构一致，解决了32位、64位对齐不一致的问题
// https://docs.microsoft.com/zh-cn/windows/desktop/api/winternl/ns-winternl-_peb
// http://terminus.rewolf.pl/terminus/structures/ntdll/_PEB_combined.html
//typedef struct _PEB {
//BYTE Reserved1[2];
//BYTE BeingDebugged;
//BYTE Reserved2[21];
//PPEB_LDR_DATA LoaderData;
//PRTL_USER_PROCESS_PARAMETERS ProcessParameters;
//BYTE Reserved3[520];
//PPS_POST_PROCESS_INIT_ROUTINE PostProcessInitRoutine;
//BYTE Reserved4[136];
//ULONG SessionId;
//} PEB;
type PEB64 struct {
	Reserved1              [2]byte
	BeingDebugged          byte
	Reserved2              [21]byte
	Ldr                    uint64
	ProcessParameters      uint64
	Reserved3              [520]byte
	PostProcessInitRoutine uint64
	Reserved4              [136]byte
	SessionId              uint32
	_                      uint32 // 保证 32 位下长度和64位一致，64位最后会有个 uint32 的空白用于对齐
}

// https://docs.microsoft.com/zh-cn/windows/desktop/api/winternl/ns-winternl-_rtl_user_process_parameters
//typedef struct _RTL_USER_PROCESS_PARAMETERS {
//BYTE           Reserved1[16];
//PVOID          Reserved2[10];
//UNICODE_STRING ImagePathName;
//UNICODE_STRING CommandLine;
//} RTL_USER_PROCESS_PARAMETERS, *PRTL_USER_PROCESS_PARAMETERS;
type RTL_USER_PROCESS_PARAMETERS struct {
	Reserved1     [16]byte
	Reserved2     [10]uint
	ImagePathName UNICODE_STRING
	CommandLine   UNICODE_STRING
}
type RTL_USER_PROCESS_PARAMETERS64 struct {
	Reserved1     [16]byte
	Reserved2     [10]uint64
	ImagePathName UNICODE_STRING64
	CommandLine   UNICODE_STRING64
}

// https://docs.microsoft.com/en-us/windows/desktop/winprog/windows-data-types
//typedef struct _UNICODE_STRING {
//USHORT  Length;
//USHORT  MaximumLength;
//PWSTR  Buffer;
//} UNICODE_STRING;
type UNICODE_STRING struct {
	Length        uint16 // buffer的字节长度，不包括终止符“NULL”
	MaximumLength uint16

	// 	指向以null结尾的16位Unicode字符串的指针。有关更多信息，请参阅[字体使用的字符集]（https://msdn.microsoft.com/library/windows/desktop/dd183415）。
	Buffer uint
}
type UNICODE_STRING64 struct {
	Length        uint16 // buffer的字节长度，不包括终止符“NULL”
	MaximumLength uint16
	_             [4]byte // 32位下强制64位对齐

	// 	指向以null结尾的16位Unicode字符串的指针。有关更多信息，请参阅[字体使用的字符集]（https://msdn.microsoft.com/library/windows/desktop/dd183415）。
	Buffer uint64
}
