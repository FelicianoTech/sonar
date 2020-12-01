package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var packagesDiffCmd = &cobra.Command{
	Use:   "diff <image> <image>",
	Short: "Displays the difference in installed packages between two Docker images",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 2 {
			log.Fatal("Please provide exactly two Docker image names.")
			os.Exit(1)
		}

		images, err := getImageRefs(args)
		if err != nil {
			fmt.Errorf("%s, err")
		}

		packages1 := listPackages(images[0])
		packages2 := listPackages(images[1])

		diffPackages := pkgMissingFromA(packages2, packages1)
		fmt.Printf("Packages only in %s\n", images[0].String())
		fmt.Println("==============================")
		for _, pkg := range diffPackages {
			fmt.Printf("%s\t\t%s\t%s\n", pkg.Name, pkg.Version, pkg.Manager)
		}

		fmt.Printf("\n")

		diffPackages = pkgMissingFromA(packages1, packages2)
		fmt.Printf("Packages only in %s\n", images[1].String())
		fmt.Println("==============================")
		for _, pkg := range diffPackages {
			fmt.Printf("%s\t\t%s\t%s\n", pkg.Name, pkg.Version, pkg.Manager)
		}
	},
}

func init() {
	packagesCmd.AddCommand(packagesDiffCmd)
}

func pkgMissingFromA(listA, listB []packageInfo) []packageInfo {

	listAMap := make(map[string]packageInfo, len(listA))

	bOnlyPkgs := []packageInfo{}

	for _, v := range listA {
		listAMap[v.Name] = v
	}

	for _, v := range listB {

		if _, found := listAMap[v.Name]; !found {
			bOnlyPkgs = append(bOnlyPkgs, v)
		}
	}

	return bOnlyPkgs
}
