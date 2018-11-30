// +build windows

package gowindows

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32      = syscall.NewLazyDLL("user32.dll")
	messageBox  = user32.NewProc("MessageBoxW")
	findWindow  = user32.NewProc("FindWindowW")
	sendMessage = user32.NewProc("SendMessageW")
)

func MessageBox(caption, text string, style uintptr) (result int, err error) {
	// var hwnd HWND

	ret, _, callErr := messageBox.Call(
		0, // HWND
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),    // Text
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))), // Caption
		style, // type
	)
	//TODO: 需要处理 err
	if err != nil {
		return 0, fmt.Errorf("Call MessageBox:%v", callErr)
	}
	result = int(ret)
	return
}

func FindWindow(className, windowsName string) (result Handle, err error) {

	class := uintptr(0)
	win := uintptr(0)

	if len(className) != 0 {
		class = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(className)))
	}
	if len(windowsName) != 0 {
		win = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowsName)))
	}

	ret, _, callErr := findWindow.Call(class, win)
	//TODO: 需要处理 err

	if ret == 0 {
		return 0, fmt.Errorf("Call MessageBox:%v", callErr)
	}

	return Handle(ret), nil
}

func SendMessage(hWnd Handle, msg uint32, wParam, lParam uintptr) error {
	_, _, _ = sendMessage.Call(
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)
	//TODO: 需要处理 err

	return nil
}

func SendMessageData(hWnd Handle, data []byte) error {
	pCopyData := new(COPYDATASTRUCT)
	pCopyData.dwData = 0
	pCopyData.cbData = uint32(len(data))
	pCopyData.lpData = uintptr(unsafe.Pointer(&data[0]))

	return SendMessage(hWnd, WM_COPYDATA, 0, uintptr(unsafe.Pointer(pCopyData)))
}
