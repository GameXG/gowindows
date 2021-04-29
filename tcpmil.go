package gowindows

import (
	"fmt"
	"net"
)

//typedef struct _MIB_TCPTABLE {
//    DWORD      dwNumEntries;
//    MIB_TCPROW table[ANY_SIZE];
//} MIB_TCPTABLE, *PMIB_TCPTABLE;

// The MIB_TCPTABLE2 structure contains a table of IPv4 TCP connections on the local computer.
//typedef struct _MIB_TCPTABLE2 {
//    DWORD       dwNumEntries;
//    MIB_TCPROW2 table[ANY_SIZE];
//} MIB_TCPTABLE2, *PMIB_TCPTABLE2;
// https://docs.microsoft.com/en-us/windows/win32/api/tcpmib/ns-tcpmib-mib_tcptable2
type MibTcpTable2 struct {
	NumEntries DWord
	Table      [ANY_SIZE]MibTcpRow2
}

// typedef struct _MIB_TCPROW2 {
//    DWORD dwState;
//    DWORD dwLocalAddr;
//    DWORD dwLocalPort;
//    DWORD dwRemoteAddr;
//    DWORD dwRemotePort;
//    DWORD dwOwningPid;
//    TCP_CONNECTION_OFFLOAD_STATE dwOffloadState;
//} MIB_TCPROW2, *PMIB_TCPROW2;
// https://docs.microsoft.com/en-us/windows/win32/api/tcpmib/ns-tcpmib-mib_tcprow2
type MibTcpRow2 struct {
	State        MibTcpRow2Status
	LocalAddr    DWord
	LocalPort    DWord
	RemoteAddr   DWord
	RemotePort   DWord
	OwningPid    DWord
	OffloadState TcpConnectionOffloadState
}

func (r *MibTcpRow2) GetLocalAddr() net.IP {
	return uint322Ip(r.LocalAddr)
}

func (r *MibTcpRow2) GetLocalPort() int {
	return int(Ntohs(uint16(r.LocalPort)))
}

func (r *MibTcpRow2) GetRemoteAddr() net.IP {
	return uint322Ip(r.RemoteAddr)
}

func (r *MibTcpRow2) GetRemotePort() int {
	return int(Ntohs(uint16(r.RemotePort)))
}

func (r *MibTcpRow2) String() string {
	if r == nil {
		return ""
	}

	return fmt.Sprintf("[%v]%v:%v-%v:%v %v", r.OwningPid, r.GetLocalAddr(), r.GetLocalPort(), r.GetRemoteAddr(), r.GetRemotePort(), r.State)
}

type MibTcpRow2Status DWord

// typedef enum {
//    MIB_TCP_STATE_CLOSED     =  1,
//    MIB_TCP_STATE_LISTEN     =  2,
//    MIB_TCP_STATE_SYN_SENT   =  3,
//    MIB_TCP_STATE_SYN_RCVD   =  4,
//    MIB_TCP_STATE_ESTAB      =  5,
//    MIB_TCP_STATE_FIN_WAIT1  =  6,
//    MIB_TCP_STATE_FIN_WAIT2  =  7,
//    MIB_TCP_STATE_CLOSE_WAIT =  8,
//    MIB_TCP_STATE_CLOSING    =  9,
//    MIB_TCP_STATE_LAST_ACK   = 10,
//    MIB_TCP_STATE_TIME_WAIT  = 11,
//    MIB_TCP_STATE_DELETE_TCB = 12,
//    //
//    // Extra TCP states not defined in the MIB
//    //
//    MIB_TCP_STATE_RESERVED      = 100
//} MIB_TCP_STATE;
// https://docs.microsoft.com/en-us/windows/win32/api/tcpmib/ns-tcpmib-mib_tcprow2
const (
	MIB_TCP_STATE_CLOSED     MibTcpRow2Status = 1
	MIB_TCP_STATE_LISTEN     MibTcpRow2Status = 2
	MIB_TCP_STATE_SYN_SENT   MibTcpRow2Status = 3
	MIB_TCP_STATE_SYN_RCVD   MibTcpRow2Status = 4
	MIB_TCP_STATE_ESTAB      MibTcpRow2Status = 5
	MIB_TCP_STATE_FIN_WAIT1  MibTcpRow2Status = 6
	MIB_TCP_STATE_FIN_WAIT2  MibTcpRow2Status = 7
	MIB_TCP_STATE_CLOSE_WAIT MibTcpRow2Status = 8
	MIB_TCP_STATE_CLOSING    MibTcpRow2Status = 9
	MIB_TCP_STATE_LAST_ACK   MibTcpRow2Status = 10
	MIB_TCP_STATE_TIME_WAIT  MibTcpRow2Status = 11
	MIB_TCP_STATE_DELETE_TCB MibTcpRow2Status = 12
)

func (s MibTcpRow2Status) String() string {
	switch s {
	case MIB_TCP_STATE_CLOSED:
		return "CLOSED"
	case MIB_TCP_STATE_LISTEN:
		return "LISTEN"
	case MIB_TCP_STATE_SYN_SENT:
		return "SYN_SENT"
	case MIB_TCP_STATE_SYN_RCVD:
		return "SYN_RCVD"
	case MIB_TCP_STATE_ESTAB:
		return "ESTABLISHED"
	case MIB_TCP_STATE_FIN_WAIT1:
		return "FIN_WAIT1"
	case MIB_TCP_STATE_FIN_WAIT2:
		return "FIN_WAIT2"
	case MIB_TCP_STATE_CLOSE_WAIT:
		return "CLOSE_WAIT"
	case MIB_TCP_STATE_CLOSING:
		return "CLOSING"
	case MIB_TCP_STATE_LAST_ACK:
		return "LAST_ACK"
	case MIB_TCP_STATE_TIME_WAIT:
		return "TIME_WAIT"
	case MIB_TCP_STATE_DELETE_TCB:
		return "DELETE_TCB"
	default:
		return fmt.Sprintf("%v", DWord(s))
	}
}

// //
//// Various Offload states a TCP connection can be in.
////
//typedef enum {
//    TcpConnectionOffloadStateInHost,
//    TcpConnectionOffloadStateOffloading,
//    TcpConnectionOffloadStateOffloaded,
//    TcpConnectionOffloadStateUploading,
//    TcpConnectionOffloadStateMax
//} TCP_CONNECTION_OFFLOAD_STATE, *PTCP_CONNECTION_OFFLOAD_STATE;
type TcpConnectionOffloadState int32

const (
	TcpConnectionOffloadStateInHost     TcpConnectionOffloadState = 0
	TcpConnectionOffloadStateOffloading TcpConnectionOffloadState = 1
	TcpConnectionOffloadStateOffloaded  TcpConnectionOffloadState = 2
	TcpConnectionOffloadStateUploading  TcpConnectionOffloadState = 3
	TcpConnectionOffloadStateMax        TcpConnectionOffloadState = 4
)

func (s TcpConnectionOffloadState) String() string {
	switch s {
	case TcpConnectionOffloadStateInHost:
		return "InHost"
	case TcpConnectionOffloadStateOffloading:
		return "Offloading"
	case TcpConnectionOffloadStateOffloaded:
		return "Offloaded"
	case TcpConnectionOffloadStateUploading:
		return "Uploading"
	case TcpConnectionOffloadStateMax:
		return "Max"
	default:
		return fmt.Sprintf("%v", uint32(s))
	}
}
