package gowindows

type LUID struct {
	LowPart  uint32
	HighPart int32
}

type LUID_AND_ATTRIBUTES struct {
	Luid       LUID
	Attributes uint32
}

type TOKEN_PRIVILEGES struct {
	PrivilegeCount uint32
	Privileges     [1]LUID_AND_ATTRIBUTES
}

const (
	SE_PRIVILEGE_ENABLED = 0x00000002
)

type Privilege string

const (
	////////////////////////////////////////////////////////////////////////
	//                                                                    //
	//               NT Defined Privileges                                //
	//                                                                    //
	////////////////////////////////////////////////////////////////////////

	SE_CREATE_TOKEN_NAME           Privilege = "SeCreateTokenPrivilege"
	SE_ASSIGNPRIMARYTOKEN_NAME     Privilege = "SeAssignPrimaryTokenPrivilege"
	SE_LOCK_MEMORY_NAME            Privilege = "SeLockMemoryPrivilege"
	SE_INCREASE_QUOTA_NAME         Privilege = "SeIncreaseQuotaPrivilege"
	SE_UNSOLICITED_INPUT_NAME      Privilege = "SeUnsolicitedInputPrivilege"
	SE_MACHINE_ACCOUNT_NAME        Privilege = "SeMachineAccountPrivilege"
	SE_TCB_NAME                    Privilege = "SeTcbPrivilege"
	SE_SECURITY_NAME               Privilege = "SeSecurityPrivilege"
	SE_TAKE_OWNERSHIP_NAME         Privilege = "SeTakeOwnershipPrivilege"
	SE_LOAD_DRIVER_NAME            Privilege = "SeLoadDriverPrivilege"
	SE_SYSTEM_PROFILE_NAME         Privilege = "SeSystemProfilePrivilege"
	SE_SYSTEMTIME_NAME             Privilege = "SeSystemtimePrivilege"
	SE_PROF_SINGLE_PROCESS_NAME    Privilege = "SeProfileSingleProcessPrivilege"
	SE_INC_BASE_PRIORITY_NAME      Privilege = "SeIncreaseBasePriorityPrivilege"
	SE_CREATE_PAGEFILE_NAME        Privilege = "SeCreatePagefilePrivilege"
	SE_CREATE_PERMANENT_NAME       Privilege = "SeCreatePermanentPrivilege"
	SE_BACKUP_NAME                 Privilege = "SeBackupPrivilege"
	SE_RESTORE_NAME                Privilege = "SeRestorePrivilege"
	SE_SHUTDOWN_NAME               Privilege = "SeShutdownPrivilege"
	SE_DEBUG_NAME                  Privilege = "SeDebugPrivilege"
	SE_AUDIT_NAME                  Privilege = "SeAuditPrivilege"
	SE_SYSTEM_ENVIRONMENT_NAME     Privilege = "SeSystemEnvironmentPrivilege"
	SE_CHANGE_NOTIFY_NAME          Privilege = "SeChangeNotifyPrivilege"
	SE_REMOTE_SHUTDOWN_NAME        Privilege = "SeRemoteShutdownPrivilege"
	SE_UNDOCK_NAME                 Privilege = "SeUndockPrivilege"
	SE_SYNC_AGENT_NAME             Privilege = "SeSyncAgentPrivilege"
	SE_ENABLE_DELEGATION_NAME      Privilege = "SeEnableDelegationPrivilege"
	SE_MANAGE_VOLUME_NAME          Privilege = "SeManageVolumePrivilege"
	SE_IMPERSONATE_NAME            Privilege = "SeImpersonatePrivilege"
	SE_CREATE_GLOBAL_NAME          Privilege = "SeCreateGlobalPrivilege"
	SE_TRUSTED_CREDMAN_ACCESS_NAME Privilege = "SeTrustedCredManAccessPrivilege"
	SE_RELABEL_NAME                Privilege = "SeRelabelPrivilege"
	SE_INC_WORKING_SET_NAME        Privilege = "SeIncreaseWorkingSetPrivilege"
	SE_TIME_ZONE_NAME              Privilege = "SeTimeZonePrivilege"
	SE_CREATE_SYMBOLIC_LINK_NAME   Privilege = "SeCreateSymbolicLinkPrivilege"
)

type SddlRevision DWord

const (
	SDDL_REVISION_1 = 1
	SDDL_REVISION   = SDDL_REVISION_1
)

//typedef PVOID PSECURITY_DESCRIPTOR;
type SecurityDescriptor Pointer

/*
typedef struct _ACL {
    BYTE  AclRevision;
    BYTE  Sbz1;
    WORD   AclSize;
    WORD   AceCount;
    WORD   Sbz2;
} ACL;
typedef ACL *PACL;
*/
type ACL struct {
	AclRevision byte
	Sbz1        byte
	AclSize     Word
	AceCount    Word
	Sbz2        Word
}

const LOW_INTEGRITY_SDDL_SACL_W = "S:(ML;;NW;;;LW)"

const (
	SE_UNKNOWN_OBJECT_TYPE SeObjectType = iota + 0
	SE_FILE_OBJECT
	SE_SERVICE
	SE_PRINTER
	SE_REGISTRY_KEY
	SE_LMSHARE
	SE_KERNEL_OBJECT
	SE_WINDOW_OBJECT
	SE_DS_OBJECT
	SE_DS_OBJECT_ALL
	SE_PROVIDER_DEFINED_OBJECT
	SE_WMIGUID_OBJECT
	SE_REGISTRY_WOW64_32KEY
)

type SeObjectType int32

const (
	OWNER_SECURITY_INFORMATION               SecurityInformation = (0x00000001)
	GROUP_SECURITY_INFORMATION               SecurityInformation = (0x00000002)
	DACL_SECURITY_INFORMATION                SecurityInformation = (0x00000004)
	SACL_SECURITY_INFORMATION                SecurityInformation = (0x00000008)
	LABEL_SECURITY_INFORMATION               SecurityInformation = (0x00000010)
	ATTRIBUTE_SECURITY_INFORMATION           SecurityInformation = (0x00000020)
	SCOPE_SECURITY_INFORMATION               SecurityInformation = (0x00000040)
	PROCESS_TRUST_LABEL_SECURITY_INFORMATION SecurityInformation = (0x00000080)
	BACKUP_SECURITY_INFORMATION              SecurityInformation = (0x00010000)

	PROTECTED_DACL_SECURITY_INFORMATION   SecurityInformation = (0x80000000)
	PROTECTED_SACL_SECURITY_INFORMATION   SecurityInformation = (0x40000000)
	UNPROTECTED_DACL_SECURITY_INFORMATION SecurityInformation = (0x20000000)
	UNPROTECTED_SACL_SECURITY_INFORMATION SecurityInformation = (0x10000000)
)

type SecurityInformation DWord
type PSId Pointer
