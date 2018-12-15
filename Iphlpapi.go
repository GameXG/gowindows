package gowindows

import (
	"fmt"
	"net"
)

const (
	GAA_FLAG_SKIP_UNICAST                = 0x0001
	GAA_FLAG_SKIP_ANYCAST                = 0x0002
	GAA_FLAG_SKIP_MULTICAST              = 0x0004
	GAA_FLAG_SKIP_DNS_SERVER             = 0x0008
	GAA_FLAG_INCLUDE_PREFIX              = 0x0010
	GAA_FLAG_SKIP_FRIENDLY_NAME          = 0x0020
	GAA_FLAG_INCLUDE_WINS_INFO           = 0x0040
	GAA_FLAG_INCLUDE_GATEWAYS            = 0x0080
	GAA_FLAG_INCLUDE_ALL_INTERFACES      = 0x0100
	GAA_FLAG_INCLUDE_ALL_COMPARTMENTS    = 0x0200
	GAA_FLAG_INCLUDE_TUNNEL_BINDINGORDER = 0x0400
)

// NET_LUID_LH
// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ns-ifdef-_net_luid_lh
//typedef union _NET_LUID_LH {
//ULONG64 Value;
//struct {
//ULONG64 Reserved : 24;
//ULONG64 NetLuidIndex : 24;
//ULONG64 IfType : 16;
//} Info;
//} NET_LUID_LH, *PNET_LUID_LH;
type NetLuidLh uint64

type IfLuid = NetLuidLh

type CompartmentId uint32
type NetworkGuid = GUID
type ConnectionType int32
type TunnelType int32
type IpAdapterDnsSuffix struct {
	Next   *IpAdapterDnsSuffix
	String [MAX_DNS_SUFFIX_STRING_LENGTH]uint16
}

const (
	NET_IF_CONNECTION_DEDICATED ConnectionType = 1
	NET_IF_CONNECTION_PASSIVE   ConnectionType = 2
	NET_IF_CONNECTION_DEMAND    ConnectionType = 3
	NET_IF_CONNECTION_MAXIMUM   ConnectionType = 4
)

const (
	TUNNEL_TYPE_NONE    TunnelType = 0
	TUNNEL_TYPE_OTHER   TunnelType = 1
	TUNNEL_TYPE_DIRECT  TunnelType = 2
	TUNNEL_TYPE_6TO4    TunnelType = 11
	TUNNEL_TYPE_ISATAP  TunnelType = 13
	TUNNEL_TYPE_TEREDO  TunnelType = 14
	TUNNEL_TYPE_IPHTTPS TunnelType = 15
)

const MAX_DHCPV6_DUID_LENGTH = 130
const MAX_DNS_SUFFIX_STRING_LENGTH = 256

const (
	ERROR_INSUFFICIENT_BUFFER = 122
	ERROR_NO_DATA             = 232
)

const (
	STATUS_WAIT_0           DWord = 0
	STATUS_ABANDONED_WAIT_0 DWord = 0x00000080
	STATUS_USER_APC         DWord = 0x000000C0

	WAIT_OBJECT_0      DWord = STATUS_WAIT_0 + 0
	WAIT_FAILED        DWord = 0xFFFFFFFF
	WAIT_ABANDONED     DWord = STATUS_ABANDONED_WAIT_0 + 0
	WAIT_ABANDONED_0   DWord = STATUS_ABANDONED_WAIT_0 + 0
	WAIT_IO_COMPLETION DWord = STATUS_USER_APC
)

type MibIpForwardRow struct {
	ForwardDest      [4]byte //目标网络
	ForwardMask      [4]byte //掩码
	ForwardPolicy    DWord   //ForwardPolicy:0x0
	ForwardNextHop   [4]byte //网关
	ForwardIfIndex   DWord   // 网卡索引 id
	ForwardType      DWord   //3 本地接口  4 远端接口
	ForwardProto     DWord   //3静态路由 2本地接口 5EGP网关
	ForwardAge       DWord   //存在时间 秒
	ForwardNextHopAS DWord   //下一跳自治域号码 0
	ForwardMetric1   DWord   //度量衡(跃点数)，根据 ForwardProto 不同意义不同。
	ForwardMetric2   DWord
	ForwardMetric3   DWord
	ForwardMetric4   DWord
	ForwardMetric5   DWord
}

type MibIpForwardTable struct {
	NumEntries DWord
	Table      [1]MibIpForwardRow //实际是 NumEntries 个
}

func (row *MibIpForwardRow) String() string {
	return fmt.Sprintf("%v/%v->%v Metric:%v", row.GetForwardDest(),
		row.GetForwardMask(), row.GetForwardNextHop(), row.ForwardMetric1)
}

func (row *MibIpForwardRow) GetForwardDest() net.IP {
	return net.IP(row.ForwardDest[:])
}

func (row *MibIpForwardRow) SetForwardDest(v net.IP) error {
	ipv4 := v.To4()
	if len(ipv4) != net.IPv4len {
		return fmt.Errorf("%v 不是ipv4地址。", v)
	}

	copy(row.ForwardDest[:], ipv4)
	return nil
}

func (row *MibIpForwardRow) GetForwardMask() net.IPMask {
	return net.IPMask(row.ForwardMask[:])
}

func (row *MibIpForwardRow) SetForwardMask(v net.IPMask) error {
	ipv4 := net.IP(v).To4()
	if len(ipv4) != net.IPv4len {
		return fmt.Errorf("%v 不是ipv4掩码。", v)
	}

	copy(row.ForwardMask[:], ipv4)
	return nil
}

func (row *MibIpForwardRow) GetForwardNextHop() net.IP {
	return net.IP(row.ForwardNextHop[:])
}

func (row *MibIpForwardRow) SetForwardNextHop(v net.IP) error {
	ipv4 := v.To4()
	if len(ipv4) != net.IPv4len {
		return fmt.Errorf("%v 不是ipv4地址。", v)
	}

	copy(row.ForwardNextHop[:], ipv4)
	return nil
}
