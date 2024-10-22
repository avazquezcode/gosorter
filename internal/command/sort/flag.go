package sort

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	flagNameUnique           = "unique"
	flagNameTop              = "top"
	flagNameMethod           = "method"
	flagNameOrder            = "order"
	flagNameIgnoreCase       = "ignore-case"
	flagNameIgnoreEmptyLines = "ignore-empty-lines"
)

type (
	flagsDTO struct {
		unique           bool
		topLimit         int
		algorithm        string
		descOrder        bool
		ignoreCase       bool
		ignoreEmptyLines bool
	}
)

func addFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool(flagNameUnique, false, "if true, only unique values will be consider in the output")
	cmd.PersistentFlags().Int(flagNameTop, 0, "used to limit the items shown in the output to the first <N> elements")
	cmd.PersistentFlags().String(flagNameMethod, "default", "refers to the sorting algorithm used to perform the sorting operation (default|insertion|selection|mergesort)")
	cmd.PersistentFlags().String(flagNameOrder, "asc", "indicates the order of the sorted output (asc|desc)")
	cmd.PersistentFlags().Bool(flagNameIgnoreCase, false, "if true, the char case will be ignored in the comparison (i.e: 'AVA' = 'ava')")
	cmd.PersistentFlags().Bool(flagNameIgnoreEmptyLines, true, "if true, empty lines are removed from the output")
}

func parseFlags(cmd *cobra.Command) (*flagsDTO, error) {
	unique, err := cmd.Flags().GetBool(flagNameUnique)
	if err != nil {
		return nil, err
	}

	topLimit, err := cmd.Flags().GetInt(flagNameTop)
	if err != nil {
		return nil, err
	}

	algorithm, err := cmd.Flags().GetString(flagNameMethod)
	if err != nil {
		return nil, err
	}

	isDescOrder, err := isDescOrder(cmd)
	if err != nil {
		return nil, err
	}

	ignoreCase, err := cmd.Flags().GetBool(flagNameIgnoreCase)
	if err != nil {
		return nil, err
	}

	ignoreEmptyLines, err := cmd.Flags().GetBool(flagNameIgnoreEmptyLines)
	if err != nil {
		return nil, err
	}

	return &flagsDTO{
		unique:           unique,
		topLimit:         topLimit,
		algorithm:        algorithm,
		descOrder:        *isDescOrder,
		ignoreCase:       ignoreCase,
		ignoreEmptyLines: ignoreEmptyLines,
	}, nil
}

func isDescOrder(cmd *cobra.Command) (*bool, error) {
	sortingOrder, err := cmd.Flags().GetString(flagNameOrder)
	if err != nil {
		return nil, err
	}

	switch sortingOrder {
	case "desc":
		trueVal := true
		return &trueVal, nil
	case "asc":
		falseVal := false
		return &falseVal, nil
	default:
		return nil, fmt.Errorf("the sorting order: %s is not valid", sortingOrder)
	}
}
