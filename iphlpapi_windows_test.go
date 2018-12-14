package gowindows

import (
	"testing"
	"unsafe"

	"golang.org/x/sys/windows"
)

func TestAdapterAddresses(t *testing.T) {
	as, err := AdapterAddresses()
	if err != nil {
		t.Fatal(err)
	}

	if len(as) <= 0 {
		t.Fatalf("无网卡")
	}

	for _, aa := range as {
		//fmt.Printf("\r\nFriendlyName:%v\r\n", aa.GetFriendlyName())
		//fmt.Printf("IfType:%v\r\n", aa.IfType)
		//fmt.Printf("Description:%v\r\n", aa.GetDescription())

		luid, err := aa.GetLuid()
		if err != nil {
			t.Error(err)
		}
		_ = luid
		//fmt.Printf("luid:%v\r\n", luid)

		guid, err := aa.GetNetworkGuid()
		if err != nil {
			t.Error(err)
		}
		guidStr, err := StringFromGUID2(&guid)
		if err != nil {
			t.Error(err)
		}
		_ = guidStr
		// 奇怪输出的全部都是一个值:{A5735777-2F40-11E8-A039-806E6F6E6963}
		//fmt.Printf("guid:%v\r\n", guidStr)

		gas, err := aa.GetGatewayAddress()
		if err != nil {
			t.Error(err)
		}
		_ = gas
		// 输出的全部是空，不清楚原因。
		//fmt.Printf("IpAdapterGatewayAddress:%#v\r\n", gas)

		//	fmt.Printf("connectionType:%v\r\n", aa.connectionType)

	}
}

func TestStruct(t *testing.T) {
	if ptrSize == 8 {
		if unsafe.Sizeof(IpAdapterAddresses{}) != 448 {
			t.Errorf("IpAdapterAddresses %v!=448", unsafe.Sizeof(IpAdapterAddresses{}))
		}

		ipAdapterAddresses := IpAdapterAddresses{}

		if o := unsafe.Offsetof(ipAdapterAddresses.Length); o != 0 {
			t.Errorf("%v !=0", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.ipv6IfIndex); o != 108 {
			t.Errorf("%v !=108", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.luid); o != 224 {
			t.Errorf("%v !=224", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.dhcpv4Server); o != 232 {
			t.Errorf("%v !=232", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.compartmentId); o != 248 {
			t.Errorf("%v !=248", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.networkGuid); o != 252 {
			t.Errorf("%v !=252", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.connectionType); o != 268 {
			t.Errorf("%v !=268", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.dhcpv6Server); o != 280 {
			t.Errorf("%v !=280", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.firstDnsSuffix); o != 440 {
			t.Errorf("%v !=440", o)
		}

		if unsafe.Sizeof(IfLuid(0)) != 8 {
			t.Errorf("IfLuid %v!=8", unsafe.Sizeof(IfLuid(0)))
		}
		if unsafe.Sizeof(IpAdapterWinsServerAddress{}) != 32 {
			t.Errorf("IpAdapterWinsServerAddress %v!=32", unsafe.Sizeof(IpAdapterWinsServerAddress{}))
		}
		if unsafe.Sizeof(IpAdapterGatewayAddress{}) != 32 {
			t.Errorf("IpAdapterGatewayAddress %v!=32", unsafe.Sizeof(IpAdapterGatewayAddress{}))
		}
		if unsafe.Sizeof(windows.IpAdapterUnicastAddress{}) != 64 {
			t.Errorf("windows.IpAdapterUnicastAddress %v!=64", unsafe.Sizeof(windows.IpAdapterUnicastAddress{}))
		}
		if unsafe.Sizeof(windows.IpAdapterAnycastAddress{}) != 32 {
			t.Errorf("windows.IpAdapterAnycastAddress %v!=32", unsafe.Sizeof(windows.IpAdapterAnycastAddress{}))
		}
		if unsafe.Sizeof(windows.IpAdapterMulticastAddress{}) != 32 {
			t.Errorf("windows.IpAdapterMulticastAddress %v!=32", unsafe.Sizeof(windows.IpAdapterMulticastAddress{}))
		}
		if unsafe.Sizeof(windows.IpAdapterDnsServerAdapter{}) != 32 {
			t.Errorf("windows.IpAdapterDnsServerAdapter %v!=32", unsafe.Sizeof(windows.IpAdapterDnsServerAdapter{}))
		}
		if unsafe.Sizeof(windows.IpAdapterPrefix{}) != 40 {
			t.Errorf("windows.IpAdapterPrefix %v!=40", unsafe.Sizeof(windows.IpAdapterPrefix{}))
		}
		if unsafe.Sizeof(windows.SocketAddress{}) != 16 {
			t.Errorf("windows.SocketAddress %v!=16", unsafe.Sizeof(windows.SocketAddress{}))
		}
		if unsafe.Sizeof(CompartmentId(0)) != 4 {
			t.Errorf("CompartmentId %v!=4", unsafe.Sizeof(CompartmentId(0)))
		}
		if unsafe.Sizeof(NetworkGuid{}) != 16 {
			t.Errorf("NetworkGuid %v!=16", unsafe.Sizeof(NetworkGuid{}))
		}
		if unsafe.Sizeof(ConnectionType(0)) != 4 {
			t.Errorf("ConnectionType %v!=4", unsafe.Sizeof(ConnectionType(0)))
		}
		if unsafe.Sizeof(TunnelType(0)) != 4 {
			t.Errorf("TunnelType %v!=8", unsafe.Sizeof(TunnelType(0)))
		}
		if unsafe.Sizeof(IpAdapterDnsSuffix{}) != 520 {
			t.Errorf("IpAdapterDnsSuffix %v!=516", unsafe.Sizeof(IpAdapterDnsSuffix{}))
		}
	} else {
		if unsafe.Sizeof(IpAdapterAddresses{}) != 376 {
			t.Errorf("IpAdapterAddresses %v!=376", unsafe.Sizeof(IpAdapterAddresses{}))
		}

		ipAdapterAddresses := IpAdapterAddresses{}

		if o := unsafe.Offsetof(ipAdapterAddresses.Length); o != 0 {
			t.Errorf("%v !=0", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.ipv6IfIndex); o != 72 {
			t.Errorf("%v !=72", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.luid); o != 176 {
			t.Errorf("%v !=176", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.dhcpv4Server); o != 184 {
			t.Errorf("%v !=184", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.compartmentId); o != 192 {
			t.Errorf("%v !=192", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.networkGuid); o != 196 {
			t.Errorf("%v !=196", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.connectionType); o != 212 {
			t.Errorf("%v !=212", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.dhcpv6Server); o != 220 {
			t.Errorf("%v !=220", o)
		}
		if o := unsafe.Offsetof(ipAdapterAddresses.firstDnsSuffix); o != 368 {
			t.Errorf("%v !=368", o)
		}

		if unsafe.Sizeof(IfLuid(0)) != 8 {
			t.Errorf("IfLuid %v!=8", unsafe.Sizeof(IfLuid(0)))
		}

		if unsafe.Sizeof(windows.SocketAddress{}) != 8 {
			t.Errorf("windows.SocketAddress %v!=8", unsafe.Sizeof(windows.SocketAddress{}))
		}

		if s := unsafe.Sizeof(IpAdapterWinsServerAddress{}); s != 24 {
			t.Errorf("IpAdapterWinsServerAddress %v!=24", s)
		}

		ipAdapterWinsServerAddress := IpAdapterWinsServerAddress{}
		if o := unsafe.Offsetof(ipAdapterWinsServerAddress.Length); o != 0 {
			t.Errorf("%v !=0", o)
		}
		if o := unsafe.Offsetof(ipAdapterWinsServerAddress.Reserved); o != 4 {
			t.Errorf("%v !=4", o)
		}
		if o := unsafe.Offsetof(ipAdapterWinsServerAddress.Next); o != 8 {
			t.Errorf("%v !=8", o)
		}
		if o := unsafe.Offsetof(ipAdapterWinsServerAddress.Address); o != 12 {
			t.Errorf("%v !=12", o)
		}

		if unsafe.Sizeof(IpAdapterGatewayAddress{}) != 24 {
			t.Errorf("IpAdapterGatewayAddress %v!=24", unsafe.Sizeof(IpAdapterGatewayAddress{}))
		}
		if unsafe.Sizeof(windows.IpAdapterUnicastAddress{}) != 48 {
			t.Errorf("windows.IpAdapterUnicastAddress %v!=48", unsafe.Sizeof(windows.IpAdapterUnicastAddress{}))
		}
		if unsafe.Sizeof(windows.IpAdapterAnycastAddress{}) != 24 {
			t.Errorf("windows.IpAdapterAnycastAddress %v!=24", unsafe.Sizeof(windows.IpAdapterAnycastAddress{}))
		}
		if unsafe.Sizeof(windows.IpAdapterMulticastAddress{}) != 24 {
			t.Errorf("windows.IpAdapterMulticastAddress %v!=24", unsafe.Sizeof(windows.IpAdapterMulticastAddress{}))
		}
		if unsafe.Sizeof(windows.IpAdapterDnsServerAdapter{}) != 24 {
			t.Errorf("windows.IpAdapterDnsServerAdapter %v!=24", unsafe.Sizeof(windows.IpAdapterDnsServerAdapter{}))
		}
		if unsafe.Sizeof(windows.IpAdapterPrefix{}) != 24 {
			t.Errorf("windows.IpAdapterPrefix %v!=24", unsafe.Sizeof(windows.IpAdapterPrefix{}))
		}
		if unsafe.Sizeof(CompartmentId(0)) != 4 {
			t.Errorf("CompartmentId %v!=4", unsafe.Sizeof(CompartmentId(0)))
		}
		if unsafe.Sizeof(NetworkGuid{}) != 16 {
			t.Errorf("NetworkGuid %v!=16", unsafe.Sizeof(NetworkGuid{}))
		}
		if unsafe.Sizeof(ConnectionType(0)) != 4 {
			t.Errorf("ConnectionType %v!=4", unsafe.Sizeof(ConnectionType(0)))
		}
		if unsafe.Sizeof(TunnelType(0)) != 4 {
			t.Errorf("TunnelType %v!=8", unsafe.Sizeof(TunnelType(0)))
		}
		if unsafe.Sizeof(IpAdapterDnsSuffix{}) != 516 {
			t.Errorf("IpAdapterDnsSuffix %v!=516", unsafe.Sizeof(IpAdapterDnsSuffix{}))
		}
	}
}

func TestGetIpForwardTable(t *testing.T) {
	rows, err := GetIpForwardTable()
	if err != nil {
		t.Fatal(err)
	}
	_ = rows

	for _, row := range rows {
		t.Log(row.String())
	}
	//	t.Logf("%#v", rows)
}


/*
// 手工wifi切换网络测试通过
func TestNotifyAddrChangeSync(t *testing.T) {
	t.Log("TestNotifyAddrChangeSync...")
	err:=NotifyAddrChange(nil,nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestNotifyAddrChangeSync ok")
}*/

/*
// 手工wifi切换网络测试通过
func TestNotifyRouteChangeSync(t *testing.T) {
	t.Log("TestNotifyRouteChangeSync ...")
	err:=NotifyRouteChange(nil,nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestNotifyRouteChangeSync ok")
}
*/


/*
// 手工wifi切换网络测试通过
func TestNotifyAddrChangeASync(t *testing.T) {
	t.Log("TestNotifyAddrChangeASync...")

	overlap := Overlapped{}
	hEvent, err := WSACreateEvent()
	if err != nil {
		t.Fatal(err)
	}
	overlap.HEvent = windows.Handle(hEvent)

	hand := Handle(0)

	err = NotifyAddrChange(&hand, &overlap)
	if err != nil {
		t.Fatal(err)
	}
	event, err := WaitForSingleObject(overlap.HEvent, INFINITE)
	if err != nil {
		t.Fatal(err)
	}

	if event != WAIT_OBJECT_0 {
		t.Fatal(event, " != WAIT_OBJECT_0")
	}

	t.Log("TestNotifyAddrChangeASync ok")
}
*/

/*
// 手工wifi切换网络测试通过
func TestNotifyRouteChangeASync(t *testing.T) {
	t.Log("TestNotifyRouteChangeASync...")

	overlap := Overlapped{}
	hEvent, err := WSACreateEvent()
	if err != nil {
		t.Fatal(err)
	}
	overlap.HEvent = windows.Handle(hEvent)

	hand := Handle(0)

	err = NotifyRouteChange(&hand, &overlap)
	if err != nil {
		t.Fatal(err)
	}
	event, err := WaitForSingleObject(overlap.HEvent, INFINITE)
	if err != nil {
		t.Fatal(err)
	}

	if event != WAIT_OBJECT_0 {
		t.Fatal(event, " != WAIT_OBJECT_0")
	}

	t.Log("TestNotifyRouteChangeASync ok")
}
*/