package gowindows

const (
	ProcessBasicInformation = 0

	PROCESS_VM_READ = 0x0010
)

// 自适应版本，支持 32位读32位，64位读64位，但是不支持 32位读64位。
type PROCESS_BASIC_INFORMATION struct {
	Reserved1       uint
	PebBaseAddress  uint
	Reserved2       [2]uint
	UniqueProcessId uint
	Reserved3       uint
}

// 64位版本，支持 32位读64位，64位读64位
type PROCESS_BASIC_INFORMATION64 struct {
	Reserved1       uint64
	PebBaseAddress  uint64
	Reserved2       [2]uint64
	UniqueProcessId uint64
	Reserved3       uint64
}
