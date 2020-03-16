package gowindows

import "fmt"

type _HRESULT_TYPEDEF_ HRESULT

//
// =======================================================
// Windows Filtering Platform Error Messages
// =======================================================
//
//
// MessageId: FWP_E_CALLOUT_NOT_FOUND
//
// MessageText:
//
// The callout does not exist.
//
const FWP_E_CALLOUT_NOT_FOUND = _HRESULT_TYPEDEF_(uint32((0x80320001)))

//
// MessageId: FWP_E_CONDITION_NOT_FOUND
//
// MessageText:
//
// The filter condition does not exist.
//
const FWP_E_CONDITION_NOT_FOUND = _HRESULT_TYPEDEF_(uint32((0x80320002)))

//
// MessageId: FWP_E_FILTER_NOT_FOUND
//
// MessageText:
//
// The filter does not exist.
//
const FWP_E_FILTER_NOT_FOUND = _HRESULT_TYPEDEF_(uint32((0x80320003)))

//
// MessageId: FWP_E_LAYER_NOT_FOUND
//
// MessageText:
//
// The layer does not exist.
//
const FWP_E_LAYER_NOT_FOUND = _HRESULT_TYPEDEF_(uint32((0x80320004)))

//
// MessageId: FWP_E_PROVIDER_NOT_FOUND
//
// MessageText:
//
// The provider does not exist.
//
const FWP_E_PROVIDER_NOT_FOUND = _HRESULT_TYPEDEF_(uint32((0x80320005)))

//
// MessageId: FWP_E_PROVIDER_CONTEXT_NOT_FOUND
//
// MessageText:
//
// The provider context does not exist.
//
const FWP_E_PROVIDER_CONTEXT_NOT_FOUND = _HRESULT_TYPEDEF_(uint32((0x80320006)))

//
// MessageId: FWP_E_SUBLAYER_NOT_FOUND
//
// MessageText:
//
// The sublayer does not exist.
//
const FWP_E_SUBLAYER_NOT_FOUND = _HRESULT_TYPEDEF_(uint32((0x80320007)))

//
// MessageId: FWP_E_NOT_FOUND
//
// MessageText:
//
// The object does not exist.
//
const FWP_E_NOT_FOUND = _HRESULT_TYPEDEF_(uint32((0x80320008)))

//
// MessageId: FWP_E_ALREADY_EXISTS
//
// MessageText:
//
// An object with that GUID or LUID already exists.
//
const FWP_E_ALREADY_EXISTS = _HRESULT_TYPEDEF_(uint32((0x80320009)))

//
// MessageId: FWP_E_IN_USE
//
// MessageText:
//
// The object is referenced by other objects so cannot be deleted.
//
const FWP_E_IN_USE = _HRESULT_TYPEDEF_(uint32((0x8032000A)))

//
// MessageId: FWP_E_DYNAMIC_SESSION_IN_PROGRESS
//
// MessageText:
//
// The call is not allowed from within a dynamic session.
//
const FWP_E_DYNAMIC_SESSION_IN_PROGRESS = _HRESULT_TYPEDEF_(uint32((0x8032000B)))

//
// MessageId: FWP_E_WRONG_SESSION
//
// MessageText:
//
// The call was made from the wrong session so cannot be completed.
//
const FWP_E_WRONG_SESSION = _HRESULT_TYPEDEF_(uint32((0x8032000C)))

//
// MessageId: FWP_E_NO_TXN_IN_PROGRESS
//
// MessageText:
//
// The call must be made from within an explicit transaction.
//
const FWP_E_NO_TXN_IN_PROGRESS = _HRESULT_TYPEDEF_(uint32((0x8032000D)))

//
// MessageId: FWP_E_TXN_IN_PROGRESS
//
// MessageText:
//
// The call is not allowed from within an explicit transaction.
//
const FWP_E_TXN_IN_PROGRESS = _HRESULT_TYPEDEF_(uint32((0x8032000E)))

//
// MessageId: FWP_E_TXN_ABORTED
//
// MessageText:
//
// The explicit transaction has been forcibly cancelled.
//
const FWP_E_TXN_ABORTED = _HRESULT_TYPEDEF_(uint32((0x8032000F)))

//
// MessageId: FWP_E_SESSION_ABORTED
//
// MessageText:
//
// The session has been cancelled.
//
const FWP_E_SESSION_ABORTED = _HRESULT_TYPEDEF_(uint32((0x80320010)))

//
// MessageId: FWP_E_INCOMPATIBLE_TXN
//
// MessageText:
//
// The call is not allowed from within a read-only transaction.
//
const FWP_E_INCOMPATIBLE_TXN = _HRESULT_TYPEDEF_(uint32((0x80320011)))

//
// MessageId: FWP_E_TIMEOUT
//
// MessageText:
//
// The call timed out while waiting to acquire the transaction lock.
//
const FWP_E_TIMEOUT = _HRESULT_TYPEDEF_(uint32((0x80320012)))

//
// MessageId: FWP_E_NET_EVENTS_DISABLED
//
// MessageText:
//
// Collection of network diagnostic events is disabled.
//
const FWP_E_NET_EVENTS_DISABLED = _HRESULT_TYPEDEF_(uint32((0x80320013)))

//
// MessageId: FWP_E_INCOMPATIBLE_LAYER
//
// MessageText:
//
// The operation is not supported by the specified layer.
//
const FWP_E_INCOMPATIBLE_LAYER = _HRESULT_TYPEDEF_(uint32((0x80320014)))

//
// MessageId: FWP_E_KM_CLIENTS_ONLY
//
// MessageText:
//
// The call is allowed for kernel-mode callers only.
//
const FWP_E_KM_CLIENTS_ONLY = _HRESULT_TYPEDEF_(uint32((0x80320015)))

//
// MessageId: FWP_E_LIFETIME_MISMATCH
//
// MessageText:
//
// The call tried to associate two objects with incompatible lifetimes.
//
const FWP_E_LIFETIME_MISMATCH = _HRESULT_TYPEDEF_(uint32((0x80320016)))

//
// MessageId: FWP_E_BUILTIN_OBJECT
//
// MessageText:
//
// The object is built in so cannot be deleted.
//
const FWP_E_BUILTIN_OBJECT = _HRESULT_TYPEDEF_(uint32((0x80320017)))

//
// MessageId: FWP_E_TOO_MANY_CALLOUTS
//
// MessageText:
//
// The maximum number of callouts has been reached.
//
const FWP_E_TOO_MANY_CALLOUTS = _HRESULT_TYPEDEF_(uint32((0x80320018)))

//
// MessageId: FWP_E_NOTIFICATION_DROPPED
//
// MessageText:
//
// A notification could not be delivered because a message queue is at its maximum capacity.
//
const FWP_E_NOTIFICATION_DROPPED = _HRESULT_TYPEDEF_(uint32((0x80320019)))

//
// MessageId: FWP_E_TRAFFIC_MISMATCH
//
// MessageText:
//
// The traffic parameters do not match those for the security association context.
//
const FWP_E_TRAFFIC_MISMATCH = _HRESULT_TYPEDEF_(uint32((0x8032001A)))

//
// MessageId: FWP_E_INCOMPATIBLE_SA_STATE
//
// MessageText:
//
// The call is not allowed for the current security association state.
//
const FWP_E_INCOMPATIBLE_SA_STATE = _HRESULT_TYPEDEF_(uint32((0x8032001B)))

//
// MessageId: FWP_E_NULL_POINTER
//
// MessageText:
//
// A required pointer is null.
//
const FWP_E_NULL_POINTER = _HRESULT_TYPEDEF_(uint32((0x8032001C)))

//
// MessageId: FWP_E_INVALID_ENUMERATOR
//
// MessageText:
//
// An enumerator is not valid.
//
const FWP_E_INVALID_ENUMERATOR = _HRESULT_TYPEDEF_(uint32((0x8032001D)))

//
// MessageId: FWP_E_INVALID_FLAGS
//
// MessageText:
//
// The flags field contains an invalid value.
//
const FWP_E_INVALID_FLAGS = _HRESULT_TYPEDEF_(uint32((0x8032001E)))

//
// MessageId: FWP_E_INVALID_NET_MASK
//
// MessageText:
//
// A network mask is not valid.
//
const FWP_E_INVALID_NET_MASK = _HRESULT_TYPEDEF_(uint32((0x8032001F)))

//
// MessageId: FWP_E_INVALID_RANGE
//
// MessageText:
//
// An FWP_RANGE is not valid.
//
const FWP_E_INVALID_RANGE = _HRESULT_TYPEDEF_(uint32((0x80320020)))

//
// MessageId: FWP_E_INVALID_INTERVAL
//
// MessageText:
//
// The time interval is not valid.
//
const FWP_E_INVALID_INTERVAL = _HRESULT_TYPEDEF_(uint32((0x80320021)))

//
// MessageId: FWP_E_ZERO_LENGTH_ARRAY
//
// MessageText:
//
// An array that must contain at least one element is zero length.
//
const FWP_E_ZERO_LENGTH_ARRAY = _HRESULT_TYPEDEF_(uint32((0x80320022)))

//
// MessageId: FWP_E_NULL_DISPLAY_NAME
//
// MessageText:
//
// The displayData.name field cannot be null.
//
const FWP_E_NULL_DISPLAY_NAME = _HRESULT_TYPEDEF_(uint32((0x80320023)))

//
// MessageId: FWP_E_INVALID_ACTION_TYPE
//
// MessageText:
//
// The action type is not one of the allowed action types for a filter.
//
const FWP_E_INVALID_ACTION_TYPE = _HRESULT_TYPEDEF_(uint32((0x80320024)))

//
// MessageId: FWP_E_INVALID_WEIGHT
//
// MessageText:
//
// The filter weight is not valid.
//
const FWP_E_INVALID_WEIGHT = _HRESULT_TYPEDEF_(uint32((0x80320025)))

//
// MessageId: FWP_E_MATCH_TYPE_MISMATCH
//
// MessageText:
//
// A filter condition contains a match type that is not compatible with the operands.
//
const FWP_E_MATCH_TYPE_MISMATCH = _HRESULT_TYPEDEF_(uint32((0x80320026)))

//
// MessageId: FWP_E_TYPE_MISMATCH
//
// MessageText:
//
// An FWP_VALUE or FWPM_CONDITION_VALUE is of the wrong type.
//
const FWP_E_TYPE_MISMATCH = _HRESULT_TYPEDEF_(uint32((0x80320027)))

//
// MessageId: FWP_E_OUT_OF_BOUNDS
//
// MessageText:
//
// An integer value is outside the allowed range.
//
const FWP_E_OUT_OF_BOUNDS = _HRESULT_TYPEDEF_(uint32((0x80320028)))

//
// MessageId: FWP_E_RESERVED
//
// MessageText:
//
// A reserved field is non-zero.
//
const FWP_E_RESERVED = _HRESULT_TYPEDEF_(uint32((0x80320029)))

//
// MessageId: FWP_E_DUPLICATE_CONDITION
//
// MessageText:
//
// A filter cannot contain multiple conditions operating on a single field.
//
const FWP_E_DUPLICATE_CONDITION = _HRESULT_TYPEDEF_(uint32((0x8032002A)))

//
// MessageId: FWP_E_DUPLICATE_KEYMOD
//
// MessageText:
//
// A policy cannot contain the same keying module more than once.
//
const FWP_E_DUPLICATE_KEYMOD = _HRESULT_TYPEDEF_(uint32((0x8032002B)))

//
// MessageId: FWP_E_ACTION_INCOMPATIBLE_WITH_LAYER
//
// MessageText:
//
// The action type is not compatible with the layer.
//
const FWP_E_ACTION_INCOMPATIBLE_WITH_LAYER = _HRESULT_TYPEDEF_(uint32((0x8032002C)))

//
// MessageId: FWP_E_ACTION_INCOMPATIBLE_WITH_SUBLAYER
//
// MessageText:
//
// The action type is not compatible with the sublayer.
//
const FWP_E_ACTION_INCOMPATIBLE_WITH_SUBLAYER = _HRESULT_TYPEDEF_(uint32((0x8032002D)))

//
// MessageId: FWP_E_CONTEXT_INCOMPATIBLE_WITH_LAYER
//
// MessageText:
//
// The raw context or the provider context is not compatible with the layer.
//
const FWP_E_CONTEXT_INCOMPATIBLE_WITH_LAYER = _HRESULT_TYPEDEF_(uint32((0x8032002E)))

//
// MessageId: FWP_E_CONTEXT_INCOMPATIBLE_WITH_CALLOUT
//
// MessageText:
//
// The raw context or the provider context is not compatible with the callout.
//
const FWP_E_CONTEXT_INCOMPATIBLE_WITH_CALLOUT = _HRESULT_TYPEDEF_(uint32((0x8032002F)))

//
// MessageId: FWP_E_INCOMPATIBLE_AUTH_METHOD
//
// MessageText:
//
// The authentication method is not compatible with the policy type.
//
const FWP_E_INCOMPATIBLE_AUTH_METHOD = _HRESULT_TYPEDEF_(uint32((0x80320030)))

//
// MessageId: FWP_E_INCOMPATIBLE_DH_GROUP
//
// MessageText:
//
// The Diffie-Hellman group is not compatible with the policy type.
//
const FWP_E_INCOMPATIBLE_DH_GROUP = _HRESULT_TYPEDEF_(uint32((0x80320031)))

//
// MessageId: FWP_E_EM_NOT_SUPPORTED
//
// MessageText:
//
// An IKE policy cannot contain an Extended Mode policy.
//
const FWP_E_EM_NOT_SUPPORTED = _HRESULT_TYPEDEF_(uint32((0x80320032)))

//
// MessageId: FWP_E_NEVER_MATCH
//
// MessageText:
//
// The enumeration template or subscription will never match any objects.
//
const FWP_E_NEVER_MATCH = _HRESULT_TYPEDEF_(uint32((0x80320033)))

//
// MessageId: FWP_E_PROVIDER_CONTEXT_MISMATCH
//
// MessageText:
//
// The provider context is of the wrong type.
//
const FWP_E_PROVIDER_CONTEXT_MISMATCH = _HRESULT_TYPEDEF_(uint32((0x80320034)))

//
// MessageId: FWP_E_INVALID_PARAMETER
//
// MessageText:
//
// The parameter is incorrect.
//
const FWP_E_INVALID_PARAMETER = _HRESULT_TYPEDEF_(uint32((0x80320035)))

//
// MessageId: FWP_E_TOO_MANY_SUBLAYERS
//
// MessageText:
//
// The maximum number of sublayers has been reached.
//
const FWP_E_TOO_MANY_SUBLAYERS = _HRESULT_TYPEDEF_(uint32((0x80320036)))

//
// MessageId: FWP_E_CALLOUT_NOTIFICATION_FAILED
//
// MessageText:
//
// The notification function for a callout returned an error.
//
const FWP_E_CALLOUT_NOTIFICATION_FAILED = _HRESULT_TYPEDEF_(uint32((0x80320037)))

//
// MessageId: FWP_E_INVALID_AUTH_TRANSFORM
//
// MessageText:
//
// The IPsec authentication transform is not valid.
//
const FWP_E_INVALID_AUTH_TRANSFORM = _HRESULT_TYPEDEF_(uint32((0x80320038)))

//
// MessageId: FWP_E_INVALID_CIPHER_TRANSFORM
//
// MessageText:
//
// The IPsec cipher transform is not valid.
//
const FWP_E_INVALID_CIPHER_TRANSFORM = _HRESULT_TYPEDEF_(uint32((0x80320039)))

//
// MessageId: FWP_E_INCOMPATIBLE_CIPHER_TRANSFORM
//
// MessageText:
//
// The IPsec cipher transform is not compatible with the policy.
//
const FWP_E_INCOMPATIBLE_CIPHER_TRANSFORM = _HRESULT_TYPEDEF_(uint32((0x8032003A)))

//
// MessageId: FWP_E_INVALID_TRANSFORM_COMBINATION
//
// MessageText:
//
// The combination of IPsec transform types is not valid.
//
const FWP_E_INVALID_TRANSFORM_COMBINATION = _HRESULT_TYPEDEF_(uint32((0x8032003B)))

//
// MessageId: FWP_E_DUPLICATE_AUTH_METHOD
//
// MessageText:
//
// A policy cannot contain the same auth method more than once.
//
const FWP_E_DUPLICATE_AUTH_METHOD = _HRESULT_TYPEDEF_(uint32((0x8032003C)))

//
// MessageId: FWP_E_INVALID_TUNNEL_ENDPOINT
//
// MessageText:
//
// A tunnel endpoint configuration is invalid.
//
const FWP_E_INVALID_TUNNEL_ENDPOINT = _HRESULT_TYPEDEF_(uint32((0x8032003D)))

//
// MessageId: FWP_E_L2_DRIVER_NOT_READY
//
// MessageText:
//
// The WFP MAC Layers are not ready.
//
const FWP_E_L2_DRIVER_NOT_READY = _HRESULT_TYPEDEF_(uint32((0x8032003E)))

//
// MessageId: FWP_E_KEY_DICTATOR_ALREADY_REGISTERED
//
// MessageText:
//
// A key manager capable of key dictation is already registered
//
const FWP_E_KEY_DICTATOR_ALREADY_REGISTERED = _HRESULT_TYPEDEF_(uint32((0x8032003F)))

//
// MessageId: FWP_E_KEY_DICTATION_INVALID_KEYING_MATERIAL
//
// MessageText:
//
// A key manager dictated invalid keys
//
const FWP_E_KEY_DICTATION_INVALID_KEYING_MATERIAL = _HRESULT_TYPEDEF_(uint32((0x80320040)))

//
// MessageId: FWP_E_CONNECTIONS_DISABLED
//
// MessageText:
//
// The BFE IPsec Connection Tracking is disabled.
//
const FWP_E_CONNECTIONS_DISABLED = _HRESULT_TYPEDEF_(uint32((0x80320041)))

//
// MessageId: FWP_E_INVALID_DNS_NAME
//
// MessageText:
//
// The DNS name is invalid.
//
const FWP_E_INVALID_DNS_NAME = _HRESULT_TYPEDEF_(uint32((0x80320042)))

//
// MessageId: FWP_E_STILL_ON
//
// MessageText:
//
// The engine option is still enabled due to other configuration settings.
//
const FWP_E_STILL_ON = _HRESULT_TYPEDEF_(uint32((0x80320043)))

//
// MessageId: FWP_E_IKEEXT_NOT_RUNNING
//
// MessageText:
//
// The IKEEXT service is not running.  This service only runs when there is IPsec policy applied to the machine.
//
const FWP_E_IKEEXT_NOT_RUNNING = _HRESULT_TYPEDEF_(uint32((0x80320044)))

//
// MessageId: FWP_E_DROP_NOICMP
//
// MessageText:
//
// The packet should be dropped, no ICMP should be sent.
//
const FWP_E_DROP_NOICMP = _HRESULT_TYPEDEF_(uint32((0x80320104)))

var (
	ERR_FWP_E_ALREADY_EXISTS error = &FwpmError{r1: FWP_E_ALREADY_EXISTS}
)

type FwpmError struct {
	r1 _HRESULT_TYPEDEF_
}

func newFwpmError(r1 _HRESULT_TYPEDEF_) error {
	switch r1 {
	case FWP_E_ALREADY_EXISTS:
		return ERR_FWP_E_ALREADY_EXISTS
	default:
		return &FwpmError{r1: r1}
	}
}

func (e *FwpmError) Error() string {
	return fmt.Sprintf("r1:%X", e.r1)
}
