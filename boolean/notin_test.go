package boolean

import (
	"github.com/skroczek/gohelper/test/assert"
	"testing"
)

func TestNotIn(t *testing.T) {
	t.Run("Negative integer test", func(t *testing.T) {
		assert.Equal(t, NotIn(1, []int{1, 2, 3}...), false)
	})
	t.Run("Positive integer test", func(t *testing.T) {
		assert.Equal(t, NotIn(1, 2, 3), true)
	})
	t.Run("Negative string test", func(t *testing.T) {
		assert.Equal(t, NotIn("a", "a", "b", "c"), false)
	})
	t.Run("Positive string test", func(t *testing.T) {
		assert.Equal(t, NotIn("a", "b", "c"), true)
	})
	t.Run("Negative rune test", func(t *testing.T) {
		assert.Equal(t, NotIn('a', 'a', 'b', 'c'), false)
	})
	t.Run("Positive rune test", func(t *testing.T) {
		assert.Equal(t, NotIn('a', 'b', 'c'), true)
	})
	t.Run("Negative boolean test", func(t *testing.T) {
		assert.Equal(t, NotIn(true, true, false), false)
	})
	t.Run("Positive boolean test", func(t *testing.T) {
		assert.Equal(t, NotIn(true, false), true)
	})
	t.Run("Negative comparable struct test", func(t *testing.T) {
		assert.Equal(t, NotIn(struct {
			a int
		}{1}, struct {
			a int
		}{1}, struct {
			a int
		}{2}), false)
	})
	t.Run("Positive comparable struct test", func(t *testing.T) {
		assert.Equal(t, NotIn(struct {
			a int
		}{1}, struct {
			a int
		}{3}, struct {
			a int
		}{2}), true)
	})
}
