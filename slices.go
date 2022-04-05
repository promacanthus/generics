package generics

func Filter[T any](s []T, f func(T) bool) []T {
	var res []T
	if len(s) == 0 {
		return res
	}

	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
