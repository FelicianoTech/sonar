package cmd

import (
	"fmt"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"
)

var layersListCmd = &cobra.Command{
	Use:   "list <image>",
	Short: "Displays the layers for a given Docker image",
	Long:  `The output of instruction is limited to 55 characters.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		image, err := docker.ParseImageRef(args[0])
		if err != nil {
			return fmt.Errorf("%s", err)
		}

		image.ShowTag = false

		dockerLayers, err := docker.GetAllLayers(image.String(), image.Tag)
		if err != nil {
			return fmt.Errorf("Failed getting layers for Docker tag: %s", err)
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

		return nil
	},
}

func init() {
	layersCmd.AddCommand(layersListCmd)
}
