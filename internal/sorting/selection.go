package sorting

import "strings"

// selectionSort implements selection sort algorithm
type selectionSort struct {
	options SortOptions
}

// Sort sorts the input
func (s selectionSort) Sort(items []string) []string {
	// make a copy to avoid side effects on the input
	itemsCp := make([]string, len(items))
	copy(itemsCp, items)

	n := len(itemsCp)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if s.less(itemsCp[j], itemsCp[minIdx]) {
				minIdx = j
			}
		}
		itemsCp[i], itemsCp[minIdx] = itemsCp[minIdx], itemsCp[i]
	}

	return itemsCp
}

func (s selectionSort) less(a string, b string) bool {
	if s.options.IgnoreCase {
		return strings.ToLower(a) < strings.ToLower(b)
	}
	return a < b
}
