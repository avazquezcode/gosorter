package string_slice

// RemoveDuplicates removes duplicates from a slice (it doesn't cause a side effect in the input)
func RemoveDuplicates(list []string) []string {
	result := []string{}
	set := make(map[string]struct{}, 0)
	for _, item := range list {
		if _, exist := set[item]; exist {
			continue
		}
		set[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// Reverse returns the reversed list (it doesn't cause a side effect in the input)
func Reverse(list []string) []string {
	reversed := make([]string, len(list))
	copy(reversed, list)

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return reversed
}
