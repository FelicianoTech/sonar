package cmd

import (
	"fmt"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	labelsListCmd = &cobra.Command{
		Use:   "list <image-name>",
		Short: "Displays all labels for a given Docker image name + tag",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			labels, err := docker.GetLabels(args[0])
			if err != nil {
				log.Fatal("Failed to get labels.")
			}

			for k, v := range labels {
				fmt.Printf("%s: %s\n", k, v)
			}
		},
	}
)

func init() {
	labelsCmd.AddCommand(labelsListCmd)
}
