package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gopherlibs/pmm/pmm"
	"github.com/spf13/cobra"
)

var (
	outputFl string

	packagesCompileCmd = &cobra.Command{
		Use:    "compile",
		Short:  "Displays installed packages for the system",
		Hidden: true,
		Args:   cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			var packages []pmm.PkgInfo

			if typeRequested(typeFl, "apk") {

				apk, err := pmm.New(pmm.TypeAPK)
				if err != nil {
					fmt.Errorf("Error: APK is not available - %s", err)
				} else {
					packages = append(packages, apk.List()...)
				}
			}

			if typeRequested(typeFl, "apt") {

				apt, err := pmm.New(pmm.TypeAPT)
				if err != nil {
					fmt.Errorf("Error: APT is not available - %s", err)
				} else {
					packages = append(packages, apt.List()...)
				}
			}

			if typeRequested(typeFl, "pip") {

				pip, err := pmm.New(pmm.TypePIP)
				if err != nil {
					fmt.Errorf("Error: PIP is not available - %s", err)
				} else {
					packages = append(packages, pip.List()...)
				}
			}

			if typeRequested(typeFl, "rpm") {

				rpm, err := pmm.New(pmm.TypeRPM)
				if err != nil {
					fmt.Errorf("Error: RPM is not available - %s", err)
				} else {
					packages = append(packages, rpm.List()...)
				}
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
