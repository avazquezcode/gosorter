package sorting

import (
	"sort"
	"strings"
)

type defaultSort struct {
	options SortOptions
}

func (s defaultSort) Sort(items []string) []string {
	sort.Slice(items, func(i, j int) bool {
		return s.less(items[i], items[j])
	})
	return items
}

func (s defaultSort) less(a string, b string) bool {
	if s.options.IgnoreCase {
		return strings.ToLower(a) < strings.ToLower(b)
	}
	return a < b
}
