package sort

import (
	"github.com/avazquezcode/gosorter/internal/sorting"
	"github.com/avazquezcode/gosorter/internal/utils/slicetools"
)

type Service struct {
	parameters Parameters
	sorter     sorting.Sorter
}

func NewService(parameters Parameters, sorter sorting.Sorter) *Service {
	return &Service{
		parameters: parameters,
		sorter:     sorter,
	}
}

// Process executes the sort process
func (s *Service) Process(lines []string) ([]string, error) {
	if s.parameters.RemoveDuplicates {
		lines = slicetools.RemoveDuplicates(lines)
	}

	sortedLines := s.sorter.Sort(lines)

	if s.parameters.DescOrder {
		sortedLines = slicetools.Reverse(sortedLines)
	}

	if len(sortedLines) > 0 && s.parameters.TopLimit > 0 {
		sortedLines = sortedLines[0:min(len(sortedLines), s.parameters.TopLimit)]
	}

	return sortedLines, nil
}
