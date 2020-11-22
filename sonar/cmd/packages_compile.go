package cmd

import (
	"encoding/json"
	"io/ioutil"

	"github.com/arduino/go-apt-client"
	"github.com/spf13/cobra"
)

var (
	packagesCompileCmd = &cobra.Command{
		Use:    "compile",
		Short:  "Displays installed packages for the system",
		Long:   "Currently only supports Apt packages",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {

			packages := listPackages()

			jsonFile, _ := json.MarshalIndent(packages, "", "")
			_ = ioutil.WriteFile("/sonar-packages.json", jsonFile, 0644)
		},
	}
)

func init() {
	packagesCmd.AddCommand(packagesCompileCmd)
}

func listPackages() []string {

	var packages []string

	allPackages, _ := apt.List()

	for _, pkg := range allPackages {

		if pkg.Status == "installed" {
			packages = append(packages, pkg.Name)
		}
	}

	return packages
}
