package boolean

func In[T comparable](a T, b ...T) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}
