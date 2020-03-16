package gowindows

import (
	"unsafe"

	"fmt"
)

type FwpmSessionType uint32

const (

	// 设置此标志后，会话结束时会自动删除会话期间添加的所有对象。
	FWPM_SESSION_FLAG_DYNAMIC FwpmSessionType = 0x00000001

	FWPM_SESSION_FLAG_RESERVED FwpmSessionType = 0x10000000
)

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmtypes/ns-fwpmtypes-fwpm_sublayer0_
type FwpmSublayerFlag uint32

const (
	// 导致子层持久化，在BFE停止/启动时存活。
	FWPM_SUBLAYER_FLAG_PERSISTENT FwpmSublayerFlag = (0x00000001)
)

type FwpDataType int

const (
	FWP_EMPTY                         FwpDataType = 0
	FWP_UINT8                         FwpDataType = (FWP_EMPTY + 1)
	FWP_UINT16                        FwpDataType = (FWP_UINT8 + 1)
	FWP_UINT32                        FwpDataType = (FWP_UINT16 + 1)
	FWP_UINT64                        FwpDataType = (FWP_UINT32 + 1)
	FWP_INT8                          FwpDataType = (FWP_UINT64 + 1)
	FWP_INT16                         FwpDataType = (FWP_INT8 + 1)
	FWP_INT32                         FwpDataType = (FWP_INT16 + 1)
	FWP_INT64                         FwpDataType = (FWP_INT32 + 1)
	FWP_FLOAT                         FwpDataType = (FWP_INT64 + 1)
	FWP_DOUBLE                        FwpDataType = (FWP_FLOAT + 1)
	FWP_BYTE_ARRAY16_TYPE             FwpDataType = (FWP_DOUBLE + 1)
	FWP_BYTE_BLOB_TYPE                FwpDataType = (FWP_BYTE_ARRAY16_TYPE + 1)
	FWP_SID                           FwpDataType = (FWP_BYTE_BLOB_TYPE + 1)
	FWP_SECURITY_DESCRIPTOR_TYPE      FwpDataType = (FWP_SID + 1)
	FWP_TOKEN_INFORMATION_TYPE        FwpDataType = (FWP_SECURITY_DESCRIPTOR_TYPE + 1)
	FWP_TOKEN_ACCESS_INFORMATION_TYPE FwpDataType = (FWP_TOKEN_INFORMATION_TYPE + 1)
	FWP_UNICODE_STRING_TYPE           FwpDataType = (FWP_TOKEN_ACCESS_INFORMATION_TYPE + 1)
	FWP_BYTE_ARRAY6_TYPE              FwpDataType = (FWP_UNICODE_STRING_TYPE + 1)
	FWP_SINGLE_DATA_TYPE_MAX          FwpDataType = 0xff
	FWP_V4_ADDR_MASK                  FwpDataType = (FWP_SINGLE_DATA_TYPE_MAX + 1)
	FWP_V6_ADDR_MASK                  FwpDataType = (FWP_V4_ADDR_MASK + 1)
	FWP_RANGE_TYPE                    FwpDataType = (FWP_V6_ADDR_MASK + 1)
	FWP_DATA_TYPE_MAX                 FwpDataType = (FWP_RANGE_TYPE + 1)
)

type FwpMatchType int

const (
	FWP_MATCH_EQUAL                  FwpMatchType = 0
	FWP_MATCH_GREATER                FwpMatchType = (FWP_MATCH_EQUAL + 1)
	FWP_MATCH_LESS                   FwpMatchType = (FWP_MATCH_GREATER + 1)
	FWP_MATCH_GREATER_OR_EQUAL       FwpMatchType = (FWP_MATCH_LESS + 1)
	FWP_MATCH_LESS_OR_EQUAL          FwpMatchType = (FWP_MATCH_GREATER_OR_EQUAL + 1)
	FWP_MATCH_RANGE                  FwpMatchType = (FWP_MATCH_LESS_OR_EQUAL + 1)
	FWP_MATCH_FLAGS_ALL_SET          FwpMatchType = (FWP_MATCH_RANGE + 1)
	FWP_MATCH_FLAGS_ANY_SET          FwpMatchType = (FWP_MATCH_FLAGS_ALL_SET + 1)
	FWP_MATCH_FLAGS_NONE_SET         FwpMatchType = (FWP_MATCH_FLAGS_ANY_SET + 1)
	FWP_MATCH_EQUAL_CASE_INSENSITIVE FwpMatchType = (FWP_MATCH_FLAGS_NONE_SET + 1)
	FWP_MATCH_NOT_EQUAL              FwpMatchType = (FWP_MATCH_EQUAL_CASE_INSENSITIVE + 1)
	FWP_MATCH_TYPE_MAX               FwpMatchType = (FWP_MATCH_NOT_EQUAL + 1)
)

type FwpActionType uint32

const (
	FWP_ACTION_FLAG_TERMINATING     = (0x00001000)
	FWP_ACTION_FLAG_NON_TERMINATING = (0x00002000)
	FWP_ACTION_FLAG_CALLOUT         = (0x00004000)

	FWP_ACTION_BLOCK               = (0x00000001 | FWP_ACTION_FLAG_TERMINATING)
	FWP_ACTION_PERMIT              = (0x00000002 | FWP_ACTION_FLAG_TERMINATING) //允许
	FWP_ACTION_CALLOUT_TERMINATING = (0x00000003 | FWP_ACTION_FLAG_CALLOUT | FWP_ACTION_FLAG_TERMINATING)
	FWP_ACTION_CALLOUT_INSPECTION  = (0x00000004 | FWP_ACTION_FLAG_CALLOUT | FWP_ACTION_FLAG_NON_TERMINATING)
	FWP_ACTION_CALLOUT_UNKNOWN     = (0x00000005 | FWP_ACTION_FLAG_CALLOUT)
	FWP_ACTION_CONTINUE            = (0x00000006 | FWP_ACTION_FLAG_NON_TERMINATING) //继续下一个过滤器
	FWP_ACTION_NONE                = (0x00000007)
	FWP_ACTION_NONE_NO_MATCH       = (0x00000008)
)

const (
	FWPM_FILTER_FLAG_NONE                 = (0x00000000)
	FWPM_FILTER_FLAG_PERSISTENT           = (0x00000001)
	FWPM_FILTER_FLAG_BOOTTIME             = (0x00000002)
	FWPM_FILTER_FLAG_HAS_PROVIDER_CONTEXT = (0x00000004)
	// 高优先
	// https://docs.microsoft.com/zh-cn/windows/win32/fwp/filter-arbitration
	FWPM_FILTER_FLAG_CLEAR_ACTION_RIGHT                  = (0x00000008)
	FWPM_FILTER_FLAG_PERMIT_IF_CALLOUT_UNREGISTERED      = (0x00000010)
	FWPM_FILTER_FLAG_DISABLED                            = (0x00000020)
	FWPM_FILTER_FLAG_INDEXED                             = (0x00000040)
	FWPM_FILTER_FLAG_HAS_SECURITY_REALM_PROVIDER_CONTEXT = (0x00000080)
	FWPM_FILTER_FLAG_SYSTEMOS_ONLY                       = (0x00000100)
	FWPM_FILTER_FLAG_GAMEOS_ONLY                         = (0x00000200)
	FWPM_FILTER_FLAG_SILENT_MODE                         = (0x00000400)
	FWPM_FILTER_FLAG_IPSEC_NO_ACQUIRE_INITIATE           = (0x00000800)
)

//
//typedef enum {
//#if(_WIN32_WINNT >= 0x0501)
//    IPPROTO_HOPOPTS       = 0,  // IPv6 Hop-by-Hop options
//#endif//(_WIN32_WINNT >= 0x0501)
//    IPPROTO_ICMP          = 1,
//    IPPROTO_IGMP          = 2,
//    IPPROTO_GGP           = 3,
//#if(_WIN32_WINNT >= 0x0501)
//    IPPROTO_IPV4          = 4,
//#endif//(_WIN32_WINNT >= 0x0501)
//#if(_WIN32_WINNT >= 0x0600)
//    IPPROTO_ST            = 5,
//#endif//(_WIN32_WINNT >= 0x0600)
//    IPPROTO_TCP           = 6,
//#if(_WIN32_WINNT >= 0x0600)
//    IPPROTO_CBT           = 7,
//    IPPROTO_EGP           = 8,
//    IPPROTO_IGP           = 9,
//#endif//(_WIN32_WINNT >= 0x0600)
//    IPPROTO_PUP           = 12,
//    IPPROTO_UDP           = 17,
//    IPPROTO_IDP           = 22,
//#if(_WIN32_WINNT >= 0x0600)
//    IPPROTO_RDP           = 27,
//#endif//(_WIN32_WINNT >= 0x0600)
//
//#if(_WIN32_WINNT >= 0x0501)
//    IPPROTO_IPV6          = 41, // IPv6 header
//    IPPROTO_ROUTING       = 43, // IPv6 Routing header
//    IPPROTO_FRAGMENT      = 44, // IPv6 fragmentation header
//    IPPROTO_ESP           = 50, // encapsulating security payload
//    IPPROTO_AH            = 51, // authentication header
//    IPPROTO_ICMPV6        = 58, // ICMPv6
//    IPPROTO_NONE          = 59, // IPv6 no next header
//    IPPROTO_DSTOPTS       = 60, // IPv6 Destination options
//#endif//(_WIN32_WINNT >= 0x0501)
//
//    IPPROTO_ND            = 77,
//#if(_WIN32_WINNT >= 0x0501)
//    IPPROTO_ICLFXBM       = 78,
//#endif//(_WIN32_WINNT >= 0x0501)
//#if(_WIN32_WINNT >= 0x0600)
//    IPPROTO_PIM           = 103,
//    IPPROTO_PGM           = 113,
//    IPPROTO_L2TP          = 115,
//    IPPROTO_SCTP          = 132,
//#endif//(_WIN32_WINNT >= 0x0600)
//    IPPROTO_RAW           = 255,
//
//    IPPROTO_MAX           = 256,
////
////  These are reserved for internal use by Windows.
////
//    IPPROTO_RESERVED_RAW  = 257,
//    IPPROTO_RESERVED_IPSEC  = 258,
//    IPPROTO_RESERVED_IPSECOFFLOAD  = 259,
//    IPPROTO_RESERVED_WNV = 260,
//    IPPROTO_RESERVED_MAX  = 261
//} IPPROTO, *PIPROTO;
const (
	//typedef enum {
	//#if(_WIN32_WINNT >= 0x0501)
	IPPROTO_HOPOPTS = 0 // IPv6 Hop-by-Hop options
	IPPROTO_ICMP    = 1
	IPPROTO_IGMP    = 2
	IPPROTO_GGP     = 3
	//#if(_WIN32_WINNT >= 0x0501)
	IPPROTO_IPV4 = 4
	//#endif//(_WIN32_WINNT >= 0x0501)
	//#if(_WIN32_WINNT >= 0x0600)
	IPPROTO_ST = 5
	//#endif//(_WIN32_WINNT >= 0x0600)
	IPPROTO_TCP = 6
	//#if(_WIN32_WINNT >= 0x0600)
	IPPROTO_CBT = 7
	IPPROTO_EGP = 8
	IPPROTO_IGP = 9
	//#endif//(_WIN32_WINNT >= 0x0600)
	IPPROTO_PUP = 12
	IPPROTO_UDP = 17
	IPPROTO_IDP = 22
	//#if(_WIN32_WINNT >= 0x0600)
	IPPROTO_RDP = 27
	//#endif//(_WIN32_WINNT >= 0x0600)

	//#if(_WIN32_WINNT >= 0x0501)
	IPPROTO_IPV6     = 41 // IPv6 header
	IPPROTO_ROUTING  = 43 // IPv6 Routing header
	IPPROTO_FRAGMENT = 44 // IPv6 fragmentation header
	IPPROTO_ESP      = 50 // encapsulating security payload
	IPPROTO_AH       = 51 // authentication header
	IPPROTO_ICMPV6   = 58 // ICMPv6
	IPPROTO_NONE     = 59 // IPv6 no next header
	IPPROTO_DSTOPTS  = 60 // IPv6 Destination options
	//#endif//(_WIN32_WINNT >= 0x0501)

	IPPROTO_ND = 77
	//#if(_WIN32_WINNT >= 0x0501)
	IPPROTO_ICLFXBM = 78
	//#endif//(_WIN32_WINNT >= 0x0501)
	//#if(_WIN32_WINNT >= 0x0600)
	IPPROTO_PIM  = 103
	IPPROTO_PGM  = 113
	IPPROTO_L2TP = 115
	IPPROTO_SCTP = 132
	//#endif//(_WIN32_WINNT >= 0x0600)
	IPPROTO_RAW = 255

	IPPROTO_MAX = 256
	//
	//  These are reserved for internal use by Windows.
	//
	IPPROTO_RESERVED_RAW          = 257
	IPPROTO_RESERVED_IPSEC        = 258
	IPPROTO_RESERVED_IPSECOFFLOAD = 259
	IPPROTO_RESERVED_WNV          = 260
	IPPROTO_RESERVED_MAX          = 261

//} IPPROTO, *PIPROTO;
)

//typedef /* [v1_enum] */
//enum FWP_MATCH_TYPE_
//{

//} 	FWP_MATCH_TYPE;

//typedef struct FWPM_SESSION0_
//{
//GUID sessionKey;
//FWPM_DISPLAY_DATA0 displayData;
//UINT32 flags;
//UINT32 txnWaitTimeoutInMSec;
//DWORD processId;
///* [unique] */ SID *sid;
///* [unique][string] */ wchar_t *username;
//BOOL kernelMode;
//} 	FWPM_SESSION0;
type FwpmSession0 struct {
	SessionKey           GUID
	DisplayData          FwpmDisplayData0
	Flags                FwpmSessionType
	TxnWaitTimeoutInMSec uint32
	ProcessId            uint32
	Sid                  *SID
	Username             *uint16
	KernelMode           int
}

//typedef struct FWPM_DISPLAY_DATA0_
//{
///* [unique][string] */ wchar_t *name;
///* [unique][string] */ wchar_t *description;
//} 	FWPM_DISPLAY_DATA0;

type FwpmDisplayData0 struct {
	Name        *uint16
	Description *uint16
}

//typedef struct FWPM_SUBLAYER0_
//{
//GUID subLayerKey;
//FWPM_DISPLAY_DATA0 displayData;
//UINT32 flags;
///* [unique] */ GUID *providerKey;
//FWP_BYTE_BLOB providerData;
//UINT16 weight;
//} 	FWPM_SUBLAYER0;
type FwpmSublayer0 struct {
	SubLayerKey  GUID
	DisplayData  FwpmDisplayData0
	Flags        FwpmSublayerFlag
	ProviderKey  *GUID
	ProviderData FwpByteBlob
	Weight       uint16 //权重
}

//typedef struct FWP_BYTE_BLOB_
//{
//UINT32 size;
///* [unique][size_is] */ UINT8 *data;
//} 	FWP_BYTE_BLOB;
type FwpByteBlob struct {
	Size uint32
	Data *byte // uint8
}

type PSecurityDescriptor unsafe.Pointer

type FilterId uint64

//typedef struct FWP_VALUE0_
//{
//FWP_DATA_TYPE type;
///* [switch_is][switch_type] */ union
//{
///* [case()] */  /* Empty union arm */
///* [case()] */ UINT8 uint8;
///* [case()] */ UINT16 uint16;
///* [case()] */ UINT32 uint32;
///* [case()][unique] */ UINT64 *uint64;
///* [case()] */ INT8 int8;
///* [case()] */ INT16 int16;
///* [case()] */ INT32 int32;
///* [case()][unique] */ INT64 *int64;
///* [case()] */ float float32;
///* [case()][unique] */ double *double64;
///* [case()][unique] */ FWP_BYTE_ARRAY16 *byteArray16;
///* [case()][unique] */ FWP_BYTE_BLOB *byteBlob;
///* [case()][unique] */ SID *sid;
///* [case()][unique] */ FWP_BYTE_BLOB *sd;
///* [case()][unique] */ FWP_TOKEN_INFORMATION *tokenInformation;
///* [case()][unique] */ FWP_BYTE_BLOB *tokenAccessInformation;
///* [case()][string] */ LPWSTR unicodeString;
///* [case()][unique] */ FWP_BYTE_ARRAY6 *byteArray6;
//} 	;
//} 	FWP_VALUE0;
type FwpValue0 struct {
	Type FwpDataType
	// 需要存在指针，必须随位数变化，所以是 int。但是指针 gc 需要注意，记得使用 runtime.KeepAlive 保留引用，防止被垃圾回收。
	Data int
}

func (fv *FwpValue0) SetUint8(v uint8) error {
	if fv.Type != FWP_UINT8 {
		return fmt.Errorf("Type %v != FWP_UINT8", fv.Type)
	}
	*((*uint8)(unsafe.Pointer(&fv.Data))) = v
	return nil
}

func (fv *FwpValue0) GetUint8() (uint8, error) {
	if fv.Type != FWP_UINT8 {
		return 0, fmt.Errorf("Type %v != FWP_UINT8", fv.Type)
	}
	return *((*uint8)(unsafe.Pointer(&fv.Data))), nil
}

func (fv *FwpValue0) SetUint16(v uint16) error {
	if fv.Type != FWP_UINT16 {
		return fmt.Errorf("Type %v != FWP_UINT16", fv.Type)
	}
	*((*uint16)(unsafe.Pointer(&fv.Data))) = v
	return nil
}
func (fv *FwpValue0) GetUint16() (uint16, error) {
	if fv.Type != FWP_UINT16 {
		return 0, fmt.Errorf("Type %v != FWP_UINT16", fv.Type)
	}
	return *((*uint16)(unsafe.Pointer(&fv.Data))), nil
}

//typedef struct FWPM_FILTER_CONDITION0_
//{
//GUID fieldKey;
//FWP_MATCH_TYPE matchType;
//FWP_CONDITION_VALUE0 conditionValue;
//} 	FWPM_FILTER_CONDITION0;
type FwpmFilterCondition0 struct {
	FieldKey       GUID
	MatchType      FwpMatchType
	ConditionValue FwpConditionValue0
}

//typedef struct FWP_CONDITION_VALUE0_
//{
//FWP_DATA_TYPE type;
///* [switch_is][switch_type] */ union
//{
///* [case()] */  /* Empty union arm */
///* [case()] */ UINT8 uint8;
///* [case()] */ UINT16 uint16;
///* [case()] */ UINT32 uint32;
///* [case()][unique] */ UINT64 *uint64;
///* [case()] */ INT8 int8;
///* [case()] */ INT16 int16;
///* [case()] */ INT32 int32;
///* [case()][unique] */ INT64 *int64;
///* [case()] */ float float32;
///* [case()][unique] */ double *double64;
///* [case()][unique] */ FWP_BYTE_ARRAY16 *byteArray16;
///* [case()][unique] */ FWP_BYTE_BLOB *byteBlob;
///* [case()][unique] */ SID *sid;
///* [case()][unique] */ FWP_BYTE_BLOB *sd;
///* [case()][unique] */ FWP_TOKEN_INFORMATION *tokenInformation;
///* [case()][unique] */ FWP_BYTE_BLOB *tokenAccessInformation;
///* [case()][string] */ LPWSTR unicodeString;
///* [case()][unique] */ FWP_BYTE_ARRAY6 *byteArray6;
///* [case()][unique] */ FWP_V4_ADDR_AND_MASK *v4AddrMask;
///* [case()][unique] */ FWP_V6_ADDR_AND_MASK *v6AddrMask;
///* [case()][unique] */ FWP_RANGE0 *rangeValue;
//} 	;
//} 	FWP_CONDITION_VALUE0;
type FwpConditionValue0 struct {
	Type FwpDataType
	Data uint
}

func (fv *FwpConditionValue0) SetUint8(v uint8) error {
	if fv.Type != FWP_UINT8 {
		return fmt.Errorf("Type %v != FWP_UINT8", fv.Type)
	}
	*((*uint8)(unsafe.Pointer(&fv.Data))) = v
	return nil
}

func (fv *FwpConditionValue0) GetUint8() (uint8, error) {
	if fv.Type != FWP_UINT8 {
		return 0, fmt.Errorf("Type %v != FWP_UINT8", fv.Type)
	}
	return *((*uint8)(unsafe.Pointer(&fv.Data))), nil
}

func (fv *FwpConditionValue0) SetUint16(v uint16) error {
	if fv.Type != FWP_UINT16 {
		return fmt.Errorf("Type %v != FWP_UINT16", fv.Type)
	}
	*((*uint16)(unsafe.Pointer(&fv.Data))) = v
	return nil
}
func (fv *FwpConditionValue0) GetUint16() (uint16, error) {
	if fv.Type != FWP_UINT16 {
		return 0, fmt.Errorf("Type %v != FWP_UINT16", fv.Type)
	}
	return *((*uint16)(unsafe.Pointer(&fv.Data))), nil
}

// 注意，FwpConditionValue0 不会保留指针的引用，调用者记得执行 runtime.KeepAlive 防止 v 被垃圾回收。
func (fv *FwpConditionValue0) SetPUint64(v *uint64) error {
	if fv.Type != FWP_UINT64 {
		return fmt.Errorf("Type %v != FWP_UINT64", fv.Type)
	}

	fv.Data = uint(uintptr(unsafe.Pointer(v)))
	return nil
}

// 注意：垃圾回收需要自己注意！
func (fv *FwpConditionValue0) GetPUint64() (*uint64, error) {
	if fv.Type != FWP_UINT64 {
		return nil, fmt.Errorf("Type %v != FWP_UINT64", fv.Type)
	}
	return (*uint64)(unsafe.Pointer(uintptr(fv.Data))), nil
}

// 注意，FwpConditionValue0 不会保留指针的引用，调用者记得执行 runtime.KeepAlive 防止 v 被垃圾回收。
func (fv *FwpConditionValue0) SetByteBlob(v *FwpByteBlob) error {
	if fv.Type != FWP_BYTE_BLOB_TYPE {
		return fmt.Errorf("Type %v != FWP_BYTE_BLOB_TYPE", fv.Type)
	}
	fv.Data = uint(uintptr(unsafe.Pointer(v)))
	return nil
}

// 注意，垃圾回收需要自己注意！
func (fv *FwpConditionValue0) GetByteBlob() (*FwpByteBlob, error) {
	if fv.Type != FWP_BYTE_BLOB_TYPE {
		return nil, fmt.Errorf("Type %v != FWP_BYTE_BLOB_TYPE", fv.Type)
	}
	return (*FwpByteBlob)(unsafe.Pointer(uintptr(fv.Data))), nil
}

//typedef struct FWPM_ACTION0_
//{
//FWP_ACTION_TYPE type;
///* [switch_is] */ /* [switch_type] */ union
//{
///* [case()] */ GUID filterType;
///* [case()] */ GUID calloutKey;
//} 	;
//} 	FWPM_ACTION0;
type FwpmAction0 struct {
	Type                   FwpActionType
	FilterTypeOrCalloutKey GUID
}
