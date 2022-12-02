package interfaces

type Equaler[T any] interface {
	Equal(T) bool
}
