package gowindows

import (
	"golang.org/x/sys/windows"
)

var (
	ws2_32         = windows.NewLazyDLL("Ws2_32.dll")
	wSACreateEvent = ws2_32.NewProc("WSACreateEvent")
	wSACloseEvent  = ws2_32.NewProc("WSACloseEvent")
	wSAResetEvent  = ws2_32.NewProc("WSAResetEvent")
)

type WSAEvent Handle

const (
	WSA_INVALID_EVENT = WSAEvent(0)
)

// https://docs.microsoft.com/en-us/windows/desktop/api/winsock2/nf-winsock2-wsacreateevent
func WSACreateEvent() (WSAEvent, error) {
	r1, _, e1 := wSACreateEvent.Call()
	if WSAEvent(r1) == WSA_INVALID_EVENT {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, newCallError(DWord(r1))
		}
	}
	return WSAEvent(r1), nil
}

//BOOL WSAAPI WSACloseEvent(
//  WSAEVENT hEvent
//);
// https://docs.microsoft.com/zh-cn/windows/desktop/api/winsock2/nf-winsock2-wsacloseevent
func WSACloseEvent(event WSAEvent) error {
	r1, _, e1 := wSACloseEvent.Call(uintptr(event))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return newCallError(DWord(r1))
		}
	} else {
		return nil
	}
}

//BOOL WSAAPI WSAResetEvent(
//  WSAEVENT hEvent
//);
// https://docs.microsoft.com/zh-cn/windows/desktop/api/winsock2/nf-winsock2-wsaresetevent
func WSAResetEvent(event WSAEvent) error {
	r1, _, e1 := wSAResetEvent.Call(uintptr(event))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return newCallError(DWord(r1))
		}
	} else {
		return nil
	}
}
