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

func Factory(algorithm string) (Sorter, error) {
	sorter, exist := sorters[algorithm]
	if !exist {
		return nil, fmt.Errorf("algorithm: %s is not supported", algorithm)
	}

	return sorter, nil
}
