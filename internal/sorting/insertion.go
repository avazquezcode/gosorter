package sorting

import "strings"

// insertionSort implements insertion sort algorithm
type insertionSort struct {
	options SortOptions
}

// Sort ...
func (s insertionSort) Sort(items []string) []string {
	n := len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && s.less(items[j], items[j-1]) {
			items[j], items[j-1] = items[j-1], items[j]
			j--
		}
	}

	return items
}

func (s insertionSort) less(a string, b string) bool {
	if s.options.IgnoreCase {
		return strings.ToLower(a) < strings.ToLower(b)
	}
	return a < b
}
