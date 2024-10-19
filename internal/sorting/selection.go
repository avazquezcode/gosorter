package sorting

type selectionSort struct{}

func (s selectionSort) Sort(items []string) []string {
	n := len(items)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}

	return items
}
