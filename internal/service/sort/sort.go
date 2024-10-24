package sort

import (
	"github.com/avazquezcode/gosorter/internal/domain/sorting"
	"github.com/avazquezcode/gosorter/internal/utils/slicetools"
)

type (
	// Parameters are the parameters that modify the behaviour of the service
	Parameters struct {
		RemoveDuplicates bool
		TopLimit         int
		DescOrder        bool
		IgnoreEmptyLines bool
	}

	// Service is the struct used to represent the service
	Service struct {
		parameters Parameters
		sorter     sorting.Sorter
	}
)

func NewService(parameters Parameters, sorter sorting.Sorter) *Service {
	return &Service{
		parameters: parameters,
		sorter:     sorter,
	}
}

// Process executes the sort process
func (s *Service) Process(lines []string) []string {
	lines = s.preProcess(lines)

	sortedLines := s.sorter.Sort(lines)

	if s.parameters.DescOrder {
		sortedLines = slicetools.Reverse(sortedLines)
	}

	if len(sortedLines) > 0 && s.parameters.TopLimit > 0 {
		sortedLines = sortedLines[0:min(len(sortedLines), s.parameters.TopLimit)]
	}

	return sortedLines
}

// preProcess prepares the array for the sorting
// it takes care of cleaning it up based on the parameters passed by the user
func (s *Service) preProcess(lines []string) []string {
	var result []string
	set := make(map[string]struct{}, 0)
	for _, line := range lines {
		if s.parameters.RemoveDuplicates {
			if _, exist := set[line]; exist {
				continue
			}
		}

		removableLines := s.removableLines()
		if _, shouldRemove := removableLines[line]; shouldRemove {
			continue
		}

		set[line] = struct{}{}
		result = append(result, line)
	}

	return result
}

func (s *Service) removableLines() map[string]struct{} {
	lines := make(map[string]struct{}, 0)

	if s.parameters.IgnoreEmptyLines {
		lines[""] = struct{}{}
	}

	return lines
}
