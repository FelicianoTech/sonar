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

	"github.com/felicianotech/sonar/sonar/docker"

	"github.com/spf13/cobra"

	docker2 "github.com/fsouza/go-dockerclient"
	log "github.com/sirupsen/logrus"
)

var (
	formatFl string

	packagesListCmd = &cobra.Command{
		Use:   "list <image>",
		Short: "Displays installed packages for a given Docker image",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				log.Fatal("Please provide a Docker image name.")
				os.Exit(1)
			}

			if formatFl != "terminal" && formatFl != "json" {

				log.Fatal("Error: the format flag can only be 'terminal' or 'json'.")
			}

			images, err := getImageRefs(args)
			if err != nil {
				log.Fatalf("%s", err)
			}

			packages := listPackages(images[0])

			if formatFl == "json" {

				jsonData, err := json.Marshal(packages)
				if err != nil {
					log.Fatalf("%s", err)
				}

				fmt.Printf("%s", jsonData)
			} else {

				for _, pkg := range packages {
					fmt.Printf("%s\t\t%s\t%s\n", pkg.Name, pkg.Version, pkg.Manager)
				}
			}
		},
	}
)

func init() {

	packagesListCmd.Flags().StringVar(&formatFl, "format", "terminal", "the format in which to show results")

	packagesCmd.AddCommand(packagesListCmd)
}

func listPackages(imgRef *docker.ImageRef) []packageInfo {

	image := imgRef.String()

	client, err := docker2.NewClientFromEnv()
	if err != nil {
		log.Error("Error: Failed to create new Docker client.")
		log.Fatal(err)
	}

	if _, err := client.InspectImage(image); err != nil {
		if err == docker2.ErrNoSuchImage {
			if err := client.PullImage(docker2.PullImageOptions{
				Repository: strings.Split(image, ":")[0],
				Tag:        strings.Split(image, ":")[1],
			}, docker2.AuthConfiguration{}); err != nil {
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

	// Conbine types into a string
	types := strings.Join(typeFl, ",")

	container, err := client.CreateContainer(docker2.CreateContainerOptions{
		Name: "sonar-compile",
		Config: &docker2.Config{
			Image: image,
			Cmd:   []string{"/sonar", "packages", "compile", "--type=" + types},
		},
		HostConfig: &docker2.HostConfig{
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
	if err := client.DownloadFromContainer(container.ID, docker2.DownloadFromContainerOptions{
		OutputStream: &buf2,
		Path:         "/tmp/sonar-packages.json",
	}); err != nil {
		log.Error("Error: Failed to retrieve file from container.")
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

	if err := client.RemoveContainer(docker2.RemoveContainerOptions{
		ID:    container.ID,
		Force: true,
	}); err != nil {
		log.Error("Error: Failed to remove container.")
		log.Fatal(err)
	}

	var packages []packageInfo
	err = json.Unmarshal(jsonData, &packages)
	if err != nil {
		log.Fatal(err)
	}

	return packages
}
