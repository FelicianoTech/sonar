package cmd

import (
	"fmt"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"
)

var getPullsCmd = &cobra.Command{
	Use:   "pulls <image> [<image>...]",
	Short: "Get the number of pulls for one or more images on Docker Hub",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		images, err := getImageRefs(args)
		if err != nil {
			return err
		}

		for _, image := range images {

			image.ShowTag = false

			pulls, err := docker.ImagePulls(image.String())
			if err != nil {
				fmt.Errorf("Error retrieving pulls: %s", err)
			}

			fmt.Printf("The number of %v pulls is: %v\n", image, pulls)
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getPullsCmd)
}
