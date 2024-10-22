package sort_test

import (
	"testing"

	"github.com/avazquezcode/gosorter/internal/service/sort"
	"github.com/avazquezcode/gosorter/internal/sorting"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	tests := map[string]struct {
		parameters sort.Parameters
		input      []string
		expected   []string
		expectsErr bool
	}{
		"base case (default values)": {
			input:      []string{"b", "c", "a", "a"},
			expected:   []string{"a", "a", "b", "c"},
			expectsErr: false,
		},
		"removing duplicates (default values)": {
			parameters: sort.Parameters{
				RemoveDuplicates: true,
			},
			input:      []string{"b", "c", "a", "a"},
			expected:   []string{"a", "b", "c"},
			expectsErr: false,
		},
		"limiting output": {
			parameters: sort.Parameters{
				TopLimit: 2,
			},
			input:      []string{"b", "c", "a"},
			expected:   []string{"a", "b"},
			expectsErr: false,
		},
		"desc order is expected": {
			input:    []string{"b", "c", "a"},
			expected: []string{"c", "b", "a"},
			parameters: sort.Parameters{
				DescOrder: true,
			},
			expectsErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sorter, _ := sorting.Factory("default", sorting.SortOptions{})
			sortSvc := sort.NewService(test.parameters, sorter)
			output, err := sortSvc.Process(test.input)
			assert.Equal(t, test.expected, output)
			assert.Equal(t, test.expectsErr, err != nil)
		})
	}
}
