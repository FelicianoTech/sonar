package cmd

import (
	"github.com/spf13/cobra"
)

var layersCmd = &cobra.Command{
	Use:   "layers",
	Short: "A group of commands related to Docker image layers",
}

func init() {

	rootCmd.AddCommand(layersCmd)
}
