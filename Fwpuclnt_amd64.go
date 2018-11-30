package gowindows

//typedef struct FWPM_FILTER0_
//{
//GUID filterKey;
//FWPM_DISPLAY_DATA0 displayData;
//UINT32 flags;
///* [unique] */ GUID *providerKey;
//FWP_BYTE_BLOB providerData;
//GUID layerKey;
//GUID subLayerKey;
//FWP_VALUE0 weight;
//UINT32 numFilterConditions;
///* [unique][size_is] */ FWPM_FILTER_CONDITION0 *filterCondition;
//FWPM_ACTION0 action;
///* [switch_is] */ /* [switch_type] */ union
//{
///* [case()] */ UINT64 rawContext;
///* [case()] */ GUID providerContextKey;
//} 	;
///* [unique] */ GUID *reserved;
//UINT64 filterId;
//FWP_VALUE0 effectiveWeight;
//} 	FWPM_FILTER0;
type FwpmFilter0 struct {
	FilterKey           GUID
	DisplayData         FwpmDisplayData0
	Flags               uint32
	ProviderKey         *GUID
	ProviderData        FwpByteBlob
	LayerKey            GUID
	SubLayerKey         GUID
	Weight              FwpValue0
	NumFilterConditions uint32
	FilterCondition     *FwpmFilterCondition0
	Action              FwpmAction0
	ProviderContextKey  GUID // 另一个可能，UINT64 rawContext
	Reserved            *GUID
	FilterId            uint64
	EffectiveWeight     FwpValue0
}
