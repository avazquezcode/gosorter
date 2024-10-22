package slicetools_test

import (
	"testing"

	"github.com/avazquezcode/gosorter/internal/utils/slicetools"

	"github.com/stretchr/testify/assert"
)

// TestRemoveDuplicates tests the RemoveDuplicates function
func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty input",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "duplicates are removed",
			input:    []int{1, 2, 2, 3, 4, 4, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "all duplicates",
			input:    []int{1, 1, 1, 1},
			expected: []int{1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slicetools.RemoveDuplicates(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestRemoveIfMatches(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		removableItems []int
		expected       []int
	}{
		{
			name:           "empty input",
			input:          []int{},
			removableItems: []int{},
			expected:       []int{},
		},
		{
			name:           "no characters to remove",
			input:          []int{1, 2, 3, 4},
			removableItems: []int{},
			expected:       []int{1, 2, 3, 4},
		},
		{
			name:           "remove items (no duplicates)",
			input:          []int{1, 2, 3, 4},
			removableItems: []int{1, 3},
			expected:       []int{2, 4},
		},
		{
			name:           "remove items (more than 1 occurrence)",
			input:          []int{1, 2, 2, 3, 4, 4, 4},
			removableItems: []int{1, 2},
			expected:       []int{3, 4, 4, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slicetools.RemoveIfMatches(test.input, test.removableItems)
			assert.Equal(t, test.expected, result)
		})
	}
}

// TestReverse tests the Reverse function
func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "empty input",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "base case",
			input:    []string{"a", "b", "c", "d"},
			expected: []string{"d", "c", "b", "a"},
		},
		{
			name:     "single element",
			input:    []string{"a"},
			expected: []string{"a"},
		},
		{
			name:     "palindrome",
			input:    []string{"a", "b", "a"},
			expected: []string{"a", "b", "a"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slicetools.Reverse(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
