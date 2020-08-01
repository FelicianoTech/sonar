package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var tagsListCmd = &cobra.Command{
	Use:   "list <image-name>",
	Short: "Displays tags for a given Docker image name",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			log.Fatal("Please provide a Docker image name.")
			os.Exit(1)
		}

		reqURL := "https://hub.docker.com/v2/repositories/" + args[0] + "/tags/?page_size=100"
		var results []string

		for {

			req, err := http.NewRequest("GET", reqURL, nil)
			if err != nil {
				log.Fatal(err)
			}

			resp, err := sendRequest(req, "", "")
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			var respPage map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&respPage)

			for _, v := range respPage["results"].([]interface{}) {
				results = append(results, v.(map[string]interface{})["name"].(string))
			}

			if respPage["next"] == nil {
				break
			} else {
				reqURL = respPage["next"].(string)
			}
		}

		fmt.Println("The tags are: ")

		for _, tag := range results {
			fmt.Println(tag)
		}
	},
}

func init() {
	tagsCmd.AddCommand(tagsListCmd)
}
