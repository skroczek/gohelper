package diff

import (
	"github.com/skroczek/gohelper/interfaces"
	"sync"
)

// Diff returns the added and removed elements between two slices. The slices must be of the same type and must
// implement the Equaler interface.
func Diff[T interfaces.Equaler[T]](a, b []T) (added, removed []T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		added = Added(a, b)
		wg.Done()
	}()
	go func() {
		removed = Removed(a, b)
		wg.Done()
	}()
	wg.Wait()
	return
}
