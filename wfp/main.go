package main

import (
	"fmt"

	"github.com/gamexg/gowindows"
	"golang.org/x/sys/windows"
)

const FIREWALL_SUBLAYER_NAMEW = "aaa1"
const FIREWALL_SERVICE_NAMEW = "aaa1.1"

func main() {
	err := t(100, true)
	if err != nil {
		panic(err)
	}

	err = t(300, false)
	if err != nil {
		panic(err)
	}
	fmt.Print("等待退出\r\n")
	var s string
	fmt.Scanln(&s)
}

func t(Weight uint16, block bool) error {
	engineHandle := gowindows.Handle(0)

	session := gowindows.FwpmSession0{
		Flags: gowindows.FWPM_SESSION_FLAG_DYNAMIC,
	}

	err := gowindows.FwpmEngineOpen0("", gowindows.RPC_C_AUTHN_WINNT, nil, &session, &engineHandle)
	if err != nil {
		return fmt.Errorf("FwpmEngineOpen0,%v", err)
	}

	subLayer := gowindows.FwpmSublayer0{}
	subLayer.DisplayData.Name = windows.StringToUTF16Ptr(FIREWALL_SUBLAYER_NAMEW)
	subLayer.DisplayData.Description = windows.StringToUTF16Ptr(FIREWALL_SUBLAYER_NAMEW)
	subLayer.Flags = 0
	subLayer.Weight = Weight

	err = gowindows.UuidCreate(&subLayer.SubLayerKey)
	if err != nil {
		return fmt.Errorf("UuidCreate ,%v", err)
	}

	err = gowindows.FwpmSubLayerAdd0(engineHandle, &subLayer, nil)
	if err != nil {
		return fmt.Errorf("FwpmSubLayerAdd0, %v", err)
	}

	filter := gowindows.FwpmFilter0{}
	condition := make([]gowindows.FwpmFilterCondition0, 2)

	filter.SubLayerKey = subLayer.SubLayerKey
	filter.DisplayData.Name = windows.StringToUTF16Ptr(FIREWALL_SERVICE_NAMEW)
	filter.Weight.Type = gowindows.FWP_UINT8
	filter.Weight.SetUint8(0xF)
	filter.FilterCondition = &condition[0]
	filter.NumFilterConditions = uint32(len(condition))

	condition[0].FieldKey = gowindows.FWPM_CONDITION_IP_REMOTE_PORT
	condition[0].MatchType = gowindows.FWP_MATCH_EQUAL
	condition[0].ConditionValue.Type = gowindows.FWP_UINT16
	condition[0].ConditionValue.SetUint16(80)

	// 拦截 IPv4 所有 DNS 请求
	filter.LayerKey = gowindows.FWPM_LAYER_ALE_AUTH_CONNECT_V4
	filter.Action.Type = gowindows.FWP_ACTION_BLOCK
	filter.Weight.Type = gowindows.FWP_EMPTY
	filter.NumFilterConditions = 1

	if block == false {
		filter.Action.Type = gowindows.FWP_ACTION_PERMIT
	}

	var filterId gowindows.FilterId
	err = gowindows.FwpmFilterAdd0(engineHandle, &filter, nil, &filterId)
	if err != nil {
		return fmt.Errorf("ipv4-FwpmFilterAdd0, %v", err)
	}

	return nil
}
