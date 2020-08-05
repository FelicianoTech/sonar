package cmd

import (
	"fmt"
	"os"

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

			dockerTags := getAllTags(args[0])

			fmt.Println("The tags are: ")
			var totalSize uint64

			for _, tag := range dockerTags {
				fmt.Println(tag.name)
				if sumSizeFl {
					totalSize += uint64(tag.size)
				}
			}

			fmt.Printf("Total tags: %d\n", len(dockerTags))
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
