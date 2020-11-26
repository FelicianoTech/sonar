package cmd

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	docker "github.com/fsouza/go-dockerclient"
	log "github.com/sirupsen/logrus"
)

var (
	formatFl string

	packagesListCmd = &cobra.Command{
		Use:   "list <image-name>",
		Short: "Displays installed packages for a given Docker image",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				log.Fatal("Please provide a Docker image name.")
				os.Exit(1)
			}

			if formatFl != "terminal" && formatFl != "json" {

				log.Fatal("Error: the format flag can only be 'terminal' or 'json'.")
			}

			client, err := docker.NewClientFromEnv()
			if err != nil {
				log.Error("Error: Failed to create new Docker client.")
				log.Fatal(err)
			}

			if _, err := client.InspectImage(args[0]); err != nil {
				if err == docker.ErrNoSuchImage {
					if err := client.PullImage(docker.PullImageOptions{
						Repository: strings.Split(args[0], ":")[0],
						Tag:        strings.Split(args[0], ":")[1],
					}, docker.AuthConfiguration{}); err != nil {
						log.Error("Error: Failed to start container.")
						log.Fatal(err)
					}
				} else {
					log.Error("Error: Failed to inspect image.")
					log.Fatal(err)
				}
			}

			absPath, err := os.Executable()
			if err != nil {
				log.Error("Error: Can't properly locate sonar executable.")
				log.Fatal(err)
			}

			container, err := client.CreateContainer(docker.CreateContainerOptions{
				Name: "sonar-compile",
				Config: &docker.Config{
					Image: args[0],
					Cmd:   []string{"/sonar", "packages", "compile"},
				},
				HostConfig: &docker.HostConfig{
					Binds: []string{absPath + ":/sonar"},
				},
			})
			if err != nil {
				log.Error("Error: Failed to create container.")
				log.Fatal(err)
			}

			if err := client.StartContainer(container.ID, nil); err != nil {
				log.Error("Error: Failed to start container.")
				log.Fatal(err)
			}

			exitCode, err := client.WaitContainer(container.ID)
			if err != nil {
				log.Error("Error: Failed to wait for container.")
				log.Fatal(err)
			}
			if exitCode != 0 {
				log.Error("Error: Sonar failed to run within the container.")
				log.Fatal(err)
			}

			var buf2 bytes.Buffer
			if err := client.DownloadFromContainer(container.ID, docker.DownloadFromContainerOptions{
				OutputStream: &buf2,
				Path:         "/tmp/sonar-packages.json",
			}); err != nil {
				log.Error("Error: Failed")
				log.Fatal(err)
			}

			tr := tar.NewReader(&buf2)
			_, err = tr.Next()
			if err == io.EOF {
				log.Fatal("We never got the file.")
			}
			if err != nil {
				log.Fatal(err)
			}

			jsonData, err := ioutil.ReadAll(tr)
			if err != nil {
				log.Error("Error: Failed to read tar archive.")
				log.Fatal(err)
			}

			if formatFl == "json" {
				fmt.Printf("%s", jsonData)
			} else {

				var packages []string
				err = json.Unmarshal(jsonData, &packages)
				if err != nil {
					log.Fatal(err)
				}

				for _, pkg := range packages {
					fmt.Println(pkg)
				}
			}

			if err := client.RemoveContainer(docker.RemoveContainerOptions{
				ID:    container.ID,
				Force: true,
			}); err != nil {
				log.Error("Error: Failed to remove container.")
				log.Fatal(err)
			}
		},
	}
)

func init() {

	packagesListCmd.Flags().StringVar(&formatFl, "format", "terminal", "the format in which to show results")

	packagesCmd.AddCommand(packagesListCmd)
}
