package file_test

import (
	"testing"

	"github.com/avazquezcode/gosorter/internal/utils/file"
	"github.com/stretchr/testify/assert"
)

func TestLinesFromFile(t *testing.T) {
	tests := map[string]struct {
		fileName   string
		expected   []string
		expectsErr bool
	}{
		"valid file": {
			fileName: "testdata/testfile.txt",
			expected: []string{"a", "b", "c", "d", ""},
		},
		"file does not exist": {
			fileName:   "testdata/testfile_404.txt",
			expectsErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := file.LinesFromFile(test.fileName)
			assert.Equal(t, test.expected, output)
			assert.Equal(t, test.expectsErr, err != nil)
		})
	}
}
