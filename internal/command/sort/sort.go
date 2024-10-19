package sort

import (
	"fmt"

	"github.com/avazquez/gosorter/internal/file"
	"github.com/avazquez/gosorter/internal/service"
	"github.com/avazquez/gosorter/internal/sorting"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

const (
	expectedArgsNumber = 1
	indexFileNameArg   = 0
)

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
			return process(flags, args)
		},
	}

	addFlags(cmd)
	return cmd
}

func process(flags *flagsDTO, args []string) error {
	sorter, err := sorting.Factory(flags.algorithm)
	if err != nil {
		return errors.Wrapf(err, "error constructing sorter")
	}

	lines, err := file.LinesFromFile(args[indexFileNameArg])
	if err != nil {
		return errors.Wrapf(err, "error extracting lines from")
	}

	sortService := service.NewSortService(flags.unique, flags.topLimit, sorter, flags.descOrder)
	output, err := sortService.Process(lines)
	if err != nil {
		return errors.Wrapf(err, "error while processing")
	}

	for _, v := range output {
		fmt.Println(v)
	}
	return nil
}
