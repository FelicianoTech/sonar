package cmd

import (
	"fmt"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	labelsGetCmd = &cobra.Command{
		Use:   "get <image-name> <key>",
		Short: "Display the value of a single label for a given Docker image name",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			label, err := docker.GetLabel(args[0], args[1])
			if err != nil {
				log.Fatal("Failed to get label.")
			}

			fmt.Printf("%s: %s\n", args[1], label)
		},
	}
)

func init() {
	labelsCmd.AddCommand(labelsGetCmd)
}
