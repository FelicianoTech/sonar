package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/arduino/go-apt-client"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	aptFl bool
	pipFl bool

	packagesCompileCmd = &cobra.Command{
		Use:    "compile",
		Short:  "Displays installed packages for the system",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {

			if pipFl {
				packages := listPackagesPip()

				for _, pkg := range packages {
					fmt.Printf("%s: %s\n", pkg.Name, pkg.Version)
				}

				jsonFile, _ := json.Marshal(packages)
				fmt.Printf("%s", string(jsonFile))
				_ = ioutil.WriteFile("/tmp/sonar-packages.json", jsonFile, 0644)
			} else {

				packages := listPackages()

				jsonFile, _ := json.Marshal(packages)
				fmt.Printf("%s", string(jsonFile))
				_ = ioutil.WriteFile("/tmp/sonar-packages.json", jsonFile, 0644)
			}
		},
	}
)

func init() {
	packagesCompileCmd.Flags().BoolVar(&aptFl, "apt", true, "show apt packages?")
	packagesCompileCmd.Flags().BoolVar(&pipFl, "pip", false, "show pip packages?")

	packagesCmd.AddCommand(packagesCompileCmd)
}

func listPackages() []packageInfo {

	var packages []packageInfo

	allPackages, _ := apt.List()

	for _, pkg := range allPackages {

		if pkg.Status == "installed" {

			packages = append(packages, packageInfo{
				Name:    pkg.Name,
				Version: pkg.Version,
				Manager: "apt",
			})
		}
	}

	return packages
}

func listPackagesPip() []packageInfo {

	var pipJSON []map[string]string
	var packages []packageInfo

	for _, pipCmd := range []string{"pip", "pip3"} {

		if !commandExists(pipCmd) {
			continue
		}

		output, err := exec.Command(pipCmd, "list", "--format=json").Output()
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(output, &pipJSON)
		if err != nil {
			log.Fatal(err)
		}

		for _, pkg := range pipJSON {
			packages = append(packages, packageInfo{
				Name:    pkg["name"],
				Version: pkg["version"],
				Manager: "pip",
			})
		}
	}

	return packages
}

func commandExists(command string) bool {

	_, err := exec.LookPath(command)

	return err == nil
}
