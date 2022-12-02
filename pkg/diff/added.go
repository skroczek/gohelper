package diff

import (
	"github.com/skroczek/gohelper/pkg/boolean"
	"github.com/skroczek/gohelper/pkg/interfaces"
)

func Added[T interfaces.Equaler[T]](a, b []T) (added []T) {
	for _, v := range b {
		if !boolean.Contains(a, v) {
			added = append(added, v)
		}
	}
	return
}
