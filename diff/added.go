package diff

import (
	"github.com/skroczek/gohelper/boolean"
	"github.com/skroczek/gohelper/interfaces"
)

func Added[T interfaces.Equaler[T]](a, b []T) (added []T) {
	for _, v := range b {
		if !boolean.Contains(a, v) {
			added = append(added, v)
		}
	}
	return
}
