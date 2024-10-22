package sorting

import "strings"

// insertionSort implements insertion sort algorithm
type insertionSort struct {
	options SortOptions
}

// Sort sorts the input
func (s insertionSort) Sort(items []string) []string {
	// make a copy to avoid side effects on the input
	itemsCp := make([]string, len(items))
	copy(itemsCp, items)

	n := len(itemsCp)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && s.less(itemsCp[j], itemsCp[j-1]) {
			itemsCp[j], itemsCp[j-1] = itemsCp[j-1], itemsCp[j]
			j--
		}
	}

	return itemsCp
}

func (s insertionSort) less(a string, b string) bool {
	if s.options.IgnoreCase {
		return strings.ToLower(a) < strings.ToLower(b)
	}
	return a < b
}
