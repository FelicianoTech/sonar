package cmd

import (
	"github.com/spf13/cobra"
)

var (
	fieldFl string
	gtFl    string
	ltFl    string

	tagsCmd = &cobra.Command{
		Use:   "tags",
		Short: "A group of commands related to Docker image tags",
	}
)

func init() {

	tagsCmd.PersistentFlags().StringVar(&fieldFl, "field", "", "the field to filter what tags will be selected. Only 'date' is supported right now.")
	tagsCmd.PersistentFlags().StringVar(&gtFl, "gt", "", "filter tags 'greater than' this value - allowed values is based on the 'field' choosen. A relative time in seconds, minutes, or hours is supported.")
	tagsCmd.PersistentFlags().StringVar(&ltFl, "lt", "", "filter tags 'less than' this value - allowed values is based on the 'field' choosen. A relative time in seconds, minutes, or hours is supported.")

	rootCmd.AddCommand(tagsCmd)
}
