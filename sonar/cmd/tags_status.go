package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/felicianotech/sonar/sonar/docker"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status <image-name>",
	Short: "Displays the push/pull status of local tags",
	Long:  `Displays the tags for a particular image that you have locally in a table. Provides info on if the tag also exists on Docker Hub, and if it does, is the local or Hub version newer.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		dCLI, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			return err
		}

		images, err := dCLI.ImageList(context.Background(), types.ImageListOptions{})
		if err != nil {
			return err
		}

		var localTags []docker.Tag

		// Loop through all images
		for _, image := range images {

			// Look through each image's tags
			for _, tag := range image.RepoTags {

				if !strings.HasPrefix(tag, args[0]) {
					break
				}

				// this next section is to get the digest since while the image
				// ID is available the digest is the preferred identifer.
				var digestOrID string
				if image.RepoDigests != nil {

					digestOrID = strings.Split(image.RepoDigests[0], "@")[1]
				} else {
					digestOrID = image.ID
				}

				localTags = append(localTags, docker.Tag{
					Name:   strings.Split(tag, ":")[1],
					Size:   image.Size,
					Date:   time.Unix(image.Created, 0).UTC(),
					Digest: digestOrID,
				})
			}
		}

		dCLI.Close()

		// output data

		if len(localTags) == 0 {
			fmt.Println("The image doesn't have any tags local.")
			return nil
		}

		hubTags, hubTagsErr := docker.GetAllTags(args[0])

		fmt.Printf("Local tags for %s:\n\n", args[0])
		fmt.Println(" Tag        Docker Hub Status")
		fmt.Println("========== ===================")
		for _, tag := range localTags {

			status := "local-only"

			if hubTagsErr == nil {

				for _, hubTag := range hubTags {
					if tag.Name == hubTag.Name {

						if hubTag.Digest == "" {
							status = "can't check"
						} else if tag.Digest == hubTag.Digest {
							status = "synced"
						} else if tag.Date.After(hubTag.Date) {
							status = "newer"
						} else if tag.Date.Equal(hubTag.Date) {
							status = "synced"
						} else if tag.Date.Before(hubTag.Date) {
							status = "older"
						}

						break
					}
				}
			}

			fmt.Printf(" %-10s %5s\n", tag.Name, status)
		}

		// compare digest and.... creation time?

		return nil
	},
}

func init() {
	tagsCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
