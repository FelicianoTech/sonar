package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/arduino/go-apt-client"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	outputFl string

	packagesCompileCmd = &cobra.Command{
		Use:    "compile",
		Short:  "Displays installed packages for the system",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {

			var packages []packageInfo

			if typeRequested(typeFl, "apt") {
				packages = append(packages, listPackagesAPT()...)
			}

			if typeRequested(typeFl, "pip") {
				packages = append(packages, listPackagesPIP()...)
			}

			if typeRequested(typeFl, "rpm") {
				packages = append(packages, listPackagesRPM()...)
			}

			jsonFile, _ := json.Marshal(packages)

			if outputFl != "stdout" {
				_ = ioutil.WriteFile(outputFl, jsonFile, 0644)
			} else {
				fmt.Println(string(jsonFile))
			}
		},
	}
)

func init() {

	packagesCompileCmd.Flags().StringVar(&outputFl, "output", "/tmp/sonar-packages.json", "where to store results, use 'stdout' or a filepath")
	packagesCmd.AddCommand(packagesCompileCmd)
}

func commandExists(command string) bool {

	_, err := exec.LookPath(command)

	return err == nil
}

func listPackagesAPT() []packageInfo {

	var packages []packageInfo

	allPackages, _ := apt.List()

	for _, pkg := range allPackages {

		if pkg.Status == "installed" {

			packages = append(packages, packageInfo{
				Name:    pkg.Name,
				Version: pkg.Version,
				Manager: "apt",
				Source:  "self",
			})
		}
	}

	return packages
}

func listPackagesPIP() []packageInfo {

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
				Source:  "self",
			})
		}
	}

	return packages
}

func listPackagesRPM() []packageInfo {

	var packages []packageInfo

	output, err := exec.Command("rpm", "-qa", "--qf", "%{NAME}\t%{VERSION}\n").Output()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(output), "\n")
	lines = lines[0 : len(lines)-1]

	for _, pkg := range lines {

		pkgSplit := strings.Split(pkg, "\t")
		packages = append(packages, packageInfo{
			Name:    pkgSplit[0],
			Version: pkgSplit[1],
			Manager: "rpm",
			Source:  "self",
		})
	}

	return packages
}
