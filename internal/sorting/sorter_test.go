package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// this slice contains the available sorting "methods" that are tested in this test file
// if a new "sorter" (method) is added, it should also be added here
var methods = []string{"default", "insertion", "selection", "mergesort"}

// Test all the sort methods against the same inputs
func TestSort(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected []string
		options  SortOptions
	}{
		"empty slice": {
			input:    []string{},
			expected: []string{},
		},
		"already sorted slice": {
			input: []string{
				"ava",
				"ben",
				"carla",
			},
			expected: []string{
				"ava",
				"ben",
				"carla",
			},
		},
		"unsorted slice without duplicated items": {
			input: []string{
				"carla",
				"ben",
				"ava",
			},
			expected: []string{
				"ava",
				"ben",
				"carla",
			},
		},
		"unsorted slice with duplicated items": {
			input: []string{
				"carla",
				"ben",
				"ava",
				"ben",
				"ava",
			},
			expected: []string{
				"ava",
				"ava",
				"ben",
				"ben",
				"carla",
			},
		},
		"unsorted slice with difference in the 3d char": {
			input: []string{
				"aav",
				"aab",
				"aad",
				"aac",
			},
			expected: []string{
				"aab",
				"aac",
				"aad",
				"aav",
			},
		},
		"case-sensitive sorting": {
			options: SortOptions{
				IgnoreCase: false,
			},
			input: []string{
				"BBB",
				"aaa",
			},
			expected: []string{
				"BBB",
				"aaa",
			},
		},
		"case-insensitive sorting": {
			options: SortOptions{
				IgnoreCase: true,
			},
			input: []string{
				"BBB",
				"aaa",
			},
			expected: []string{
				"aaa",
				"BBB",
			},
		},
	}

	for _, method := range methods {
		for name, test := range tests {
			t.Run(method+"_"+name, func(t *testing.T) {
				sorter, _ := Factory(method, test.options)
				got := sorter.Sort(test.input)
				assert.Equal(t, test.expected, got)
			})
		}
	}
}
