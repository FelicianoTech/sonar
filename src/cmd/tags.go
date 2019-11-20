package cmd

import (
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "A group of commands related to Docker image tags",
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
