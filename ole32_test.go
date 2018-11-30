package gowindows

import (
	"testing"
)

func TestGUIDFormString(t *testing.T) {
	guid, err := GUIDFormString("{3F2504E0-4F89-11D3-9A0C-0305E82C3301}")
	if err != nil {
		t.Fatal(err)
	}

	s, err := StringFromGUID2(&guid)
	if err != nil {
		t.Fatal(err)
	}

	if s != "{3F2504E0-4F89-11D3-9A0C-0305E82C3301}" {
		t.Fatal(s)
	}
}

func TestGUIDFormString2(t *testing.T) {
	guid, err := GUIDFormString("{FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF}")
	if err != nil {
		t.Fatal(err)
	}

	s, err := StringFromGUID2(&guid)
	if err != nil {
		t.Fatal(err)
	}

	if s != "{FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF}" {
		t.Fatal(s)
	}
}
