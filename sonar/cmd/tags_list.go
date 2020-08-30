package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	sumSizeFl bool

	tagsListCmd = &cobra.Command{
		Use:   "list <image-name>",
		Short: "Displays tags for a given Docker image name",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				log.Fatal("Please provide a Docker image name.")
				os.Exit(1)
			}

			if fieldFl != "date" && fieldFl != "" {
				log.Fatal("if 'field' is set it must be 'date'.")
				os.Exit(1)
			}

			gDuration, err := parseDuration(gtFl)
			if err != nil {
				fmt.Errorf("Cannot parse duration from 'gt'", err)
			}
			gCutDate := time.Now().Add(-gDuration)

			lDuration, err := parseDuration(ltFl)
			if err != nil {
				fmt.Errorf("Cannot parse duration from 'lt'", err)
			}
			lCutDate := time.Now().Add(-lDuration)

			dockerTags, err := docker.GetAllTags(args[0])
			if err != nil {
				fmt.Errorf("Failed retrieving Docker tags", err)
			}
			var filteredTags []docker.Tag

			for _, tag := range dockerTags {

				if fieldFl != "" {
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
				return
			}

			var totalSize uint64

			for _, tag := range filteredTags {
				fmt.Println(tag.Name)
				if sumSizeFl {
					totalSize += uint64(tag.Size)
				}
			}

			fmt.Println("====================")
			fmt.Printf("Tags: showing %d of %d\n", len(filteredTags), len(dockerTags))
			if sumSizeFl {
				fmt.Printf("Total size: %s\n", ByteCountBinary(totalSize))
			}
		},
	}
)

func init() {
	tagsListCmd.Flags().BoolVar(&sumSizeFl, "sum-size", false, "output the storage size of tags")

	tagsCmd.AddCommand(tagsListCmd)
}
