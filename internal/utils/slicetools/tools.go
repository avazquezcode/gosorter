package slicetools

// Reverse returns the reversed list (it doesn't cause a side effect in the input)
func Reverse[T any](list []T) []T {
	reversed := make([]T, len(list))
	copy(reversed, list)

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return reversed
}
