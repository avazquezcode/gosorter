package sorting

// insertionSort implements insertion sort algorithm
type insertionSort struct {
	options SortOptions
}

// Sort ...
func (s insertionSort) Sort(items []string) []string {
	n := len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && items[j] < items[j-1] {
			items[j], items[j-1] = items[j-1], items[j]
			j--
		}
	}

	return items
}
