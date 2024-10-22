package sort

import (
	"github.com/avazquezcode/gosorter/internal/domain/sorting"
	"github.com/avazquezcode/gosorter/internal/service/sort"
	"github.com/avazquezcode/gosorter/internal/utils/file"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewCommand creates the sort CLI command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sort [file_path] ... [OPTIONAL file_N_path]",
		Short: "sorts lines of a file (or lines of multiple files merged alltogether)",
		Args:  cobra.MinimumNArgs(1), // we need at least 1 filename
		RunE: func(cmd *cobra.Command, args []string) error {
			flags, err := parseFlags(cmd)
			if err != nil {
				return errors.Wrapf(err, "error parsing flags")
			}
			return process(cmd, flags, args)
		},
	}

	addFlags(cmd)
	return cmd
}

func process(cmd *cobra.Command, flags *flagsDTO, args []string) error {
	sorter, err := createSorter(flags)
	if err != nil {
		return errors.Wrapf(err, "error creating the sorter")
	}

	lines := make([]string, 0)

	for _, fileName := range args {
		fileLines, err := file.LinesFromFile(fileName)
		if err != nil {
			return errors.Wrapf(err, "error extracting lines from file: %s", fileName)
		}
		lines = append(lines, fileLines...)
	}

	sortSvc := createSortService(flags, sorter)
	output := sortSvc.Process(lines)

	for _, v := range output {
		cmd.Println(v)
	}

	return nil
}

func createSorter(flags *flagsDTO) (sorting.Sorter, error) {
	sortOptions := sorting.SortOptions{
		IgnoreCase: flags.ignoreCase,
	}

	sorter, err := sorting.Factory(flags.algorithm, sortOptions)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing sorter")
	}

	return sorter, nil
}

func createSortService(flags *flagsDTO, sorter sorting.Sorter) *sort.Service {
	parameters := sort.Parameters{
		RemoveDuplicates: flags.unique,
		DescOrder:        flags.descOrder,
		TopLimit:         flags.topLimit,
		IgnoreEmptyLines: flags.ignoreEmptyLines,
	}

	return sort.NewService(parameters, sorter)
}
