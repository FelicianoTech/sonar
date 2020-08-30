package cmd

import (
	"strconv"
	"time"

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
	tagsCmd.PersistentFlags().StringVar(&gtFl, "gt", "", "filter tags 'greater than' this value - allowed values are based on the 'field' choosen. A relative time in seconds (s), minutes (m), hours (h), days (d), or weeks (w) is supported.")
	tagsCmd.PersistentFlags().StringVar(&ltFl, "lt", "", "filter tags 'less than' this value - allowed values are based on the 'field' choosen. A relative time in seconds (s), minutes (m), hours (h), days (d), or weeks (w) is supported.")

	rootCmd.AddCommand(tagsCmd)
}

// This is a wrapper for time.ParseDuration that adds support for days and weeks
func parseDuration(duration string) (time.Duration, error) {

	if duration == "" {
		return time.ParseDuration(duration)
	}

	cutPoint := len(duration) - 1
	var multiple int

	if duration[cutPoint:] == "d" {
		multiple = 24 // number of hours in a day
	} else if duration[cutPoint:] == "w" {
		multiple = 24 * 7 // number of hours in a week
	}

	integer, err := strconv.Atoi(duration[:cutPoint])
	if err != nil {
		return 0, err
	}

	if multiple != 0 {
		return time.ParseDuration(strconv.Itoa(integer*multiple) + "h")
	}

	return time.ParseDuration(duration)
}
