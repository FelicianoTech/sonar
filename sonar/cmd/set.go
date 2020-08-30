package cmd

import (
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the value of several Docker Hub settings",
}

func init() {
	rootCmd.AddCommand(setCmd)
}
