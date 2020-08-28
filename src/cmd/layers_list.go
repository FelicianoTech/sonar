package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	//sumSizeFl bool

	layersListCmd = &cobra.Command{
		Use:   "list <image-name> <tag>",
		Short: "Displays the layers for a given Docker image",
		Long: `The output of instruction is limited to 55 characters.
`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				log.Fatal("Please provide a Docker image.")
				os.Exit(1)
			}

			dockerLayers := getAllLayers(args[0], args[1])

			fmt.Println("The layers are: ")

			var counter int

			for _, layer := range dockerLayers {

				var digestStr string
				var sizeStr string

				counter++

				if layer.digest == "" {
					digestStr = "<not-assigned>\t\t\t\t\t\t\t"
				} else {
					digestStr = layer.digest
				}

				if layer.size == 0 {
					sizeStr = "0 B\t"
				} else {
					sizeStr = ByteCountBinary(layer.size)
				}

				fmt.Printf("%d:\t%s\t%s\t%.55s\n", counter, digestStr, sizeStr, layer.instruction)
			}

			fmt.Printf("Total layers: %d\n", len(dockerLayers))
		},
	}
)

func init() {
	layersCmd.AddCommand(layersListCmd)
}
