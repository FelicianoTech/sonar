package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/felicianotech/sonar/docker"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var setDescriptionCmd = &cobra.Command{
	Use:   "summary <image-name> <summary-string>",
	Short: "Set the summary for an image on Docker Hub",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			log.Fatal("Please provide a summary string.")
			os.Exit(1)
		}

		// Escape file content for use in JSON
		content := []byte(strconv.Quote(args[1]))

		content = append([]byte("{\"description\": "), content[:len(content)-1]...)
		content = append(content, []byte("\"}")...)

		req, err := http.NewRequest("PATCH", "https://hub.docker.com/v2/repositories/"+args[0]+"/", bytes.NewBuffer(content))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		if len(viper.Get("user").(string)) == 0 || len(viper.Get("pass").(string)) == 0 {
			log.Fatal("This command requires Docker Hub credentials to be set in your environment.")
		}

		resp, err := docker.SendRequest(req, viper.Get("user").(string), viper.Get("pass").(string))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if code, _ := strconv.Atoi(resp.Status); code >= 300 {
			log.Fatal("There was an error updating the summary. Code " + resp.Status)
		} else {
			fmt.Println("Successfully updated.")
		}
	},
}

func init() {
	setCmd.AddCommand(setDescriptionCmd)
}
