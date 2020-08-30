package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value of several Docker Hub metrics",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
