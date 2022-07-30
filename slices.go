package generics

// Filter returns any list of elements according to the boolean value of the function.
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
