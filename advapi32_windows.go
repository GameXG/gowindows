package gowindows

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")

	procLookupPrivilegeValueW                                = modadvapi32.NewProc("LookupPrivilegeValueW")
	procAdjustTokenPrivileges                                = modadvapi32.NewProc("AdjustTokenPrivileges")
	procConvertStringSecurityDescriptorToSecurityDescriptorW = modadvapi32.NewProc("ConvertStringSecurityDescriptorToSecurityDescriptorW")
	getSecurityDescriptorSacl                                = modadvapi32.NewProc("GetSecurityDescriptorSacl")
	setSecurityInfo                                          = modadvapi32.NewProc("SetSecurityInfo")
	setNamedSecurityInfo                                     = modadvapi32.NewProc("SetNamedSecurityInfoW")
)

func adjustTokenPrivileges(token windows.Token, disableAllPrivileges bool, newstate *TOKEN_PRIVILEGES, buflen uint32, prevstate *TOKEN_PRIVILEGES, returnlen *uint32) (ret uint32, err error) {
	var _p0 uint32
	if disableAllPrivileges {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(), 6, uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(newstate)), uintptr(buflen), uintptr(unsafe.Pointer(prevstate)), uintptr(unsafe.Pointer(returnlen)))
	ret = uint32(r0)
	if true {
		if e1 != 0 {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// AdjustTokenPrivileges功能允许或禁止特定的特权访问令牌。启用或禁用访问令牌中的权限需要TOKEN_ADJUST_PRIVILEGES访问权限。
// https://msdn.microsoft.com/en-us/library/windows/desktop/aa375202(v=vs.85).aspx
func AdjustTokenPrivileges(token windows.Token, disableAllPrivileges bool, newstate *TOKEN_PRIVILEGES, buflen uint32, prevstate *TOKEN_PRIVILEGES, returnlen *uint32) error {
	ret, err := adjustTokenPrivileges(token, disableAllPrivileges, newstate, buflen, prevstate, returnlen)
	if ret == 0 {
		// AdjustTokenPrivileges call failed
		return err
	}
	// AdjustTokenPrivileges call succeeded
	if err == syscall.EINVAL {
		// GetLastError returned ERROR_SUCCESS
		return nil
	}
	return err
}

// LookupPrivilegeValue函数检索 本地唯一性标识符（LUID）一个指定系统上用于局部地表示指定的权限名称。
// https://docs.microsoft.com/en-us/windows/desktop/api/winbase/nf-winbase-lookupprivilegevaluea
func LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) (err error) {
	r1, _, e1 := syscall.Syscall(procLookupPrivilegeValueW.Addr(), 3, uintptr(unsafe.Pointer(systemname)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(luid)))
	if r1 == 0 {
		if e1 != 0 {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

//
// https://docs.microsoft.com/en-us/windows/desktop/api/sddl/nf-sddl-convertstringsecuritydescriptortosecuritydescriptorw
//BOOL
//WINAPI
//ConvertStringSecurityDescriptorToSecurityDescriptorW(
//    _In_ LPCWSTR StringSecurityDescriptor,
//    _In_ DWORD StringSDRevision,
//    _Outptr_ PSECURITY_DESCRIPTOR * SecurityDescriptor,
//    _Out_opt_ PULONG SecurityDescriptorSize
//    );
// 注意：需要自己调用 LocalFree 释放 securityDescriptor 占用空间。
func ConvertStringSecurityDescriptorToSecurityDescriptor(stringSecurityDescriptor string, stringSDRevision SddlRevision,
	securityDescriptor *SecurityDescriptor, securityDescriptorSize *ULong) error {
	stringSecurityDescriptorUtf16, err := windows.UTF16PtrFromString(stringSecurityDescriptor)
	if err != nil {
		return err
	}

	r1, _, e1 := procConvertStringSecurityDescriptorToSecurityDescriptorW.Call(
		uintptr(unsafe.Pointer(stringSecurityDescriptorUtf16)), uintptr(stringSDRevision),
		uintptr(unsafe.Pointer(securityDescriptor)), uintptr(unsafe.Pointer(securityDescriptorSize)))

	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}

	return nil
}

// https://docs.microsoft.com/en-us/windows/desktop/api/securitybaseapi/nf-securitybaseapi-getsecuritydescriptorsacl
//BOOL GetSecurityDescriptorSacl(
//  PSECURITY_DESCRIPTOR pSecurityDescriptor,
//  LPBOOL               lpbSaclPresent,
//  PACL                 *pSacl,
//  LPBOOL               lpbSaclDefaulted
//);
func GetSecurityDescriptorSacl(securityDescriptor SecurityDescriptor, lpbSaclPresent *Bool, sacl **ACL, saclDefaulted *Bool) error {
	r1, _, e1 := getSecurityDescriptorSacl.Call(uintptr(unsafe.Pointer(securityDescriptor)), uintptr(unsafe.Pointer(lpbSaclPresent)), uintptr(unsafe.Pointer(sacl)), uintptr(unsafe.Pointer(saclDefaulted)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

/*
// https://docs.microsoft.com/en-us/windows/desktop/api/aclapi/nf-aclapi-setsecurityinfo
DWORD
WINAPI
SetSecurityInfo(
    _In_     HANDLE                handle,
    _In_     SE_OBJECT_TYPE        ObjectType,
    _In_     SECURITY_INFORMATION  SecurityInfo,
    _In_opt_ PSID                  psidOwner,
    _In_opt_ PSID                  psidGroup,
    _In_opt_ PACL                  pDacl,
    _In_opt_ PACL                  pSacl
    );
*/
func SetSecurityInfo(handle Handle, objectType SeObjectType, securityInfo SecurityInformation,
	psidOwner, psidGroup PSId, pDacl, pSacl *ACL) error {
	r1, _, e1 := setSecurityInfo.Call(uintptr(handle), uintptr(objectType), uintptr(securityInfo),
		uintptr(unsafe.Pointer(psidOwner)), uintptr(unsafe.Pointer(psidGroup)), uintptr(unsafe.Pointer(pDacl)),
		uintptr(unsafe.Pointer(pSacl)))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/desktop/api/aclapi/nf-aclapi-setnamedsecurityinfow
//DWORD SetNamedSecurityInfoW(
//  LPWSTR               pObjectName,
//  SE_OBJECT_TYPE       ObjectType,
//  SECURITY_INFORMATION SecurityInfo,
//  PSID                 psidOwner,
//  PSID                 psidGroup,
//  PACL                 pDacl,
//  PACL                 pSacl
//);
func SetNamedSecurityInfoW(objectName string, objectType SeObjectType, securityInfo SecurityInformation,
	psidOwner, psidGroup PSId, pDacl, pSacl *ACL) error {

	_objectName, err := windows.UTF16PtrFromString(objectName)
	if err != nil {
		return err
	}

	r1, _, e1 := setNamedSecurityInfo.Call(uintptr(unsafe.Pointer(_objectName)), uintptr(objectType), uintptr(securityInfo),
		uintptr(unsafe.Pointer(psidOwner)), uintptr(unsafe.Pointer(psidGroup)), uintptr(unsafe.Pointer(pDacl)),
		uintptr(unsafe.Pointer(pSacl)))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

// 降低内核对象的安全级别
func SetObjectToLowIntegrity(h Handle) error {
	return SetObjectIntegrity(h, LOW_INTEGRITY_SDDL_SACL_W)
}
func SetObjectIntegrity(h Handle, stringSecurityDescriptor string) error {

	var securityDescriptor SecurityDescriptor
	var securityDescriptorSize ULong

	err := ConvertStringSecurityDescriptorToSecurityDescriptor(stringSecurityDescriptor, SDDL_REVISION_1,
		&securityDescriptor, &securityDescriptorSize)
	if err != nil {
		return fmt.Errorf("ConvertStringSecurityDescriptorToSecurityDescriptor(), %v", err)
	}
	defer func() {
		LocalFree(windows.Pointer(securityDescriptor))
	}()

	var fSaclPresent Bool = 0
	var pSacl *ACL
	var fSaclDefaulted Bool = 0

	err = GetSecurityDescriptorSacl(securityDescriptor, &fSaclPresent, &pSacl, &fSaclDefaulted)
	if err != nil {
		return fmt.Errorf("GetSecurityDescriptorSacl(), %v", err)
	}

	err = SetSecurityInfo(h, SE_KERNEL_OBJECT, LABEL_SECURITY_INFORMATION,
		nil, nil, nil, pSacl)
	if err != nil {
		return fmt.Errorf("SetSecurityInfo(), %v", err)
	}

	return nil
}

// 降低内核对象的安全级别
// 不知道什么原因，对 mmap 无效，mmap直接在创建时指定可行。
func SetObjectToLowIntegrityWithName(objectName string) error {
	return SetObjectIntegrityWithName(objectName, LOW_INTEGRITY_SDDL_SACL_W)
}
func SetObjectIntegrityWithName(objectName string, stringSecurityDescriptor string) error {

	var securityDescriptor SecurityDescriptor
	var securityDescriptorSize ULong

	err := ConvertStringSecurityDescriptorToSecurityDescriptor(stringSecurityDescriptor, SDDL_REVISION_1,
		&securityDescriptor, &securityDescriptorSize)
	if err != nil {
		return fmt.Errorf("ConvertStringSecurityDescriptorToSecurityDescriptor(), %v", err)
	}
	defer func() {
		LocalFree(windows.Pointer(securityDescriptor))
	}()

	var fSaclPresent Bool = 0
	var pSacl *ACL
	var fSaclDefaulted Bool = 0

	err = GetSecurityDescriptorSacl(securityDescriptor, &fSaclPresent, &pSacl, &fSaclDefaulted)
	if err != nil {
		return fmt.Errorf("GetSecurityDescriptorSacl(), %v", err)
	}

	err = SetNamedSecurityInfoW(objectName, SE_KERNEL_OBJECT, LABEL_SECURITY_INFORMATION,
		nil, nil, nil, pSacl)
	if err != nil {
		return fmt.Errorf("SetSecurityInfo(), %v", err)
	}

	return nil
}
