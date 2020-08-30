package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var setReadmeCmd = &cobra.Command{
	Use:   "readme <image-name> <file>",
	Short: "Set the readme for an image on Docker Hub",
	Long: `This command was previously called 'set description'. Docker Hub now 
refers to the summary as a description and the old description as a readme.`,
	// This alias is temporarly. Docker Hub changed how they word things. The
	// summary is called a description and the readme is called
	// full_description. Since we previously called readme description, this
	// alias allows for a smooth transition over to the new names.
	Aliases: []string{"description"},
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			log.Fatal("Please provide a text file.")
			os.Exit(1)
		}

		content, err := ioutil.ReadFile(args[1])
		if err != nil {
			log.Fatal(err)
		}

		// Escape file content for use in JSON
		content = []byte(strconv.Quote(string(content)))

		content = append([]byte("{\"full_description\": "), content[:len(content)-1]...)
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
			log.Fatal("There was an error updating the description. Code " + resp.Status)
		} else {
			fmt.Println("Successfully updated.")
		}
	},
}

func init() {
	setCmd.AddCommand(setReadmeCmd)
}
