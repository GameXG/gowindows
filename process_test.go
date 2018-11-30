package gowindows

import (
	"testing"
	"unsafe"
)

// 测试PEB结构是否正确
func TestPEB(t *testing.T) {
	p := PEB{}

	processParameters := unsafe.Offsetof(p.ProcessParameters)
	atlThunkSListPtr := unsafe.Offsetof(p.AtlThunkSListPtr)
	atlThunkSListPtr32 := unsafe.Offsetof(p.AtlThunkSListPtr32)
	postProcessInitRoutine := unsafe.Offsetof(p.PostProcessInitRoutine)
	sessionId := unsafe.Offsetof(p.SessionId)

	switch ptrSize {
	case 4:
		if processParameters != 0x10 {
			t.Errorf("%X!=0x10", processParameters)
		}

		if atlThunkSListPtr != 0x20 {
			t.Errorf("%X!=0x20", atlThunkSListPtr)
		}

		if atlThunkSListPtr32 != 0x34 {
			t.Errorf("%X!=0x34", atlThunkSListPtr32)
		}

		if postProcessInitRoutine != 0x14c {
			t.Errorf("%X!=0x1D4", postProcessInitRoutine)
		}

		if sessionId != 0x1D4 {
			t.Errorf("%X!=0x1D4", sessionId)
		}

		s := unsafe.Sizeof(p)
		if s != 0x1d8 {
			t.Errorf("%X!=0x1D8", s)
		}

	case 8:
		if processParameters != 0x20 {
			t.Errorf("%X!=0x20", processParameters)
		}

		if atlThunkSListPtr != 0x40 {
			t.Errorf("%X!=0x40", atlThunkSListPtr)
		}

		if atlThunkSListPtr32 != 0x64 {
			t.Errorf("%X!=0x64", atlThunkSListPtr32)
		}

		if postProcessInitRoutine != 0x230 {
			t.Errorf("%X!=0x230", postProcessInitRoutine)
		}

		if sessionId != 0x2c0 {
			t.Errorf("%X!=0x2c0", sessionId)
		}

		// 最后一个字段是32位，当64位环境下会在其后部空出32位字节的空白，用于下一结构的对齐。
		// 尺寸将是 0x2c8
		s := unsafe.Sizeof(p)
		if s != 0x2c8 {
			t.Errorf("%X!=0x2c8", s)
		}
	}
}

func TestPEB64(t *testing.T) {
	p := PEB64{}
	beingDebugged := unsafe.Offsetof(p.BeingDebugged)
	ldr := unsafe.Offsetof(p.Ldr)
	processParameters := unsafe.Offsetof(p.ProcessParameters)
	postProcessInitRoutine := unsafe.Offsetof(p.PostProcessInitRoutine)
	sessionId := unsafe.Offsetof(p.SessionId)

	if beingDebugged != 0x2 {
		t.Errorf("%X!=0x2", beingDebugged)
	}
	if ldr != 0x18 {
		t.Errorf("%X!=0x18", ldr)
	}
	if processParameters != 0x20 {
		t.Errorf("%X!=0x20", processParameters)
	}
	if postProcessInitRoutine != 0x230 {
		t.Errorf("%X!=0x230", postProcessInitRoutine)
	}

	if sessionId != 0x2c0 {
		t.Errorf("%X!=0x2c0", sessionId)
	}

	s := unsafe.Sizeof(p)
	if s != 0x2c8 {
		t.Errorf("%X!=0x2c8", s)
	}
}

func TestRTL_USER_PROCESS_PARAMETERS(t *testing.T) {
	p := RTL_USER_PROCESS_PARAMETERS{}

	imagePathName := unsafe.Offsetof(p.ImagePathName)
	commandLine := unsafe.Offsetof(p.CommandLine)

	switch ptrSize {
	case 4:
		if imagePathName != 0x38 {
			t.Errorf("%X!=0x38", imagePathName)
		}

		if commandLine != 0x40 {
			t.Errorf("%X!=0x40", commandLine)
		}

		s := unsafe.Sizeof(p)
		if s != 0x48 {
			t.Errorf("%X!=0x48", s)
		}

	case 8:
		if imagePathName != 0x60 {
			t.Errorf("%X!=0x60", imagePathName)
		}

		if commandLine != 0x70 {
			t.Errorf("%X!=0x70", commandLine)
		}

		// 最后一个字段是32位，当64位环境下会在其后部空出32位字节的空白，用于下一结构的对齐。
		// 尺寸将是 0x2c8
		s := unsafe.Sizeof(p)
		if s != 0x80 {
			t.Errorf("%X!=0x80", s)
		}
	}
}

func TestRTL_USER_PROCESS_PARAMETERS64(t *testing.T) {
	p := RTL_USER_PROCESS_PARAMETERS64{}

	imagePathName := unsafe.Offsetof(p.ImagePathName)
	commandLine := unsafe.Offsetof(p.CommandLine)

	if imagePathName != 0x60 {
		t.Errorf("%X!=0x60", imagePathName)
	}
	if commandLine != 0x70 {
		t.Errorf("%X!=0x70", commandLine)
	}

	s := unsafe.Sizeof(p)
	if s != 0x80 {
		t.Errorf("%X!=0x80", s)
	}
}

func TestUNICODE_STRING(t *testing.T) {
	p := UNICODE_STRING{}

	length := unsafe.Offsetof(p.Length)
	maximumLength := unsafe.Offsetof(p.MaximumLength)
	buffer := unsafe.Offsetof(p.Buffer)

	switch ptrSize {
	case 4:
		if length != 0x0 {
			t.Errorf("%X!=0x0", length)
		}

		if maximumLength != 0x2 {
			t.Errorf("%X!=0x2", maximumLength)
		}

		if buffer != 0x4 {
			t.Errorf("%X!=0x4", buffer)
		}

		s := unsafe.Sizeof(p)
		if s != 0x8 {
			t.Errorf("%X!=0x8", s)
		}

	case 8:
		if length != 0x0 {
			t.Errorf("%X!=0x0", length)
		}

		if maximumLength != 0x2 {
			t.Errorf("%X!=0x2", maximumLength)
		}

		if buffer != 0x8 {
			t.Errorf("%X!=0x8", buffer)
		}

		// 最后一个字段是32位，当64位环境下会在其后部空出32位字节的空白，用于下一结构的对齐。
		// 尺寸将是 0x2c8
		s := unsafe.Sizeof(p)
		if s != 0x10 {
			t.Errorf("%X!=0x10", s)
		}
	}
}

func TestUNICODE_STRING64(t *testing.T) {
	p := UNICODE_STRING64{}

	Length := unsafe.Offsetof(p.Length)
	MaximumLength := unsafe.Offsetof(p.MaximumLength)
	Buffer := unsafe.Offsetof(p.Buffer)

	if Length != 0 {
		t.Errorf("%X!=0x60", Length)
	}
	if MaximumLength != 0x2 {
		t.Errorf("%X!=0x2", MaximumLength)
	}
	if Buffer != 0x8 {
		t.Errorf("%X!=0x8", Buffer)
	}

	s := unsafe.Sizeof(p)
	if s != 0x10 {
		t.Errorf("%X!=0x10", s)
	}
}
