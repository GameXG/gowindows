package gowindows

import (
	"testing"
)

func TestSetSelfProcessPrivilege(t *testing.T) {
	err := SetSelfProcessPrivilege(SE_DEBUG_NAME, true)
	if err != nil {
		t.Fatal(err)
	}
	err = SetSelfProcessPrivilege(SE_DEBUG_NAME, false)
	if err != nil {
		t.Fatal(err)
	}

	err = SetSelfProcessPrivilege(SE_CREATE_GLOBAL_NAME, true)
	if err != nil {
		t.Fatal(err)
	}
	err = SetSelfProcessPrivilege(SE_CREATE_GLOBAL_NAME, false)
	if err != nil {
		t.Fatal(err)
	}

	err = SetSelfProcessPrivilege("000", false)
	if err == nil {
		t.Fatal("nil")
	}
}
