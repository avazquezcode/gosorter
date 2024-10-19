package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test all the sort methods against the same inputs
func TestSort(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected []string
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
	}

	for sorterName, sorter := range sorters {
		for name, test := range tests {
			t.Run(sorterName+"_"+name, func(t *testing.T) {
				got := sorter.Sort(test.input)
				assert.Equal(t, test.expected, got)
			})
		}
	}
}
