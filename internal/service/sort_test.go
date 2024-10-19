package service_test

import (
	"testing"

	"github.com/avazquez/gosorter/internal/service"
	"github.com/avazquez/gosorter/internal/sorting"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	tests := map[string]struct {
		removeDuplicates bool
		topLimit         int
		expectsAscOrder  bool
		input            []string
		expected         []string
		expectsErr       bool
	}{
		"base case (default values)": {
			input:           []string{"b", "c", "a", "a"},
			expected:        []string{"a", "a", "b", "c"},
			expectsAscOrder: true,
			expectsErr:      false,
		},
		"removing duplicates (default values)": {
			removeDuplicates: true,
			input:            []string{"b", "c", "a", "a"},
			expected:         []string{"a", "b", "c"},
			expectsAscOrder:  true,
			expectsErr:       false,
		},
		"limiting output": {
			topLimit:        2,
			input:           []string{"b", "c", "a"},
			expected:        []string{"a", "b"},
			expectsAscOrder: true,
			expectsErr:      false,
		},
		"desc order is expected": {
			input:           []string{"b", "c", "a"},
			expected:        []string{"c", "b", "a"},
			expectsAscOrder: false,
			expectsErr:      false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sorter, _ := sorting.Factory("default")
			sortSvc := service.NewSortService(test.removeDuplicates, test.topLimit, sorter, test.expectsAscOrder)
			output, err := sortSvc.Process(test.input)
			assert.Equal(t, test.expected, output)
			assert.Equal(t, test.expectsErr, err != nil)
		})
	}
}
