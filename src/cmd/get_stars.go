package cmd

import (
	"fmt"
	"os"

	"github.com/felicianotech/sonar/docker"

	"github.com/spf13/cobra"
	"golang.org/x/text/message"

	log "github.com/sirupsen/logrus"
)

var getStarsCmd = &cobra.Command{
	Use:   "stars <image-name>",
	Short: "Get the number of stars for an image on Docker Hub",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			log.Fatal("Please provide an image name.")
			os.Exit(1)
		}

		stars, err := docker.ImageStars(args[0])
		if err != nil {
			fmt.Errorf("Error retrieving stars: %s", err)
		}

		p := message.NewPrinter(message.MatchLanguage("en"))
		p.Printf("The number of %v pulls is: %v", args[0], stars)

	},
}

func init() {
	getCmd.AddCommand(getStarsCmd)
}
