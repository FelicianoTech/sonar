package cmd

import (
	"encoding/json"
	"net/http"
	"os"

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

		p := message.NewPrinter(message.MatchLanguage("en"))
		p.Printf("The number of %v pulls is: %v", args[0], result2["star_count"])

	},
}

func init() {
	getCmd.AddCommand(getStarsCmd)
}
