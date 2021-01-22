package cmd

import (
	"fmt"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"
)

var (
	tagsAliasCmd = &cobra.Command{
		Use:   "aliases <full-image-name>",
		Short: "Display tags that point to the same image the given tag does",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			img, err := docker.ParseImageRef(args[0])
			if err != nil {
				return fmt.Errorf("Failed parsing image name: %s", err)
			}

			dockerTags, err := docker.GetAllTags(img.Namespace + "/" + img.Name)

			var theDigest string

			for _, tag := range dockerTags {

				if tag.Name == img.Tag {
					theDigest = tag.Digest
					break
				}
			}

			fmt.Printf("Listing aliases for the tag %s: \n\n", img.Tag)

			for _, tag := range dockerTags {

				if tag.Digest == theDigest && tag.Name != img.Tag {
					fmt.Println(tag.Name)
				}
			}

			return nil
		},
	}
)

func init() {
	tagsCmd.AddCommand(tagsAliasCmd)
}
