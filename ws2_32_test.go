package gowindows

import "testing"

func TestWSACreateEvent(t *testing.T) {
	e, err := WSACreateEvent()
	if err != nil {
		t.Fatal(err)
	}

	err = WSAResetEvent(e)
	if err != nil {
		t.Fatal(err)
	}

	err = WSACloseEvent(e)
	if err != nil {
		t.Fatal(err)
	}
}
