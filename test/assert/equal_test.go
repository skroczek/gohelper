package assert

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		Equal(t, 1, 1)
		NotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		Equal(t, "hello", "hello")
		NotEqual(t, "hello", "Grace")
	})

	t.Run("failed asserting on strings", func(t *testing.T) {
		tEqual := &testing.T{}
		Equal(tEqual, "hello", "hello1")
		if !tEqual.Failed() {
			t.Error("Equal() with not equal string args should fail")
		}
		tNotEqual := &testing.T{}
		NotEqual(tNotEqual, "hello", "hello")
		if !tNotEqual.Failed() {
			t.Error("NotEqual() with equal string args should fail")
		}
	})
}
