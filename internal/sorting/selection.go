package sorting

import "strings"

type selectionSort struct {
	options SortOptions
}

func (s selectionSort) Sort(items []string) []string {
	n := len(items)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if s.less(items[j], items[minIdx]) {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}

	return items
}

func (s selectionSort) less(a string, b string) bool {
	if s.options.IgnoreCase {
		return strings.ToLower(a) < strings.ToLower(b)
	}
	return a < b
}
