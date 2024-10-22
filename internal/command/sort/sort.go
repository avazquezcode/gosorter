package sort

import (
	"github.com/avazquezcode/gosorter/internal/file"
	"github.com/avazquezcode/gosorter/internal/service/sort"
	"github.com/avazquezcode/gosorter/internal/sorting"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

const (
	expectedArgsNumber = 1
	indexFileNameArg   = 0
)

// NewCommand creates the sort CLI command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sort",
		Short: "sorts lines of a file",
		Args:  cobra.ExactArgs(expectedArgsNumber),
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

	lines, err := file.LinesFromFile(args[indexFileNameArg])
	if err != nil {
		return errors.Wrapf(err, "error extracting lines from file")
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
	}

	return sort.NewService(parameters, sorter)
}
