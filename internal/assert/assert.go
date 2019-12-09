package assert

import (
	"reflect"
	"testing"
)

func Equal(t testing.TB, want, got interface{}) bool {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Logf("not equal\nwant=%v\ngot=%v", want, got)
		return false
	}

	return true
}
