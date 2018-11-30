// +build windows

package gowindows

import "testing"

func TestMessageBox(t *testing.T) {
	/* 这个测试注释掉，会阻塞自动化测试功能
	MessageBox("Cloud.exe - Unable To Locate Component",
		"This application has failed to start because Proxy Gate's configuration file was not found. Re-installing Proxy Gate may fix this problem.",
		uintptr(MB_ICONHAND))*/
}

func TestSendMessageData(t *testing.T) {
	/*
		h, err := FindWindow("", "接收消息")
		if err != nil {
			t.Fatal(err)
		}

		err = SendMessageData(h, []byte("123abc中文\000"))
		if err != nil {
			t.Fatal(err)
		}*/
}
