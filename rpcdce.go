package gowindows

// Authentication-Service Constants
// The authentication service constants represent the authentication services passed to various run-time functions.
//https://docs.microsoft.com/en-us/windows/desktop/rpc/authentication-service-constants
type RpcCAuthnType uint32

const (
	RPC_C_AUTHN_NONE          RpcCAuthnType = 0
	RPC_C_AUTHN_DCE_PRIVATE   RpcCAuthnType = 1
	RPC_C_AUTHN_DCE_PUBLIC    RpcCAuthnType = 2
	RPC_C_AUTHN_DEC_PUBLIC    RpcCAuthnType = 4
	RPC_C_AUTHN_GSS_NEGOTIATE RpcCAuthnType = 9
	RPC_C_AUTHN_WINNT         RpcCAuthnType = 10
	RPC_C_AUTHN_GSS_SCHANNEL  RpcCAuthnType = 14
	RPC_C_AUTHN_GSS_KERBEROS  RpcCAuthnType = 16
	RPC_C_AUTHN_DPA           RpcCAuthnType = 17
	RPC_C_AUTHN_MSN           RpcCAuthnType = 18
)
