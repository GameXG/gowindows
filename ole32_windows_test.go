package gowindows

import (
	"testing"
)

func TestStringFromGUID2(t *testing.T) {
	guid := GUID{}

	s, err := StringFromGUID2(&guid)
	if err != nil {
		t.Fatal(err)
	}

	if s != "{00000000-0000-0000-0000-000000000000}" {
		t.Fatalf("%v!={00000000-0000-0000-0000-000000000000}", s)
	}

	guid.Data1 = 123
	guid.Data2 = 456
	guid.Data3 = 789
	guid.Data4 = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}

	s, err = StringFromGUID2(&guid)
	if err != nil {
		t.Fatal(err)
	}

	if s != "{0000007B-01C8-0315-0102-030405060708}" {
		t.Fatalf("%v!={0000007B-01C8-0315-0102-030405060708}", s)
	}
}
