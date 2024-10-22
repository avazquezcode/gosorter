package sort_test

import (
	"bytes"
	"testing"

	"github.com/avazquezcode/gosorter/internal/command/sort"

	"github.com/stretchr/testify/assert"
)

func TestSortCmd(t *testing.T) {
	tests := map[string]struct {
		args           []string
		expectedOutput string
		expectsErr     bool
	}{
		"no arguments": {
			args:       []string{},
			expectsErr: true,
		},
		"invalid unique flag": {
			args:       []string{"testdata/test.txt", "--unique=string"},
			expectsErr: true,
		},
		"invalid top flag": {
			args:       []string{"testdata/test.txt", "--top=string"},
			expectsErr: true,
		},
		"invalid method flag (unsopported option)": {
			args:       []string{"testdata/test.txt", "--method=invalid_one"},
			expectsErr: true,
		},
		"invalid order flag (unsopported option)": {
			args:       []string{"testdata/test.txt", "--order=invalid_one"},
			expectsErr: true,
		},
		"invalid order flag (invalid type)": {
			args:       []string{"testdata/test.txt", "--order"},
			expectsErr: true,
		},
		"invalid ignore case flag": {
			args:       []string{"testdata/test.txt", "--ignore-case=string"},
			expectsErr: true,
		},
		"valid case - default flags": {
			args:           []string{"testdata/test.txt"},
			expectedOutput: "BBB\naaa\nb\nb\nc\n",
			expectsErr:     false,
		},
		"valid case - unique": {
			args:           []string{"testdata/test.txt", "--unique"},
			expectedOutput: "BBB\naaa\nb\nc\n",
			expectsErr:     false,
		},
		"valid case - top limit": {
			args:           []string{"testdata/test.txt", "--top=1"},
			expectedOutput: "BBB\n",
			expectsErr:     false,
		},
		"valid case - ignore case true": {
			args:           []string{"testdata/test.txt", "--ignore-case=true"},
			expectedOutput: "aaa\nb\nb\nBBB\nc\n",
			expectsErr:     false,
		},
		"valid case - sort output desc order": {
			args:           []string{"testdata/test.txt", "--order=desc"},
			expectedOutput: "c\nb\nb\naaa\nBBB\n",
			expectsErr:     false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var testStdOut bytes.Buffer
			cmd := sort.NewCommand()
			cmd.SetArgs(test.args)
			cmd.SetOut(&testStdOut)

			err := cmd.Execute()
			assert.Equal(t, test.expectsErr, err != nil)

			if err == nil {
				// check stdout
				assert.Equal(t, test.expectedOutput, testStdOut.String())
			}
		})
	}
}
