package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var imagesListCmd = &cobra.Command{
	Use:   "list <namespace>",
	Short: "Displays a list of images for a given Docker Hub namespace",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			log.Fatal("Please provide a Docker Hub namespace.")
			os.Exit(1)
		}

		req, err := http.NewRequest("GET", "https://hub.docker.com/v2/repositories/"+args[0]+"/", nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := sendRequest(req, "", "")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var result2 map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result2)

		fmt.Println("The images are: ")

		for _, v := range result2["results"].([]interface{}) {
			//log.Println(v)
			//var images []interface{}
			//images = v.(map[string]interface{})["images"].([]interface{})
			fmt.Println(v.(map[string]interface{})["name"])
			//log.Println("digest: ", images[0].(map[string]interface{})["digest"])
		}
	},
}

func init() {
	imagesCmd.AddCommand(imagesListCmd)
}
