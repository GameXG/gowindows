package gowindows

import (
	"net"
	"testing"
	"time"
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
		t.Logf("\r\nFriendlyName:%v\r\n", aa.GetFriendlyName())
		t.Logf("IfType:%v\r\n", aa.IfType)
		t.Logf("Description:%v\r\n", aa.GetDescription())

		luid, err := aa.GetLuid()
		if err != nil {
			t.Error(err)
		}
		_ = luid
		t.Logf("luid:%v\r\n", luid)

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
		t.Logf("guid:%v\r\n", guidStr)

		uas, err := aa.GetUnicastIpAddress()
		if err != nil {
			t.Errorf("GetUnicastIpAddress,%v", err)
		}
		for _, v := range uas {
			t.Logf("UnicastAddress:%v(%#v)", v.IP, v.IP)
		}

		gas, err := aa.GetGatewayIpAddress()
		if err != nil {
			t.Error(err)
		}
		for _, v := range gas {
			t.Logf("Gateway:%v", v)
		}

		t.Logf("connectionType:%v\r\n", aa.connectionType)

		dns, err := aa.GetDnsServerIpAddress()
		if err != nil {
			t.Error(err)
		}

		for _, v := range dns {
			t.Logf("DNS:%v(%#v)", v, v)
		}

	}
}

func TestStruct(t *testing.T) {
	if ptrSize == 8 {
		t.Log("64")

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

		if unsafe.Sizeof(MibIpAddrTable{}) != 28 {
			t.Errorf("MibIpAddrTable %v!=28", unsafe.Sizeof(MibIpAddrTable{}))
		}
		mibIpAddrTable := MibIpAddrTable{}
		if o := unsafe.Offsetof(mibIpAddrTable.Table); o != 4 {
			t.Errorf("%v !=4", o)
		}

		if unsafe.Sizeof(MibIpAddrRowW2k{}) != 24 {
			t.Errorf("MibIpAddrRowW2k %v!=24", unsafe.Sizeof(MibIpAddrRowW2k{}))
		}
		mibIpAddrRowW2k := MibIpAddrRowW2k{}
		if o := unsafe.Offsetof(mibIpAddrRowW2k.Mask); o != 8 {
			t.Errorf("%v !=8", o)
		}
		if o := unsafe.Offsetof(mibIpAddrRowW2k.Unused2); o != 22 {
			t.Errorf("%v !=22", o)
		}

	} else {
		t.Log("32")

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

		if unsafe.Sizeof(MibIpAddrTable{}) != 28 {
			t.Errorf("MibIpAddrTable %v!=28", unsafe.Sizeof(MibIpAddrTable{}))
		}
		mibIpAddrTable := MibIpAddrTable{}
		if o := unsafe.Offsetof(mibIpAddrTable.Table); o != 4 {
			t.Errorf("%v !=4", o)
		}

		if unsafe.Sizeof(MibIpAddrRowW2k{}) != 24 {
			t.Errorf("MibIpAddrRowW2k %v!=24", unsafe.Sizeof(MibIpAddrRowW2k{}))
		}
		mibIpAddrRowW2k := MibIpAddrRowW2k{}
		if o := unsafe.Offsetof(mibIpAddrRowW2k.Mask); o != 8 {
			t.Errorf("%v !=8", o)
		}
		if o := unsafe.Offsetof(mibIpAddrRowW2k.Unused2); o != 22 {
			t.Errorf("%v !=22", o)
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

// 手工wifi切换网络测试通过
func TestCancelIPChangeNotify(t *testing.T) {
	t.Log("TestCancelIPChangeNotify...")

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

	go func() {
		time.Sleep(1 * time.Second)
		r1, err := CancelIPChangeNotify(&overlap)
		if err != nil {
			t.Fatal(err)
		}

		if r1 != true {
			t.Errorf("r1!=true")
		}
	}()

	event, err := WaitForSingleObject(overlap.HEvent, INFINITE)
	if err != nil {
		t.Fatal(err)
	}

	// 取消时值也是 WAIT_OBJECT_0
	if event != WAIT_OBJECT_0 {
		t.Fatal(event, " != WAIT_OBJECT_0")
	}

	t.Log("TestNotifyRouteChangeASync ok")
}

/*
func TestIPChangeNotify_Reset(t *testing.T) {

	f:=func(name string){
	n:=IPChangeNotify{}

	err:=n.Reset(true,true)
	if err != nil {
		t.Fatal(err)
	}


	// 10 秒后取消
	go func(){
		time.Sleep(20*time.Second)
		t.Log(time.Now(), " [",name,"] close")
		err:=n.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 接收更改
	for {
		select {
		case v:=<-n.C:
			t.Logf("%v [%v] %#v\r\n",time.Now(),name,v)
			case <-n.Done():
				return
		}
	}
	}

	go f("1")
	go f("2")

	time.Sleep(20*time.Second)
}*/

/*
wifi 断开
    iphlpapi_windows_test.go:421: 2018-12-31 21:29:27.3559343 +0800 CST m=+3.445919201 [2] &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:421: 2018-12-31 21:29:27.3739352 +0800 CST m=+3.463920101 [1] &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:410: 2018-12-31 21:29:44.9400476 +0800 CST m=+21.030032501  [ 2 ] close
    iphlpapi_windows_test.go:410: 2018-12-31 21:29:44.9400476 +0800 CST m=+21.030032501  [ 1 ] close
*/

/*
单个断开 wifi + 重新连接 wifi 时的消息


    iphlpapi_windows_test.go:419: 2018-12-25 16:02:23.7224344 +0800 CST m=+4.436488501 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:23.7734356 +0800 CST m=+4.487489701 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:23.8034356 +0800 CST m=+4.517489701 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:23.8194343 +0800 CST m=+4.533488401 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:23.9044358 +0800 CST m=+4.618489901 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:23.9044358 +0800 CST m=+4.618489901 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:26.6590696 +0800 CST m=+7.373123701 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:27.6715333 +0800 CST m=+8.385587401 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:27.6795325 +0800 CST m=+8.393586601 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:27.7855319 +0800 CST m=+8.499586001 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:27.7875314 +0800 CST m=+8.501585501 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:27.8055308 +0800 CST m=+8.519584901 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:32.2337128 +0800 CST m=+12.947766901 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:32.2757113 +0800 CST m=+12.989765401 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:32.2857124 +0800 CST m=+12.999766501 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:32.3257123 +0800 CST m=+13.039766401 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:32.372712 +0800 CST m=+13.086766101 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:32.372712 +0800 CST m=+13.086766101 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:32.4317122 +0800 CST m=+13.145766301 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:32.4317122 +0800 CST m=+13.145766301 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.7755426 +0800 CST m=+16.489596701 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.7845405 +0800 CST m=+16.498594601 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.809544 +0800 CST m=+16.523598101 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.809544 +0800 CST m=+16.523598101 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.809544 +0800 CST m=+16.523598101 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.8305413 +0800 CST m=+16.544595401 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.8305413 +0800 CST m=+16.544595401 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.8735428 +0800 CST m=+16.587596901 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.9055429 +0800 CST m=+16.619597001 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.930545 +0800 CST m=+16.644599101 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.930545 +0800 CST m=+16.644599101 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:35.930545 +0800 CST m=+16.644599101 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:36.0365442 +0800 CST m=+16.750598301 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:36.0465464 +0800 CST m=+16.760600501 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:36.0785433 +0800 CST m=+16.792597401 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:39.1555427 +0800 CST m=+19.869596801 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:39.1615416 +0800 CST m=+19.875595701 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:39.1705432 +0800 CST m=+19.884597301 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:39.2375421 +0800 CST m=+19.951596201 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:39.3815429 +0800 CST m=+20.095597001 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:true, IsRoute:false}
    iphlpapi_windows_test.go:419: 2018-12-25 16:02:39.3815429 +0800 CST m=+20.095597001 &gowindows.IPChangeNotifyChanData{Err:error(nil), IsAddr:false, IsRoute:true}
    iphlpapi_windows_test.go:408: 2018-12-25 16:02:40.3450281 +0800 CST m=+21.059082201  close

*/

func TestGetIpAddrTable(t *testing.T) {
	rows, err := GetIpAddrTable(false)
	if err != nil {
		panic(err)
	}

	if len(rows) == 0 {
		t.Error("len(rows)==0")
	}

	for _, v := range rows {
		t.Logf("index:%v\r\n", v.Index)
		t.Logf("ip:%v/%v\r\n", v.GetAddr(), v.GetMask())
		t.Logf("BCastAddr:%v\r\n", v.GetBCastAddr())
		t.Logf("ReasmSize:%v\r\n", v.ReasmSize)
	}
}

func TestUint322Ip(t *testing.T) {
	ip := net.IPv4(0x11, 0x22, 0x33, 0x44)

	_int, err := ip2uint32(ip)
	if err != nil {
		t.Fatal(err)
	}

	if _int != 0x44332211 {
		t.Errorf("%v!=0x44332211", _int)
	}

	ip2 := uint322Ip(_int)

	if ip2.Equal(ip) == false {
		t.Errorf("%v!=%v", ip2, ip)
	}

}
