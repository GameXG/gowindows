package gowindows

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	fwpuclnt                  = windows.NewLazyDLL("Fwpuclnt.dll")
	fwpmEngineOpen0           = fwpuclnt.NewProc("FwpmEngineOpen0")
	fwpmEngineClose0          = fwpuclnt.NewProc("FwpmEngineClose0")
	fwpmSubLayerAdd0          = fwpuclnt.NewProc("FwpmSubLayerAdd0")
	fwpmSubLayerDeleteByKey0  = fwpuclnt.NewProc("FwpmSubLayerDeleteByKey0")
	fwpmGetAppIdFromFileName0 = fwpuclnt.NewProc("FwpmGetAppIdFromFileName0")
	fwpmFreeMemory0           = fwpuclnt.NewProc("FwpmFreeMemory0")
	fwpmFilterAdd0            = fwpuclnt.NewProc("FwpmFilterAdd0")
	fwpmFilterDeleteById0     = fwpuclnt.NewProc("FwpmFilterDeleteById0")
)

// FwpmEngineOpen0
// The FwpmEngineOpen0 function opens a session to the filter engine.
// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-fwpmengineopen0
//DWORD FwpmEngineOpen0(
//const wchar_t             *serverName,
//UINT32                    authnService,
//SEC_WINNT_AUTH_IDENTITY_W *authIdentity,
//const FWPM_SESSION0       *session,
//HANDLE                    *engineHandle
//);
// 	Windows Vista [desktop apps only] Windows Server 2008 [desktop apps only]
func FwpmEngineOpen0(serverName string, authnService RpcCAuthnType, authIdentity unsafe.Pointer, session *FwpmSession0, engineHandle *Handle) error {
	if err := fwpmEngineOpen0.Find(); err != nil {
		return err
	}

	var _serverName *uint16
	if len(serverName) > 0 {
		var err error
		_serverName, err = windows.UTF16PtrFromString(serverName)
		if err != nil {
			return err
		}
	}

	r1, _, e1 := fwpmEngineOpen0.Call(uintptr(unsafe.Pointer(_serverName)), uintptr(authnService), uintptr(authIdentity), uintptr(unsafe.Pointer(session)), uintptr(unsafe.Pointer(engineHandle)))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return newFwpmError(_HRESULT_TYPEDEF_(r1))
		}
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/content/fwpmk/nf-fwpmk-fwpmengineclose0
//NTSTATUS FwpmEngineClose0(
//HANDLE engineHandle
//);
// 	Windows Vista [desktop apps only] Windows Server 2008 [desktop apps only]
func FwpmEngineClose0(engineHandle Handle) error {
	if err := fwpmEngineClose0.Find(); err != nil {
		return err
	}

	r1, _, e1 := fwpmEngineClose0.Call(uintptr(engineHandle))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return newFwpmError(_HRESULT_TYPEDEF_(r1))
		}
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-FwpmSubLayerAdd0
//DWORD
//WINAPI
//FwpmSubLayerAdd0(
//_In_ HANDLE engineHandle,
//_In_ const FWPM_SUBLAYER0* subLayer,
//_In_opt_ PSECURITY_DESCRIPTOR sd
//);
// Windows Vista [desktop apps only] Windows Server 2008 [desktop apps only]
func FwpmSubLayerAdd0(engineHandle Handle, subLayer *FwpmSublayer0, sd PSecurityDescriptor) error {
	if err := fwpmSubLayerAdd0.Find(); err != nil {
		return err
	}

	r1, _, e1 := fwpmSubLayerAdd0.Call(uintptr(engineHandle), uintptr(unsafe.Pointer(subLayer)), uintptr(sd))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return newFwpmError(_HRESULT_TYPEDEF_(r1))
		}
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-fwpmgetappidfromfilename0
//DWORD FwpmSubLayerDeleteByKey0(
////HANDLE     engineHandle,
////const GUID *key
////);
// Windows Vista [desktop apps only] Windows Server 2008 [desktop apps only]
func FwpmSubLayerDeleteByKey0(engineHandle Handle, key *GUID) error {
	if err := fwpmSubLayerDeleteByKey0.Find(); err != nil {
		return err
	}

	r1, _, e1 := fwpmSubLayerDeleteByKey0.Call(uintptr(engineHandle), uintptr(unsafe.Pointer(key)))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return newFwpmError(_HRESULT_TYPEDEF_(r1))
		}
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-fwpmgetappidfromfilename0
// 注意：使用完毕后需要使用 FwpmFreeMemory0 释放 appId 。
//DWORD FwpmGetAppIdFromFileName0(
//PCWSTR        fileName,
//FWP_BYTE_BLOB **appId
//);
// Windows Vista [desktop apps only] Windows Server 2008 [desktop apps only]
func FwpmGetAppIdFromFileName0(fileName string, appId **FwpByteBlob) error {
	if err := fwpmGetAppIdFromFileName0.Find(); err != nil {
		return err
	}

	_fileName, err := windows.UTF16PtrFromString(fileName)
	if err != nil {
		return err
	}

	r1, _, e1 := fwpmGetAppIdFromFileName0.Call(uintptr(unsafe.Pointer(_fileName)), uintptr(unsafe.Pointer(appId)))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return newFwpmError(_HRESULT_TYPEDEF_(r1))
		}
	}

	return nil
}

// https://docs.microsoft.com/zh-cn/windows/desktop/api/fwpmu/nf-fwpmu-fwpmfreememory0
//void FwpmFreeMemory0(
//void **p
//);
func FwpmFreeMemory0(p *windows.Pointer) error {
	if err := fwpmFreeMemory0.Find(); err != nil {
		return err
	}

	_, _, err := fwpmFreeMemory0.Call(uintptr(unsafe.Pointer(p)))
	if err == ERROR_SUCCESS {
		return nil
	}
	return err
}

//DWORD FwpmFilterAdd0(
//HANDLE               engineHandle,
//const FWPM_FILTER0   *filter,
//PSECURITY_DESCRIPTOR sd,
//UINT64               *id
//);
func FwpmFilterAdd0(engineHandle Handle, filter *FwpmFilter0, sd PSecurityDescriptor, id *FilterId) error {
	if err := fwpmFilterAdd0.Find(); err != nil {
		return err
	}

	r1, _, e1 := fwpmFilterAdd0.Call(uintptr(engineHandle), uintptr(unsafe.Pointer(filter)), uintptr(sd), uintptr(unsafe.Pointer(id)))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return newFwpmError(_HRESULT_TYPEDEF_(r1))
		}
	}

	return nil
}

//DWORD
//WINAPI
//FwpmFilterDeleteById0(
//_In_ HANDLE engineHandle,
//_In_ UINT64 id
//);
func FwpmFilterDeleteById0(engineHandle Handle, id FilterId) error {
	if err := fwpmFilterDeleteById0.Find(); err != nil {
		return err
	}

	var r1 uintptr
	var e1 error

	if ptrSize == 8 {
		r1, _, e1 = fwpmFilterDeleteById0.Call(uintptr(engineHandle), uintptr(id))
	} else {
		r1, _, e1 = fwpmFilterDeleteById0.Call(uintptr(engineHandle), uintptr(id), uintptr(id>>32))
	}
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return newFwpmError(_HRESULT_TYPEDEF_(r1))
		}
	}

	return nil
}
