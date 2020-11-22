package cmd

import (
	"github.com/spf13/cobra"
)

var (
	packagesCmd = &cobra.Command{
		Use:     "packages",
		Aliases: []string{"pkgs"},
		Short:   "A group of commands related to packages in an image",
	}
)

func init() {

	rootCmd.AddCommand(packagesCmd)
}
