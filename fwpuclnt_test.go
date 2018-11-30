package gowindows

import (
	"testing"

	"unsafe"
)

func TestFwpmStruct(t *testing.T) {
	if ptrSize == 8 {
		if unsafe.Sizeof(FwpmSession0{}) != 72 {
			t.Errorf("FwpmSession0 %v!=72", unsafe.Sizeof(FwpmSession0{}))
		}
		if unsafe.Sizeof(FwpmDisplayData0{}) != 16 {
			t.Errorf("FwpmDisplayData0 %v!=16", unsafe.Sizeof(FwpmDisplayData0{}))
		}
		if unsafe.Sizeof(FwpmSublayer0{}) != 72 {
			t.Errorf("FwpmSublayer0 %v!=72", unsafe.Sizeof(FwpmSublayer0{}))
		}
		if unsafe.Sizeof(FwpByteBlob{}) != 16 {
			t.Errorf("FwpByteBlob %v!=16", unsafe.Sizeof(FwpByteBlob{}))
		}
		if unsafe.Sizeof(FwpmFilter0{}) != 200 {
			t.Errorf("FwpmFilter0 %v!=200", unsafe.Sizeof(FwpmFilter0{}))
		}
		if unsafe.Sizeof(FwpValue0{}) != 16 {
			t.Errorf("FwpValue0 %v!=16", unsafe.Sizeof(FwpValue0{}))
		}
		if unsafe.Sizeof(FwpmFilterCondition0{}) != 40 {
			t.Errorf("FwpmFilterCondition0 %v!=40", unsafe.Sizeof(FwpmFilterCondition0{}))
		}
		if unsafe.Sizeof(FwpConditionValue0{}) != 16 {
			t.Errorf("FwpConditionValue0 %v!=16", unsafe.Sizeof(FwpConditionValue0{}))
		}
		if unsafe.Sizeof(FwpmAction0{}) != 20 {
			t.Errorf("FwpmAction0 %v!=20", unsafe.Sizeof(FwpmAction0{}))
		}
	} else {
		if unsafe.Sizeof(FwpmSession0{}) != 48 {
			t.Errorf("FwpmSession0 %v!=48", unsafe.Sizeof(FwpmSession0{}))
		}
		if unsafe.Sizeof(FwpmDisplayData0{}) != 8 {
			t.Errorf("FwpmDisplayData0 %v!=8", unsafe.Sizeof(FwpmDisplayData0{}))
		}
		if unsafe.Sizeof(FwpmSublayer0{}) != 44 {
			t.Errorf("FwpmSublayer0 %v!=44", unsafe.Sizeof(FwpmSublayer0{}))
		}
		if unsafe.Sizeof(FwpByteBlob{}) != 8 {
			t.Errorf("FwpByteBlob %v!=8", unsafe.Sizeof(FwpByteBlob{}))
		}
		if unsafe.Sizeof(FwpmFilter0{}) != 152 {
			t.Errorf("FwpmFilter0 %v!=152", unsafe.Sizeof(FwpmFilter0{}))
		}

		if unsafe.Sizeof(FwpValue0{}) != 8 {
			t.Errorf("FwpValue0 %v!=8", unsafe.Sizeof(FwpValue0{}))
		}
		if unsafe.Sizeof(FwpmFilterCondition0{}) != 28 {
			t.Errorf("FwpmFilterCondition0 %v!=28", unsafe.Sizeof(FwpmFilterCondition0{}))
		}

		if unsafe.Sizeof(FwpConditionValue0{}) != 8 {
			t.Errorf("FwpConditionValue0 %v!=8", unsafe.Sizeof(FwpConditionValue0{}))
		}
		if unsafe.Sizeof(FwpmAction0{}) != 20 {
			t.Errorf("FwpmAction0 %v!=20", unsafe.Sizeof(FwpmAction0{}))
		}

		if unsafe.Offsetof(FwpmFilter0{}.DisplayData) != 16 {
			t.Errorf("FwpmFilter0{}.DisplayData %v != 16", unsafe.Offsetof(FwpmFilter0{}.DisplayData))
		}

		if unsafe.Offsetof(FwpmFilter0{}.Flags) != 24 {
			t.Errorf("FwpmFilter0{}.Flags %v != 24", unsafe.Offsetof(FwpmFilter0{}.Flags))
		}

		if unsafe.Offsetof(FwpmFilter0{}.ProviderKey) != 28 {
			t.Errorf("FwpmFilter0{}.ProviderKey %v != 28", unsafe.Offsetof(FwpmFilter0{}.ProviderKey))
		}

		if unsafe.Offsetof(FwpmFilter0{}.ProviderData) != 32 {
			t.Errorf("FwpmFilter0{}.ProviderData %v != 32", unsafe.Offsetof(FwpmFilter0{}.ProviderData))
		}

		if unsafe.Offsetof(FwpmFilter0{}.LayerKey) != 40 {
			t.Errorf("FwpmFilter0{}.LayerKey %v != 40", unsafe.Offsetof(FwpmFilter0{}.LayerKey))
		}

		if unsafe.Offsetof(FwpmFilter0{}.SubLayerKey) != 56 {
			t.Errorf("FwpmFilter0{}.SubLayerKey %v != 56", unsafe.Offsetof(FwpmFilter0{}.SubLayerKey))
		}

		if unsafe.Offsetof(FwpmFilter0{}.Weight) != 72 {
			t.Errorf("FwpmFilter0{}.Weight %v != 72", unsafe.Offsetof(FwpmFilter0{}.Weight))
		}

		if unsafe.Offsetof(FwpmFilter0{}.NumFilterConditions) != 80 {
			t.Errorf("FwpmFilter0{}.NumFilterConditions %v != 80", unsafe.Offsetof(FwpmFilter0{}.NumFilterConditions))
		}

		if unsafe.Offsetof(FwpmFilter0{}.FilterCondition) != 84 {
			t.Errorf("FwpmFilter0{}.FilterCondition %v != 84", unsafe.Offsetof(FwpmFilter0{}.FilterCondition))
		}

		if unsafe.Offsetof(FwpmFilter0{}.Action) != 88 {
			t.Errorf("FwpmFilter0{}.Action %v != 88", unsafe.Offsetof(FwpmFilter0{}.Action))
		}

		if unsafe.Offsetof(FwpmFilter0{}.ProviderContextKey) != 112 {
			t.Errorf("FwpmFilter0{}.ProviderContextKey %v != 112", unsafe.Offsetof(FwpmFilter0{}.ProviderContextKey))
		}

		if unsafe.Offsetof(FwpmFilter0{}.Reserved) != 128 {
			t.Errorf("FwpmFilter0{}.Reserved %v != 128", unsafe.Offsetof(FwpmFilter0{}.Reserved))
		}

		if unsafe.Offsetof(FwpmFilter0{}.FilterId) != 136 {
			t.Errorf("FwpmFilter0{}.FilterId %v != 136", unsafe.Offsetof(FwpmFilter0{}.FilterId))
		}

		if unsafe.Offsetof(FwpmFilter0{}.EffectiveWeight) != 144 {
			t.Errorf("FwpmFilter0{}.EffectiveWeight %v != 144", unsafe.Offsetof(FwpmFilter0{}.EffectiveWeight))
		}
	}

}

func TestG(t *testing.T) {
}
