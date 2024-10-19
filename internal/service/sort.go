package service

import (
	"github.com/avazquezcode/gosorter/internal/sorting"
	"github.com/avazquezcode/gosorter/internal/utils/string_slice"
)

type SortService struct {
	removeDuplicates bool
	topLimit         int
	sorter           sorting.Sorter
	descOrder        bool
}

func NewSortService(removeDuplicates bool, topLimit int, sorter sorting.Sorter, descOrder bool) *SortService {
	return &SortService{
		removeDuplicates: removeDuplicates,
		topLimit:         topLimit,
		sorter:           sorter,
		descOrder:        descOrder,
	}
}

// Process executes the sort process
func (s *SortService) Process(lines []string) ([]string, error) {
	if s.removeDuplicates {
		lines = string_slice.RemoveDuplicates(lines)
	}

	sortedLines := s.sorter.Sort(lines)

	if s.descOrder {
		sortedLines = string_slice.Reverse(sortedLines)
	}

	if len(sortedLines) > 0 && s.topLimit > 0 {
		sortedLines = sortedLines[0:min(len(sortedLines), s.topLimit)]
	}

	return sortedLines, nil
}
