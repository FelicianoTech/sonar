package cmd

import (
	"fmt"
	"os"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var getPullsCmd = &cobra.Command{
	Use:   "pulls <image> [<image>...]",
	Short: "Get the number of pulls for one or more images on Docker Hub",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			log.Fatal("Please provide an image name.")
			os.Exit(1)
		}

		images, err := getImageRefs(args)
		if err != nil {
			fmt.Errorf("%s, err")
		}

		for _, image := range images {

			image.ShowTag = false

			pulls, err := docker.ImagePulls(image.String())
			if err != nil {
				fmt.Errorf("Error retrieving pulls: %s", err)
			}

			fmt.Printf("The number of %v pulls is: %v\n", image, pulls)
		}
	},
}

func init() {
	getCmd.AddCommand(getPullsCmd)
}
