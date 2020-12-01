package cmd

import (
	"fmt"
	"os"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var layersListCmd = &cobra.Command{
	Use:   "list <image>",
	Short: "Displays the layers for a given Docker image",
	Long: `The output of instruction is limited to 55 characters.
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			log.Fatal("Please provide a Docker image.")
			os.Exit(1)
		}

		image, err := docker.ParseImageRef(args[0])
		if err != nil {
			fmt.Errorf("%s, err")
		}

		image.ShowTag = false

		dockerLayers, err := docker.GetAllLayers(image.String(), image.Tag)
		if err != nil {
			log.Errorf("Failed getting layers for Docker tag: %s", err)
		}

		fmt.Println("The layers are: ")

		var counter int

		for _, layer := range dockerLayers {

			var digestStr string
			var sizeStr string

			counter++

			if layer.Digest == "" {
				digestStr = "<not-assigned>\t\t\t\t\t\t\t"
			} else {
				digestStr = layer.Digest
			}

			if layer.Size == 0 {
				sizeStr = "0 B\t"
			} else {
				sizeStr = ByteCountBinary(layer.Size)
			}

			fmt.Printf("%d:\t%s\t%s\t%.55s\n", counter, digestStr, sizeStr, layer.Instruction)
		}

		fmt.Printf("Total layers: %d\n", len(dockerLayers))
	},
}

func init() {
	layersCmd.AddCommand(layersListCmd)
}
