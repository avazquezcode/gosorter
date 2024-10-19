package sorting

// mergeSort implements merge sort algorithm
type mergeSort struct{}

// Sort ...
func (s mergeSort) Sort(items []string) []string {
	if len(items) < 2 {
		return items
	}

	half1 := s.Sort(items[:len(items)/2])
	half2 := s.Sort(items[len(items)/2:])
	return merge(half1, half2)
}

func merge(a []string, b []string) []string {
	output := []string{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			output = append(output, a[i])
			i++
		} else {
			output = append(output, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		output = append(output, a[i])
	}
	for ; j < len(b); j++ {
		output = append(output, b[j])
	}
	return output
}
