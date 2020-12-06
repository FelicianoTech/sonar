package cmd

import (
	"fmt"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var imagesListCmd = &cobra.Command{
	Use:   "list <namespace>",
	Short: "Displays a list of images for a given Docker Hub namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		images, err := docker.ImageList(args[0])
		if err != nil {
			log.Errorf("Error getting image list: %s", err)
		}

		fmt.Println("The images are: ")

		for _, image := range images {
			fmt.Println(image)
		}
	},
}

func init() {
	imagesCmd.AddCommand(imagesListCmd)
}
