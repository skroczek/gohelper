package boolean

func NotIn[T comparable](a T, b ...T) bool {
	for _, v := range b {
		if a == v {
			return false
		}
	}
	return true
}
