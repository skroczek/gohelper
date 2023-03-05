package boolean

import (
	"github.com/skroczek/gohelper/interfaces"
)

func Contains[T interfaces.Equaler[T]](a []T, v T) bool {
	for _, x := range a {
		if x.Equal(v) {
			return true
		}
	}
	return false
}
