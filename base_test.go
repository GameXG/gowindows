package gowindows

import "testing"

func TestChangeSliceSize(t *testing.T) {

	data := make([]int, 10)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	d1 := data[:2]

	err := ChangeSliceSize(&d1, 50, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(d1) != 50 {
		t.Errorf("%v", len(d1))
	}

	if cap(d1) != 100 {
		t.Error(cap(d1))
	}

	if d1[4] != 4 {
		t.Error(d1[4])
	}
}

func TestHRESULT_IsSucceeded(t *testing.T) {
	if HRESULT(FWP_E_ALREADY_EXISTS).IsSucceeded() == true {
		t.Errorf("")
	}
	if HRESULT(11).IsSucceeded() == false {
		t.Errorf("")
	}
}
