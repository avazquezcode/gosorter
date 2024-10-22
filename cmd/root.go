package cmd

import (
	"os"

	"github.com/avazquezcode/gosorter/internal/command/sort"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "gosorter",
	Short: "a simple cli tool that sort lines of files",
}

func init() {
	RootCmd.AddCommand(sort.NewCommand())
}
