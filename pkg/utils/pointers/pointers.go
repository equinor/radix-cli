package pointers

func Val[T any](i *T) T {
	var value T
	if i != nil {
		value = *i
	}

	return value
}
