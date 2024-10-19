package service_test

import (
	"testing"

	"github.com/avazquezcode/gosorter/internal/service"
	"github.com/avazquezcode/gosorter/internal/sorting"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	tests := map[string]struct {
		removeDuplicates bool
		topLimit         int
		expectsDescOrder bool
		input            []string
		expected         []string
		expectsErr       bool
	}{
		"base case (default values)": {
			input:      []string{"b", "c", "a", "a"},
			expected:   []string{"a", "a", "b", "c"},
			expectsErr: false,
		},
		"removing duplicates (default values)": {
			removeDuplicates: true,
			input:            []string{"b", "c", "a", "a"},
			expected:         []string{"a", "b", "c"},
			expectsErr:       false,
		},
		"limiting output": {
			topLimit:   2,
			input:      []string{"b", "c", "a"},
			expected:   []string{"a", "b"},
			expectsErr: false,
		},
		"desc order is expected": {
			input:            []string{"b", "c", "a"},
			expected:         []string{"c", "b", "a"},
			expectsDescOrder: true,
			expectsErr:       false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sorter, _ := sorting.Factory("default")
			sortSvc := service.NewSortService(test.removeDuplicates, test.topLimit, sorter, test.expectsDescOrder)
			output, err := sortSvc.Process(test.input)
			assert.Equal(t, test.expected, output)
			assert.Equal(t, test.expectsErr, err != nil)
		})
	}
}
