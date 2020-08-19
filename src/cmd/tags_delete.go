package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var (
	dryRunFl bool
	yesFl    bool

	tagsDeleteCmd = &cobra.Command{
		Use:   "delete <image-name>",
		Short: "Deletes one or more tags based on a parameter",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				log.Fatal("Please provide a Docker image name.")
				os.Exit(1)
			}

			if fieldFl != "date" {
				log.Fatal("'field' is a required field and must be 'date'.")
				os.Exit(1)
			}

			if gtFl == "" && ltFl == "" {
				log.Fatal("Either the 'gt' or 'lt' flags need to be set.")
				os.Exit(1)
			}

			gDuration, err := parseDuration(gtFl)
			if err != nil {
				fmt.Errorf("Cannot parse duration from 'gt'", err)
			}
			gCutDate := time.Now().Add(-gDuration)

			lDuration, err := parseDuration(ltFl)
			if err != nil {
				fmt.Errorf("Cannot parse duration from 'lt'", err)
			}
			lCutDate := time.Now().Add(-lDuration)

			dockerTags := getAllTags(args[0])
			var tagsToDelete []dockerTag

			for _, tag := range dockerTags {

				if gtFl != "" && fieldFl == "date" {
					if gCutDate.After(tag.date) {
						tagsToDelete = append(tagsToDelete, tag)
					}
				}

				if ltFl != "" && fieldFl == "date" {
					if lCutDate.Before(tag.date) {
						tagsToDelete = append(tagsToDelete, tag)
					}
				}
			}

			if len(tagsToDelete) == 0 {

				fmt.Println("There were no tags to delete.")
				return
			}

			if dryRunFl {

				fmt.Println("The tags that would have been deleted: ")
				for _, tag := range tagsToDelete {
					fmt.Println(tag.name)
				}

				return
			}

			if !yesFl {
				fmt.Printf("You are about to permanently delete %d tags. Continue? [y/yes/n/no] ", len(tagsToDelete))
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					if scanner.Text() == "n" || scanner.Text() == "no" {
						fmt.Println("Cancelling.")
						return
					} else if scanner.Text() == "y" || scanner.Text() == "yes" {
						break
					} else {
						fmt.Println("Invalid input. Try again.")
					}
				}
			} else {
				fmt.Printf("Permanently deleting %d tags.", len(tagsToDelete))
			}

			for _, tag := range tagsToDelete {
				err := deleteDockerTag(args[0], tag.name)
				if err != nil {
					log.Error(err)
				}
			}
		},
	}
)

func init() {

	tagsDeleteCmd.Flags().BoolVar(&dryRunFl, "dry-run", false, "show what would be deleted without actually deleting any tags")
	tagsDeleteCmd.Flags().BoolVarP(&yesFl, "yes", "y", false, "automatic yes to deletion prompt, useful for scripting")

	tagsCmd.AddCommand(tagsDeleteCmd)
}

func deleteDockerTag(image, tag string) error {

	req, err := http.NewRequest("DELETE", "https://hub.docker.com/v2/repositories/"+image+"/tags/"+tag+"/", nil)
	if err != nil {
		return errors.New("Failed to create DELETE request.")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if len(viper.Get("user").(string)) == 0 || len(viper.Get("pass").(string)) == 0 {
		return errors.New("This command requires Docker Hub credentials to be set in your environment.")
	}

	resp, err := sendRequest(req, viper.Get("user").(string), viper.Get("pass").(string))
	if err != nil {
		return (err)
	}
	defer resp.Body.Close()

	code := resp.StatusCode
	if err != nil {
		return err
	}

	if code != 204 {
		return errors.New("There was an error in deleting the tag: " + tag)
	}

	return nil
}
