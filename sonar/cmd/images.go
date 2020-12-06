package cmd

import (
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "A group of commands related to Docker images",
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}
