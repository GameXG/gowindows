package gowindows

import (
	"testing"

	"unsafe"

	"os"

	"golang.org/x/sys/windows"
)

func TestFwpmEngineOpen0AndClose(t *testing.T) {
	var engineHandle Handle
	session := FwpmSession0{
		Flags: FWPM_SESSION_FLAG_DYNAMIC,
	}

	err := FwpmEngineOpen0("", RPC_C_AUTHN_WINNT, nil, &session, &engineHandle)
	if err != nil {
		t.Fatal(err)
	}

	err = FwpmEngineClose0(engineHandle)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFwpmGetAppIdFromFileName0AndFree(t *testing.T) {
	var appId *FwpByteBlob
	err := FwpmGetAppIdFromFileName0(os.Args[0], &appId)
	if err != nil {
		t.Fatal(err)
	}

	err = FwpmFreeMemory0((*windows.Pointer)(unsafe.Pointer(&appId)))
	if err != nil {
		t.Fatal(err)
	}
}

func TestFwpmGetAppIdFromFileName0AndFree2(t *testing.T) {
	var appId *FwpByteBlob
	err := FwpmGetAppIdFromFileName0("C:\\aaaaaaa.exe", &appId)
	if err == nil {
		t.Fatal("==nil")
	}
}
