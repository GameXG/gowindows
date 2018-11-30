package gowindows

import (
	"runtime"
	"testing"

	"golang.org/x/sys/windows"
)

func TestConvertStringSecurityDescriptorToSecurityDescriptor(t *testing.T) {
	var securityDescriptor SecurityDescriptor
	var securityDescriptorSize ULong

	err := ConvertStringSecurityDescriptorToSecurityDescriptor(LOW_INTEGRITY_SDDL_SACL_W, SDDL_REVISION_1, &securityDescriptor, &securityDescriptorSize)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := LocalFree(windows.Pointer(securityDescriptor))
		if err != nil {
			t.Fatal(err)
		}
	}()

	if securityDescriptor == nil {
		t.Error("securityDescriptor==nil")
	}

	if securityDescriptorSize == 0 {
		t.Error("securityDescriptorSize == 0")
	}
}

func TestSetObjectToLowIntegrity(t *testing.T) {
	err := SetSelfProcessPrivilege(SE_CREATE_GLOBAL_NAME, true)
	if err != nil {
		t.Fatal(err)
	}

	const MMAP_NAME = "Global\\757aykjhgsdgdhrh"
	m, err := CreateMmap(MMAP_NAME, 1024, true)
	if err != nil {
		t.Fatal(err)
	}

	err = SetObjectToLowIntegrity(m.GetHandle())
	if err != nil {
		t.Fatal(err)
	}
}
func TestSetObjectToLowIntegrityWithName(t *testing.T) {
	err := SetSelfProcessPrivilege(SE_CREATE_GLOBAL_NAME, true)
	if err != nil {
		t.Fatal(err)
	}

	const MMAP_NAME = "Global\\e4h7igfdhh"
	m, err := CreateMmap(MMAP_NAME, 1024, true)
	if err != nil {
		t.Fatal(err)
	}

	err = SetObjectToLowIntegrityWithName(MMAP_NAME)
	if err != nil {
		t.Fatal(err)
	}

	runtime.KeepAlive(m)
}

func TestSE_KERNEL_OBJECT(t *testing.T) {
	if SE_KERNEL_OBJECT != 6 {
		t.Fatal("!=")
	}
}
