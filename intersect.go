package generics

// Contains checks if a slice contains given value.
func Contains[T comparable](s []T, t T) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if v == t {
			return true
		}
	}

	return false
}
