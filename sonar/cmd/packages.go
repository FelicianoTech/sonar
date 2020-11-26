package cmd

import (
	"github.com/spf13/cobra"
)

type packageInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Manager string `json:"manager"`
	Source  string `json:"source"`
}

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
