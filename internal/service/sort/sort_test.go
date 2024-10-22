package sort_test

import (
	"testing"

	"github.com/avazquezcode/gosorter/internal/domain/sorting"
	"github.com/avazquezcode/gosorter/internal/service/sort"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	tests := map[string]struct {
		parameters sort.Parameters
		input      []string
		expected   []string
	}{
		"base case (default values)": {
			input:    []string{"b", "c", "a", "a"},
			expected: []string{"a", "a", "b", "c"},
		},
		"removing duplicates (default values)": {
			parameters: sort.Parameters{
				RemoveDuplicates: true,
			},
			input:    []string{"b", "c", "a", "a"},
			expected: []string{"a", "b", "c"},
		},
		"removing empty lines": {
			parameters: sort.Parameters{
				IgnoreEmptyLines: true,
			},
			input:    []string{"b", "c", "a", "", ""},
			expected: []string{"a", "b", "c"},
		},
		"limiting output": {
			parameters: sort.Parameters{
				TopLimit: 2,
			},
			input:    []string{"b", "c", "a"},
			expected: []string{"a", "b"},
		},
		"desc order is expected": {
			input:    []string{"b", "c", "a"},
			expected: []string{"c", "b", "a"},
			parameters: sort.Parameters{
				DescOrder: true,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sorter, _ := sorting.Factory("default", sorting.SortOptions{})
			sortSvc := sort.NewService(test.parameters, sorter)
			output := sortSvc.Process(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}
