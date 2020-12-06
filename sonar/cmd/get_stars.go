package cmd

import (
	"fmt"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"
)

var getStarsCmd = &cobra.Command{
	Use:   "stars <image> [<image>...]",
	Short: "Get the number of stars for one or more images on Docker Hub",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		images, err := getImageRefs(args)
		if err != nil {
			fmt.Errorf("%s, err")
		}

		for _, image := range images {

			image.ShowTag = false

			stars, err := docker.ImageStars(image.String())
			if err != nil {
				fmt.Errorf("Error retrieving stars: %s", err)
			}

			fmt.Printf("The number of %v stars is: %v\n", image.String(), stars)
		}

	},
}

func init() {
	getCmd.AddCommand(getStarsCmd)
}
