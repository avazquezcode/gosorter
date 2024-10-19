package file

import (
	"os"
	"strings"

	"github.com/pkg/errors"
)

// LinesFromFile returns the content of a file as a slice of strings
// where each item, represents a line in the file content
func LinesFromFile(fileName string) ([]string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "error when reading file")
	}
	return strings.Split(string(content), "\n"), nil
}
