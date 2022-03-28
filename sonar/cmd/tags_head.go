package cmd

import (
	"fmt"

	"github.com/felicianotech/sonar/sonar/docker"
	"github.com/spf13/cobra"

	semver "github.com/hashicorp/go-version"
)

var (
	filterNameFl string
	methodFl     string

	headCmd = &cobra.Command{
		Use:   "head <image-name>",
		Short: "Returns the head (sequentially last) tag of the image",
		Long: `Returns what is considered to be the most recent or latest tag for 
an image, based on a criteria. By default, this criteria (or method) is date 
but SemVer is also supported.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			dockerTags, err := docker.GetFilteredTags(args[0], filterNameFl)
			if err != nil {
				return fmt.Errorf("Failed retrieving Docker tags: %s", err)
			} else if len(dockerTags) == 0 {
				return fmt.Errorf("The image %s doesn't have any tags.", args[0])
			}

			// start off with the first tag before we start comparing so that
			// we're not comparing nil
			headTag := dockerTags[0]

			for _, tag := range dockerTags {

				if methodFl == "date" {

					if tag.Date.After(headTag.Date) {
						headTag = tag
					}
				} else if methodFl == "semver" {

					headTagV, err := semver.NewSemver(headTag.Name)
					if err != nil {
						// headTag isn't a semver tag so let's reset to the current
						// tag in the loop and start over
						headTag = tag
						continue
					}

					tagV, err := semver.NewSemver(tag.Name)
					if err != nil {
						// the tag to compare isn't a semver tag so let's skip
						continue
					}

					if headTagV.LessThan(tagV) {
						headTag = tag
					}
				}
			}

			if methodFl == "semver" {

				_, err := semver.NewSemver(headTag.Name)
				if err != nil {
					return fmt.Errorf("The image %s doesn't contain at least 1 valid semver tag.", args[0])
				}
			}

			fmt.Println(headTag.Name)

			return nil
		},
	}
)

func init() {

	headCmd.PersistentFlags().StringVarP(&methodFl, "method", "m", "date", "Criteria to calculate the head tag. Supported values are 'date' (default) or 'semver'.")
	headCmd.PersistentFlags().StringVar(&filterNameFl, "filter-name", ".*", "a regex of which tag names to include")
	tagsCmd.AddCommand(headCmd)
}
