package cmd

import (
	"github.com/spf13/cobra"
)

type packageInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Manager string `json:"type"`
	Source  string `json:"source"`
}

var (
	// allowed values: apt, rpm, pip, all
	typeFl []string

	packagesCmd = &cobra.Command{
		Use:     "packages",
		Aliases: []string{"pkgs"},
		Short:   "A group of commands related to packages in an image",
	}
)

func init() {

	packagesCmd.PersistentFlags().StringSliceVar(&typeFl, "type", []string{"apt"}, "choose type of packages to list (apt, rpm, pip, or all)")
	rootCmd.AddCommand(packagesCmd)
}

func typeRequested(types []string, request string) bool {

	for _, aType := range types {

		if aType == request {
			return true
		}
	}

	return false
}
