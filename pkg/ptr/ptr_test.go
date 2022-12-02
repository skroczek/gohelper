package ptr

import (
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		if got := Ptr("1"); reflect.TypeOf(got) != reflect.TypeOf(new(string)) {
			t.Errorf("Ptr() = %T, want %T", got, new(string))
		} else if *got != "1" {
			t.Errorf("Ptr() = %v, want %v", *got, "1")
		}
	})
	t.Run("int", func(t *testing.T) {
		if got := Ptr(1); reflect.TypeOf(got) != reflect.TypeOf(new(int)) {
			t.Errorf("Ptr() = %T, want %T", got, new(int))
		} else if *got != 1 {
			t.Errorf("Ptr() = %v, want %v", *got, 1)
		}
	})
	t.Run("uint", func(t *testing.T) {
		if got := Ptr(uint(1)); reflect.TypeOf(got) != reflect.TypeOf(new(uint)) {
			t.Errorf("Ptr() = %T, want %T", got, new(uint))
		} else if *got != uint(1) {
			t.Errorf("Ptr() = %v, want %v", *got, uint(1))
		}
	})
	t.Run("float64", func(t *testing.T) {
		if got := Ptr(1.0); reflect.TypeOf(got) != reflect.TypeOf(new(float64)) {
			t.Errorf("Ptr() = %T, want %T", got, new(float64))
		} else if *got != 1.0 {
			t.Errorf("Ptr() = %v, want %v", *got, 1.0)
		}
	})
}
