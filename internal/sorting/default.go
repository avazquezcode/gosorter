package sorting

import (
	"sort"
	"strings"
)

// defaultSort uses the strings library to sort the data
type defaultSort struct {
	options SortOptions
}

// Sort sorts the input
func (s defaultSort) Sort(items []string) []string {
	// make a copy to avoid side effects on the input
	sortedItems := make([]string, len(items))
	copy(sortedItems, items)

	sort.Slice(sortedItems, func(i, j int) bool {
		return s.less(sortedItems[i], sortedItems[j])
	})
	return sortedItems
}

func (s defaultSort) less(a string, b string) bool {
	if s.options.IgnoreCase {
		return strings.ToLower(a) < strings.ToLower(b)
	}
	return a < b
}
