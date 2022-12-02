package assert

import (
	"reflect"
	"testing"
)

func NotEqual[T comparable](t *testing.T, a, b T) {
	if reflect.DeepEqual(a, b) {
		t.Errorf("didn't want %v", b)
	}
}
