package arrays

func RemoveIndex[T any](index uint, vals []T) []T {
	if index > uint(len(vals)) {
		return vals
	}

	return append(vals[:index], vals[index+1:]...)
}
