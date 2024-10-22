package sorting

import "sort"

type defaultSort struct {
	options SortOptions
}

func (s defaultSort) Sort(items []string) []string {
	sortedItems := make([]string, len(items))
	copy(sortedItems, items)
	sort.Strings(sortedItems)
	return sortedItems
}
