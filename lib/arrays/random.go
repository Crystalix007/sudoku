package arrays

import "math/rand"

func RandomIndex[T any](v []T) uint {
	if len(v) == 0 {
		return 0
	}

	return uint(rand.Int63n(int64(len(v))))
}
