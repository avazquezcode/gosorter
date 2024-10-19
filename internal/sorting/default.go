package sorting

import "sort"

type defaultSort struct{}

func (s defaultSort) Sort(items []string) []string {
	sortedItems := make([]string, len(items))
	copy(sortedItems, items)
	sort.Strings(sortedItems)
	return sortedItems
}
