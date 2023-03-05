package errh

func Must2[T any, U any](v T, w U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return v, w
}
