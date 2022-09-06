package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"
)

var (
	sumSizeFl bool
	archFl    bool

	tagsListCmd = &cobra.Command{
		Use:   "list <image-name>",
		Short: "Displays tags on Docker Hub for a given Docker image name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			var gCutDate time.Time
			var lCutDate time.Time

			if fieldFl == "date" {

				if gtFl != "" {

					gDuration, err := parseDuration(gtFl)
					if err != nil {
						return fmt.Errorf("Cannot parse duration from 'gt': %s", err)
					}
					gCutDate = time.Now().Add(-gDuration)
				}

				if ltFl != "" {

					lDuration, err := parseDuration(ltFl)
					if err != nil {
						return fmt.Errorf("Cannot parse duration from 'lt': %s", err)
					}
					lCutDate = time.Now().Add(-lDuration)
				}
			}

			dockerTags, err := docker.GetAllTags(args[0])
			if err != nil {
				return fmt.Errorf("Failed retrieving Docker tags: %s", err)
			}
			var filteredTags []docker.Tag

			for _, tag := range dockerTags {

				if fieldFl == "date" {
					if gtFl != "" && fieldFl == "date" {
						if gCutDate.After(tag.Date) {
							filteredTags = append(filteredTags, tag)
						}
					}

					if ltFl != "" && fieldFl == "date" {
						if lCutDate.Before(tag.Date) {
							filteredTags = append(filteredTags, tag)
						}
					}
				} else {
					filteredTags = append(filteredTags, tag)
				}
			}

			if len(filteredTags) == 0 {

				fmt.Println("There were no tags to list.")
				return nil
			}

			var totalSize uint64

			// This is where we actually output tags to display
			for _, tag := range filteredTags {

				fmt.Printf("%s", tag.Name)

				if sumSizeFl {
					totalSize += uint64(tag.Size)
				}

				if archFl {

					var arches []string

					for _, img := range tag.Images {
						arches = append(arches, img.Arch)
					}

					fmt.Printf(" (%s)", strings.Join(arches, ","))
				}

				fmt.Printf("\n")
			}

			fmt.Println("====================")
			fmt.Printf("Tags: showing %d of %d\n", len(filteredTags), len(dockerTags))
			if sumSizeFl {
				fmt.Printf("Total size: %s\n", ByteCountBinary(totalSize))
			}

			return nil
		},
	}
)

func init() {
	tagsListCmd.Flags().BoolVar(&sumSizeFl, "sum-size", false, "output the storage size of tags")
	tagsCmd.PersistentFlags().BoolVar(&archFl, "arch", false, "show CPU architectures available per tag. Defaults to false.")

	tagsCmd.AddCommand(tagsListCmd)
}
