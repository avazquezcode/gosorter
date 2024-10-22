package sorting

import (
	"fmt"
)

var sorters = map[string]Sorter{
	"default":   defaultSort{},
	"selection": selectionSort{},
	"insertion": insertionSort{},
	"mergesort": mergeSort{},
}

func Factory(algorithm string, options SortOptions) (Sorter, error) {
	switch algorithm {
	case "default":
		return defaultSort{options}, nil
	case "selection":
		return selectionSort{options}, nil
	case "insertion":
		return insertionSort{options}, nil
	case "mergesort":
		return mergeSort{options}, nil
	default:
		return nil, fmt.Errorf("algorithm: %s is not supported", algorithm)
	}

}
