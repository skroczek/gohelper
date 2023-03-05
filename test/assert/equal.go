package assert

import (
	"reflect"
	"testing"
)

func Equal[T any](t *testing.T, a, b T) {
	if reflect.DeepEqual(a, b) {
		t.Errorf("expected %v, got %v", b, a)
	}
}
