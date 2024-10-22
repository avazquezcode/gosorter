package slicetools

// RemoveDuplicates removes duplicates from a slice (it doesn't cause a side effect in the input)
func RemoveDuplicates[T comparable](list []T) []T {
	result := []T{}
	set := make(map[T]struct{}, 0)
	for _, item := range list {
		if _, exist := set[item]; exist {
			continue
		}
		set[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// RemoveIfMatches removes items from a slice if matches to any item in the removableItems list (it doesn't cause a side effect in the input)
func RemoveIfMatches[T comparable](list []T, removableItems []T) []T {
	result := []T{}
	set := make(map[T]struct{}, 0)
	for _, item := range list {
		shouldRemove := false
		for _, removable := range removableItems {
			if item == removable {
				shouldRemove = true
				break
			}
		}

		if shouldRemove {
			continue
		}

		set[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// Reverse returns the reversed list (it doesn't cause a side effect in the input)
func Reverse[T any](list []T) []T {
	reversed := make([]T, len(list))
	copy(reversed, list)

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return reversed
}
