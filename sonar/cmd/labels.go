package cmd

import (
	"github.com/spf13/cobra"
)

var (
	labelsCmd = &cobra.Command{
		Use:   "labels",
		Short: "A group of commands related to Docker image labels",
	}
)

func init() {

	rootCmd.AddCommand(labelsCmd)
}
