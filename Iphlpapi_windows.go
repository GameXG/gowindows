package gowindows

import (
	"context"
	"net"
	"strconv"
	"sync"
	"syscall"
	"unsafe"

	"os"

	"fmt"

	"golang.org/x/sys/windows"
)

var (
	iphlpapi             = syscall.NewLazyDLL("iphlpapi.dll")
	getIpForwardTable    = iphlpapi.NewProc("GetIpForwardTable")
	createIpForwardEntry = iphlpapi.NewProc("CreateIpForwardEntry")
	deleteIpForwardEntry = iphlpapi.NewProc("DeleteIpForwardEntry")
	notifyAddrChange     = iphlpapi.NewProc("NotifyAddrChange")
	notifyRouteChange    = iphlpapi.NewProc("NotifyRouteChange")
	cancelIPChangeNotify = iphlpapi.NewProc("CancelIPChangeNotify")
	getIpAddrTable       = iphlpapi.NewProc("GetIpAddrTable")
)

// https://docs.microsoft.com/zh-cn/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_addresses_lh
type IfOperStatus uint32

const (
	//接口已启动并且能够传递数据包。
	IfOperStatusUp IfOperStatus = 1

	//接口已关闭，并且不处于传递数据包的状态。该 IfOperStatusDown状态有两个含义，这取决于的值 AdminStatus构件。如果AdminStatus未设置为 NET_IF_ADMIN_STATUS_DOWN且ifOperStatus设置为 IfOperStatusDown，则假定接口上存在故障条件。如果 AdminStatus设置为IfOperStatusDown，那么 ifOperStatus通常也会设置为 IfOperStatusDown或IfOperStatusNotPresent 并且接口上不一定存在故障情况。
	IfOperStatusDown IfOperStatus = 2

	//接口处于测试模式。
	IfOperStatusTesting IfOperStatus = 3

	//接口的运行状态未知。
	IfOperStatusUnknown IfOperStatus = 4

	//接口实际上并不处于传递数据包的状态（它没有启动），而是处于挂起状态，等待一些外部事件。对于按需接口，此新状态标识接口正在等待事件将其置于IfOperStatusUp状态的情况。
	IfOperStatusDormant IfOperStatus = 5

	//IfOperStatusDown状态的 细化，指示相关接口是特定关闭的，因为托管系统中不存在某些组件（通常是硬件组件）。
	IfOperStatusNotPresent IfOperStatus = 6

	//IfOperStatusDown状态的 细化。此新状态表示此接口在一个或多个其他接口之上运行，并且此接口已关闭，因为这些较低层接口中的一个或多个已关闭。
	IfOperStatusLowerLayerDown IfOperStatus = 7
)

// IP_ADAPTER_ADDRESSES_LH
// https://docs.microsoft.com/zh-cn/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_addresses_lh
/*
typedef struct _IP_ADAPTER_ADDRESSES_LH {
  union {
    ULONGLONG Alignment;
    struct {
      ULONG    Length;
      IF_INDEX IfIndex;
    };
  };
  struct _IP_ADAPTER_ADDRESSES_LH    *Next;
  PCHAR                              AdapterName;
  PIP_ADAPTER_UNICAST_ADDRESS_LH     FirstUnicastAddress;
  PIP_ADAPTER_ANYCAST_ADDRESS_XP     FirstAnycastAddress;
  PIP_ADAPTER_MULTICAST_ADDRESS_XP   FirstMulticastAddress;
  PIP_ADAPTER_DNS_SERVER_ADDRESS_XP  FirstDnsServerAddress;
  PWCHAR                             DnsSuffix;
  PWCHAR                             Description;
  PWCHAR                             FriendlyName;
  BYTE                               PhysicalAddress[MAX_ADAPTER_ADDRESS_LENGTH];
  ULONG                              PhysicalAddressLength;
  union {
    ULONG Flags;
    struct {
      ULONG DdnsEnabled : 1;
      ULONG RegisterAdapterSuffix : 1;
      ULONG Dhcpv4Enabled : 1;
      ULONG ReceiveOnly : 1;
      ULONG NoMulticast : 1;
      ULONG Ipv6OtherStatefulConfig : 1;
      ULONG NetbiosOverTcpipEnabled : 1;
      ULONG Ipv4Enabled : 1;
      ULONG Ipv6Enabled : 1;
      ULONG Ipv6ManagedAddressConfigurationSupported : 1;
    };
  };
  ULONG                              Mtu;
  IFTYPE                             IfType;
  IF_OPER_STATUS                     OperStatus;
  IF_INDEX                           Ipv6IfIndex;
  ULONG                              ZoneIndices[16];
  PIP_ADAPTER_PREFIX_XP              FirstPrefix;
  ULONG64                            TransmitLinkSpeed;
  ULONG64                            ReceiveLinkSpeed;
  PIP_ADAPTER_WINS_SERVER_ADDRESS_LH FirstWinsServerAddress;
  PIP_ADAPTER_GATEWAY_ADDRESS_LH     FirstGatewayAddress;
  ULONG                              Ipv4Metric;
  ULONG                              Ipv6Metric;
  IF_LUID                            Luid;
  SOCKET_ADDRESS                     Dhcpv4Server;
  NET_IF_COMPARTMENT_ID              CompartmentId;
  NET_IF_NETWORK_GUID                NetworkGuid;
  NET_IF_CONNECTION_TYPE             ConnectionType;
  TUNNEL_TYPE                        TunnelType;
  SOCKET_ADDRESS                     Dhcpv6Server;
  BYTE                               Dhcpv6ClientDuid[MAX_DHCPV6_DUID_LENGTH];
  ULONG                              Dhcpv6ClientDuidLength;
  ULONG                              Dhcpv6Iaid;
  PIP_ADAPTER_DNS_SUFFIX             FirstDnsSuffix;
} IP_ADAPTER_ADDRESSES_LH, *PIP_ADAPTER_ADDRESSES_LH;
*/
//TODO: 结构应该存在问题，GUID 之前应该有字段不对！
type IpAdapterAddresses struct {
	Length                uint32
	IfIndex               uint32
	Next                  *IpAdapterAddresses
	AdapterName           *byte
	FirstUnicastAddress   *windows.IpAdapterUnicastAddress
	FirstAnycastAddress   *windows.IpAdapterAnycastAddress
	FirstMulticastAddress *windows.IpAdapterMulticastAddress
	FirstDnsServerAddress *windows.IpAdapterDnsServerAdapter
	DnsSuffix             *uint16
	Description           *uint16
	FriendlyName          *uint16
	PhysicalAddress       [syscall.MAX_ADAPTER_ADDRESS_LENGTH]byte
	PhysicalAddressLength uint32
	Flags                 uint32
	Mtu                   uint32
	IfType                uint32
	OperStatus            IfOperStatus

	// 以下是 windows xp sp1 之后添加的
	ipv6IfIndex uint32
	zoneIndices [16]uint32
	firstPrefix *windows.IpAdapterPrefix

	// 以下是 windows Vista 之后添加的
	transmitLinkSpeed      uint64
	receiveLinkSpeed       uint64
	firstWinsServerAddress *IpAdapterWinsServerAddress
	firstGatewayAddress    *IpAdapterGatewayAddress
	ipv4Metric             uint32
	ipv6Metric             uint32
	luid                   IfLuid
	dhcpv4Server           windows.SocketAddress
	compartmentId          CompartmentId
	networkGuid            NetworkGuid
	connectionType         ConnectionType
	tunnelType             TunnelType
	dhcpv6Server           windows.SocketAddress
	dhcpv6ClientDuid       [MAX_DHCPV6_DUID_LENGTH]byte
	dhcpv6ClientDuidLength uint32
	dhcpv6Iaid             uint32

	// 以下是 windows Vista SP1 及 windows server 2008 之后添加的
	firstDnsSuffix *IpAdapterDnsSuffix
}

//typedef struct _IP_ADAPTER_WINS_SERVER_ADDRESS_LH {
//    union {
//        ULONGLONG Alignment;
//        struct {
//            ULONG Length;
//            DWORD Reserved;
//        };
//    };
//    struct _IP_ADAPTER_WINS_SERVER_ADDRESS_LH *Next;
//    SOCKET_ADDRESS Address;
//} IP_ADAPTER_WINS_SERVER_ADDRESS_LH, *PIP_ADAPTER_WINS_SERVER_ADDRESS_LH;
type IpAdapterWinsServerAddress struct {
	Length   uint32
	Reserved int32
	Next     *IpAdapterWinsServerAddress
	Address  windows.SocketAddress
}

type IpAdapterGatewayAddress struct {
	Length   uint32
	Reserved int32
	Next     *IpAdapterGatewayAddress
	Address  windows.SocketAddress
}

//typedef struct _MIB_IPADDRTABLE {
//    DWORD dwNumEntries;
//    MIB_IPADDRROW table[ANY_SIZE];
//} MIB_IPADDRTABLE, *PMIB_IPADDRTABLE;
// https://docs.microsoft.com/zh-cn/windows/win32/api/ipmib/ns-ipmib-_mib_ipaddrtable
type MibIpAddrTable struct {
	NumEntries DWord
	Table      [ANY_SIZE]MibIpAddrRowW2k
}

// typedef struct _MIB_IPADDRROW_W2K {
//    DWORD dwAddr;
//    DWORD dwIndex;
//    DWORD dwMask;
//    DWORD dwBCastAddr;
//    DWORD dwReasmSize;
//    unsigned short Unused1;
//    unsigned short Unused2;
//} MIB_IPADDRROW_W2K, *PMIB_IPADDRROW_W2K;
//https://docs.microsoft.com/zh-cn/windows/win32/api/ipmib/ns-ipmib-mib_ipaddrrow_w2k
type MibIpAddrRowW2k struct {
	Addr      DWord
	Index     DWord
	Mask      DWord
	BCastAddr DWord
	ReasmSize DWord
	Unused1   uint16
	Unused2   uint16
}

func (r *MibIpAddrRowW2k) GetAddr() net.IP {
	return uint322Ip(r.Addr)
}

func (r *MibIpAddrRowW2k) GetMask() net.IPMask {
	return net.IPMask(uint322Ip(r.Mask))
}

func (r *MibIpAddrRowW2k) GetBCastAddr() net.IP {
	return uint322Ip(r.BCastAddr)
}

func uint322Ip(ip uint32) net.IP {
	return net.IPv4(byte(ip), byte(ip>>8), byte(ip>>16), byte(ip>>24))
}

func ip2uint32(ip net.IP) (uint32, error) {
	_ip := ip.To4()

	if len(_ip) == 4 {
		return uint32(_ip[0]) | uint32(_ip[1])<<8 | uint32(_ip[2])<<16 | uint32(_ip[3])<<24, nil
	}

	return 0, fmt.Errorf("%v 不是 ipv6 格式。", ip)
}

// 一个字符数组，包含与地址关联的适配器的名称。与适配器的友好名称不同，AdapterName中指定的适配器名称是永久性的，用户无法修改。
func (aa *IpAdapterAddresses) GetAdapterName() string {
	// C:/Go/src/net/interface_windows.go:77
	return string((*(*[10000]byte)(unsafe.Pointer(aa.AdapterName)))[:])
}
func (aa *IpAdapterAddresses) GetLuid() (IfLuid, error) {
	tz := aa.Length
	fz := unsafe.Offsetof(aa.luid) + unsafe.Sizeof(aa.luid)

	// 判断结构是否包含了指定的字段
	// 不同版本的 windows 包含的字段不同，老版本的不包含新版本的字段。
	if tz < uint32(fz) {
		return IfLuid(0), fmt.Errorf("Length(%v)<%v", tz, fz)
	}

	return aa.luid, nil
}

// 适配器的接收链路的当前速度（以每秒位数为单位）。
// 注意   此结构成员仅适用于Windows Vista及更高版本。
func (aa *IpAdapterAddresses) GetReceiveLinkSpeed() (uint64, error) {
	tz := aa.Length
	fz := unsafe.Offsetof(aa.receiveLinkSpeed) + unsafe.Sizeof(aa.receiveLinkSpeed)

	// 判断结构是否包含了指定的字段
	// 不同版本的 windows 包含的字段不同，老版本的不包含新版本的字段。
	if tz < uint32(fz) {
		return 0, fmt.Errorf("Length(%v)<%v", tz, fz)
	}

	return aa.receiveLinkSpeed, nil
}
func (aa *IpAdapterAddresses) GetNetworkGuid() (NetworkGuid, error) {
	tz := aa.Length
	fz := unsafe.Offsetof(aa.networkGuid) + unsafe.Sizeof(aa.networkGuid)

	// 判断结构是否包含了指定的字段
	// 不同版本的 windows 包含的字段不同，老版本的不包含新版本的字段。
	if tz < uint32(fz) {
		return NetworkGuid{}, fmt.Errorf("Length(%v)<%v", tz, fz)
	}

	return aa.networkGuid, nil
}

func (aa *IpAdapterAddresses) GetFriendlyName() string {
	// C:/Go/src/net/interface_windows.go:77
	return syscall.UTF16ToString((*(*[10000]uint16)(unsafe.Pointer(aa.FriendlyName)))[:])
}
func (aa *IpAdapterAddresses) GetDescription() string {
	// C:/Go/src/net/interface_windows.go:77
	return syscall.UTF16ToString((*(*[10000]uint16)(unsafe.Pointer(aa.Description)))[:])
}

func (aa *IpAdapterAddresses) GetGatewayAddress() ([]*IpAdapterGatewayAddress, error) {
	tz := aa.Length
	fz := unsafe.Offsetof(aa.firstGatewayAddress) + unsafe.Sizeof(aa.firstGatewayAddress)

	// 判断结构是否包含了指定的字段
	// 不同版本的 windows 包含的字段不同，老版本的不包含新版本的字段。
	if tz < uint32(fz) {
		return nil, fmt.Errorf("Length(%v)<%v", tz, fz)
	}

	res := make([]*IpAdapterGatewayAddress, 0, 1)
	ga := aa.firstGatewayAddress

	for ga != nil {
		res = append(res, ga)
		ga = ga.Next
	}

	return res, nil
}

func (aa *IpAdapterAddresses) GetHardwareAddr() (net.HardwareAddr, error) {
	tz := aa.Length
	fz := unsafe.Offsetof(aa.PhysicalAddressLength) + unsafe.Sizeof(aa.PhysicalAddressLength)

	// 判断结构是否包含了指定的字段
	// 不同版本的 windows 包含的字段不同，老版本的不包含新版本的字段。
	if tz < uint32(fz) {
		return nil, fmt.Errorf("Length(%v)<%v", tz, fz)
	}

	if aa.PhysicalAddressLength > 0 {
		hardwareAddr := make([]byte, aa.PhysicalAddressLength)
		copy(hardwareAddr, aa.PhysicalAddress[:])
		return hardwareAddr, nil
	}

	return nil, fmt.Errorf("PhysicalAddressLength == 0")
}

func (aa *IpAdapterAddresses) GetGatewayIpAddress() ([]net.IPAddr, error) {
	ads, err := aa.GetGatewayAddress()
	if err != nil {
		return nil, err
	}

	res := make([]net.IPAddr, 0, len(ads))
	for _, v := range ads {
		ipAddr, err := Sockaddr2IpAddr(v.Address.Sockaddr)
		if err != nil {
			return nil, err
		}
		res = append(res, ipAddr)
	}
	return res, nil
}

func Sockaddr2IpAddr(rd *syscall.RawSockaddrAny) (net.IPAddr, error) {
	sa, err := rd.Sockaddr()
	if err != nil {
		return net.IPAddr{}, err
	}

	switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
		return net.IPAddr{IP: net.IPv4(sa.Addr[0], sa.Addr[1], sa.Addr[2], sa.Addr[3])}, nil
	case *syscall.SockaddrInet6:
		return net.IPAddr{IP: make(net.IP, net.IPv6len)}, nil
	default:
		return net.IPAddr{}, fmt.Errorf("不支持的地址类型，%v", sa)
	}
}

// 注意 windows xp  IpAdapterUnicastAddress 不包含 OnLinkPrefixLength 字段，即无法获取ip掩码。
// 参考：https://docs.microsoft.com/en-us/windows/win32/api/iptypes/ns-iptypes-_ip_adapter_unicast_address_lh
func UnicastIpAddress2IpNet(ua *windows.IpAdapterUnicastAddress) (net.IPNet, error) {
	rd := ua.Address.Sockaddr
	sa, err := rd.Sockaddr()
	if err != nil {
		return net.IPNet{}, err
	}

	// windows xp 不存在 onLinkPrefixLength 字段
	// 不过实测这里判断无效，因为 winxp 32/64 结构长度都是 48 ，比 ua.OnLinkPrefixLength 长度 45 大。
	// 使用未定义行为并不是一个好主意，所以还是保留了这个部分判断。
	// https://docs.microsoft.com/en-us/windows/win32/api/iptypes/ns-iptypes-_ip_adapter_unicast_address_lh
	onLinkPrefixLength := 0
	tz := ua.Length
	fz := unsafe.Offsetof(ua.OnLinkPrefixLength) + unsafe.Sizeof(ua.OnLinkPrefixLength)
	if tz >= uint32(fz) {
		onLinkPrefixLength = int(ua.OnLinkPrefixLength)
	}

	switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
		return net.IPNet{IP: net.IPv4(sa.Addr[0], sa.Addr[1], sa.Addr[2], sa.Addr[3]), Mask: net.CIDRMask(onLinkPrefixLength, 8*net.IPv4len)}, nil
	case *syscall.SockaddrInet6:
		ipNet := net.IPNet{IP: make(net.IP, net.IPv6len), Mask: net.CIDRMask(onLinkPrefixLength, 8*net.IPv4len)}
		copy(ipNet.IP, sa.Addr[:])
		return ipNet, nil
	default:
		return net.IPNet{}, fmt.Errorf("不支持的地址类型，%v", sa)
	}
}

func (aa *IpAdapterAddresses) GetDnsServerAddress() ([]*windows.IpAdapterDnsServerAdapter, error) {
	tz := aa.Length
	fz := unsafe.Offsetof(aa.FirstDnsServerAddress) + unsafe.Sizeof(aa.FirstDnsServerAddress)

	// 判断结构是否包含了指定的字段
	// 不同版本的 windows 包含的字段不同，老版本的不包含新版本的字段。
	if tz < uint32(fz) {
		return nil, fmt.Errorf("Length(%v)<%v", tz, fz)
	}

	res := make([]*windows.IpAdapterDnsServerAdapter, 0, 1)

	for v := aa.FirstDnsServerAddress; v != nil; v = v.Next {
		res = append(res, v)
	}

	return res, nil
}
func (aa *IpAdapterAddresses) GetDnsServerIpAddress() ([]net.IPAddr, error) {
	ads, err := aa.GetDnsServerAddress()
	if err != nil {
		return nil, err
	}

	res := make([]net.IPAddr, 0, len(ads))
	for _, v := range ads {
		ipAddr, err := Sockaddr2IpAddr(v.Address.Sockaddr)
		if err != nil {
			return nil, err
		}
		res = append(res, ipAddr)
	}
	return res, nil
}

// 获得单播地址列表
// 注意 windows xp  IpAdapterUnicastAddress 不包含 OnLinkPrefixLength 字段，即无法获取ip掩码。
// 参考：https://docs.microsoft.com/en-us/windows/win32/api/iptypes/ns-iptypes-_ip_adapter_unicast_address_lh
func (aa *IpAdapterAddresses) GetUnicastAddress() ([]*windows.IpAdapterUnicastAddress, error) {
	tz := aa.Length
	fz := unsafe.Offsetof(aa.FirstUnicastAddress) + unsafe.Sizeof(aa.FirstUnicastAddress)

	// 判断结构是否包含了指定的字段
	// 不同版本的 windows 包含的字段不同，老版本的不包含新版本的字段。
	if tz < uint32(fz) {
		return nil, fmt.Errorf("Length(%v)<%v", tz, fz)
	}

	res := make([]*windows.IpAdapterUnicastAddress, 0, 1)

	for v := aa.FirstUnicastAddress; v != nil; v = v.Next {
		res = append(res, v)
	}

	return res, nil
}

// 获得单播地址列表
// 注意 windows xp  IpAdapterUnicastAddress 不包含 OnLinkPrefixLength 字段，即无法获取ip掩码。
// 所以 windows xp 下掩码将永远为0
// 参考：https://docs.microsoft.com/en-us/windows/win32/api/iptypes/ns-iptypes-_ip_adapter_unicast_address_lh
func (aa *IpAdapterAddresses) GetUnicastIpAddress() ([]net.IPNet, error) {
	ads, err := aa.GetUnicastAddress()
	if err != nil {
		return nil, err
	}

	res := make([]net.IPNet, 0, len(ads))
	for _, v := range ads {
		ipNet, err := UnicastIpAddress2IpNet(v)
		if err != nil {
			return nil, err
		}
		res = append(res, ipNet)
	}
	return res, nil
}

// https://docs.microsoft.com/en-us/windows/desktop/api/iphlpapi/nf-iphlpapi-getadaptersaddresses
// 实测 win10 x64 下可以获得断开网卡的固定ip
func AdapterAddresses() ([]*IpAdapterAddresses, error) {
	var b []byte
	l := uint32(15000) // recommended initial size
	for {
		b = make([]byte, l)
		err := windows.GetAdaptersAddresses(syscall.AF_UNSPEC, GAA_FLAG_INCLUDE_PREFIX|GAA_FLAG_INCLUDE_WINS_INFO|GAA_FLAG_INCLUDE_GATEWAYS, 0, (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])), &l)
		if err == nil {
			if l == 0 {
				return nil, nil
			}
			break
		}
		if err.(syscall.Errno) != syscall.ERROR_BUFFER_OVERFLOW {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
		if l <= uint32(len(b)) {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
	}
	//todo 需要确认是否存在内存释放问题，虽然标准库也是这样做的。
	var aas []*IpAdapterAddresses
	for aa := (*IpAdapterAddresses)(unsafe.Pointer(&b[0])); aa != nil; aa = aa.Next {
		aas = append(aas, aa)
	}
	return aas, nil
}

func GetIpForwardTable() ([]MibIpForwardRow, error) {
	buf := []byte{0}
	bufSize := uint32(len(buf))
	var r1 uintptr
	var e1 error
	for i := 0; i < 10; i++ {
		buf = make([]byte, bufSize)
		r1, _, e1 = getIpForwardTable.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&bufSize)), 0)
		if r1 == ERROR_INSUFFICIENT_BUFFER {
			// 空间不足
			continue
		}

		break
	}

	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return nil, e1
		} else {
			return nil, fmt.Errorf("r1:%v", r1)
		}
	}

	table := (*MibIpForwardTable)(unsafe.Pointer(&buf[0]))
	rows := table.Table[:]
	err := ChangeSliceSize(&rows, int(table.NumEntries), int(table.NumEntries))
	if err != nil {
		return nil, fmt.Errorf("ChangeSliceSize, %v", err)
	}

	res := make([]MibIpForwardRow, len(rows))
	copy(res, rows)
	return res, nil
}

func CreateIpForwardEntry(row *MibIpForwardRow) error {
	r1, _, e1 := createIpForwardEntry.Call(uintptr(unsafe.Pointer(row)))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return fmt.Errorf("r1:%v", r1)
		}
	}

	return nil
}

// 必须提供以下成员：dwForwardIfIndex，dwForwardDest，dwForwardMask，dwForwardNextHop和dwForwardProto
func DeleteIpForwardEntry(row *MibIpForwardRow) error {
	r1, _, e1 := deleteIpForwardEntry.Call(uintptr(unsafe.Pointer(row)))
	if r1 != 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return fmt.Errorf("r1:%v", r1)
		}
	}

	return nil
}

//https://docs.microsoft.com/en-us/windows/desktop/api/iphlpapi/nf-iphlpapi-notifyaddrchange
//DWORD NotifyAddrChange(
//  PHANDLE      Handle,
//  LPOVERLAPPED overlapped
//);
func NotifyAddrChange(handle *Handle, overlapped *Overlapped) error {
	r1, _, e1 := notifyAddrChange.Call(uintptr(unsafe.Pointer(handle)), uintptr(unsafe.Pointer(overlapped)))
	if handle == nil && overlapped == nil {
		if r1 == NO_ERROR {
			return nil
		}
	} else {
		if r1 == ERROR_IO_PENDING {
			return nil
		}
	}

	if e1 != ERROR_SUCCESS {
		return e1
	} else {
		return fmt.Errorf("r1:%v", r1)
	}
}

//DWORD NotifyRouteChange(
//  PHANDLE      Handle,
//  LPOVERLAPPED overlapped
//);
//https://docs.microsoft.com/en-us/windows/desktop/api/iphlpapi/nf-iphlpapi-notifyroutechange
func NotifyRouteChange(handle *Handle, overlapped *Overlapped) error {
	r1, _, e1 := notifyRouteChange.Call(uintptr(unsafe.Pointer(handle)), uintptr(unsafe.Pointer(overlapped)))
	if handle == nil && overlapped == nil {
		if r1 == NO_ERROR {
			return nil
		}
	} else {
		if r1 == ERROR_IO_PENDING {
			return nil
		}
	}

	if e1 != ERROR_SUCCESS {
		return e1
	} else {
		return fmt.Errorf("r1:%v", r1)
	}
}

// BOOL CancelIPChangeNotify(
//  LPOVERLAPPED notifyOverlapped
//);
// https://docs.microsoft.com/zh-cn/windows/desktop/api/iphlpapi/nf-iphlpapi-cancelipchangenotify
// 返回值：
//		bool 	如果当前没有 NotifyAddrChange 或 NotifyRouteChange 调用或 overlapped 无效，返回 false
func CancelIPChangeNotify(overlapped *Overlapped) (bool, error) {
	r1, _, _ := cancelIPChangeNotify.Call(uintptr(unsafe.Pointer(overlapped)))
	if r1 == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

type IPChangeNotify struct {
	rwm          sync.RWMutex
	ctx          context.Context
	ctxCancel    func()
	hasAddr      bool
	hasRoute     bool
	addrOverlap  *Overlapped
	routeOverlap *Overlapped
	//addrHand     Handle // 指向HANDLE变量的指针，该变量接收在异步通知中使用的句柄。
	//routeHand    Handle // 指向HANDLE变量的指针，该变量接收在异步通知中使用的句柄。
	C chan *IPChangeNotifyChanData
}

type IPChangeNotifyChanData struct {
	Err     error
	IsAddr  bool
	IsRoute bool
}

func (n *IPChangeNotify) close() error {
	if n.ctx != nil {
		select {
		case <-n.ctx.Done():
			break
		default:
			if f := n.ctxCancel; f != nil {
				f()
			}
		}
	}

	if overlap := n.routeOverlap; overlap != nil {
		CancelIPChangeNotify(overlap)
		WSACloseEvent(WSAEvent(overlap.HEvent))
	}

	if overlap := n.addrOverlap; overlap != nil {
		CancelIPChangeNotify(overlap)
		WSACloseEvent(WSAEvent(overlap.HEvent))
	}

	n.addrOverlap = &Overlapped{}
	n.routeOverlap = &Overlapped{}
	n.hasRoute = false
	n.hasAddr = false
	return nil
}

func (n *IPChangeNotify) Close() error {
	n.rwm.Lock()
	defer n.rwm.Unlock()

	return n.close()
}

func NewIPChangeNotify(hasAddr, hasRoute bool) (*IPChangeNotify, error) {
	n := new(IPChangeNotify)
	err := n.Reset(hasAddr, hasRoute)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (n *IPChangeNotify) Done() <-chan struct{} {
	n.rwm.RLock()
	defer n.rwm.RUnlock()

	if n.ctx == nil {
		return nil
	}

	return n.ctx.Done()
}

func (n *IPChangeNotify) Reset(hasAddr, hasRoute bool) error {
	n.rwm.Lock()
	defer n.rwm.Unlock()

	// 关闭可能存在的
	n.close()

	c := n.C
	if c == nil {
		c = make(chan *IPChangeNotifyChanData, 1)
		n.C = c
	}

	ctx, ctxCancel := context.WithCancel(context.Background())

	n.ctx = ctx
	n.ctxCancel = ctxCancel

	cancel := false
	defer func() {
		if cancel {
			n.close()
		}
	}()

	if hasAddr {
		hEvent, err := WSACreateEvent()
		if err != nil {
			cancel = true
			return err
		}
		n.addrOverlap.HEvent = windows.Handle(hEvent)
	}

	if hasRoute {
		hEvent, err := WSACreateEvent()
		if err != nil {
			cancel = true
			return err
		}
		n.routeOverlap.HEvent = windows.Handle(hEvent)
	}

	if hasAddr {
		overlap := n.addrOverlap
		go waitForSingleObjectLoop(ctx, ctxCancel, NotifyAddrChange, IPChangeNotifyChanData{IsAddr: true}, c, overlap)
	}
	if hasRoute {
		overlap := n.routeOverlap
		go waitForSingleObjectLoop(ctx, ctxCancel, NotifyRouteChange, IPChangeNotifyChanData{IsRoute: true}, c, overlap)
	}

	return nil
}

func waitForSingleObjectLoop(ctx context.Context, ctxCancel func(), f func(handle *Handle, overlapped *Overlapped) error, data IPChangeNotifyChanData, c chan *IPChangeNotifyChanData, overlap *Overlapped) {
	defer ctxCancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		lData := data

		hand := Handle(0)
		err := f(&hand, overlap)
		if err != nil {
			lData.Err = err
			select {
			case <-ctx.Done():
			default:
				select {
				case c <- &lData:
					return
				case <-ctx.Done():
				}
			}
		}

		event, err := WaitForSingleObject(overlap.HEvent, INFINITE)
		if err != nil {
			lData.Err = err
		}

		if event != WAIT_OBJECT_0 {
			lData.Err = fmt.Errorf("event = %v", event)
		}

		select {
		case <-ctx.Done():
			return
		default:
			select {
			case c <- &lData:
			case <-ctx.Done():
			}
		}

		if lData.Err != nil {
			return
		}
	}
}

func (s IfOperStatus) String() string {
	switch s {
	case IfOperStatusUp:
		return "Up"
	case IfOperStatusDown:
		return "Down"
	case IfOperStatusTesting:
		return "Testing"
	case IfOperStatusUnknown:
		return "Unknown"
	case IfOperStatusDormant:
		return "Dormant"
	case IfOperStatusNotPresent:
		return "NotPresent"
	case IfOperStatusLowerLayerDown:
		return "LowerLayerDown"
	default:
		return strconv.FormatUint((uint64)(s), 10)
	}
}

// https://docs.microsoft.com/zh-cn/windows/win32/api/iphlpapi/nf-iphlpapi-getipaddrtable
// IPHLPAPI_DLL_LINKAGE DWORD GetIpAddrTable(
//  PMIB_IPADDRTABLE pIpAddrTable,
//  PULONG           pdwSize,
//  BOOL             bOrder
//);
// 实测 windows 10 下无法获得断开网卡的ip
func GetIpAddrTable(order bool) ([]MibIpAddrRowW2k, error) {
	_order := 0
	if order {
		_order = 1
	}

	bufSize := 1024
	var buf []byte
	var r1 uintptr
	var e1 error
	for {
		buf = make([]byte, bufSize)

		r1, _, e1 = getIpAddrTable.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&bufSize)), uintptr(_order))
		if r1 == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		break
	}

	if r1 != NO_ERROR {
		if e1 != ERROR_SUCCESS {
			return nil, e1
		} else {
			return nil, fmt.Errorf("r1:%v", r1)
		}
	}

	table := (*MibIpAddrTable)(unsafe.Pointer(&buf[0]))
	rows := table.Table[:]
	err := ChangeSliceSize(&rows, int(table.NumEntries), int(table.NumEntries))
	if err != nil {
		return nil, fmt.Errorf("ChangeSliceSize, %v", err)
	}

	res := make([]MibIpAddrRowW2k, len(rows))
	copy(res, rows)
	return res, nil
}
