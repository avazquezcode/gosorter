package sorting

import (
	"sort"
	"strings"
)

type defaultSort struct {
	options SortOptions
}

func (s defaultSort) Sort(items []string) []string {
	// copying to avoid side effects on the input
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
