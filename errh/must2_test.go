package errh

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMust2(t *testing.T) {
	t.Run("Do not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Must2() panicked: %v", r)
			}
		}()
		a, b := Must2(1, "a", nil)
		if reflect.TypeOf(a) != reflect.TypeOf(1) {
			t.Errorf("Must2() returned wrong type: %v (expected: %v)", reflect.TypeOf(a), reflect.TypeOf(1))
		}

		if reflect.TypeOf(b) != reflect.TypeOf("a") {
			t.Errorf("Must2() returned wrong type: %v (expected: %v)", reflect.TypeOf(b), reflect.TypeOf("a"))
		}
	})
	t.Run("Panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Must2() didn't panick: %v", r)
			}
		}()
		_, _ = Must2(1, "a", fmt.Errorf("err"))
	})
}
