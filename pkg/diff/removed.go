package diff

import (
	"github.com/skroczek/gohelper/pkg/boolean"
	"github.com/skroczek/gohelper/pkg/interfaces"
)

func Removed[T interfaces.Equaler[T]](a, b []T) (removed []T) {
	for _, v := range a {
		if !boolean.Contains(b, v) {
			removed = append(removed, v)
		}
	}
	return
}
