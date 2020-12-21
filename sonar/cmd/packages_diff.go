package cmd

import (
	"fmt"

	"github.com/gopherlibs/pmm/pmm"
	"github.com/spf13/cobra"
)

var packagesDiffCmd = &cobra.Command{
	Use:   "diff <image> <image>",
	Short: "Displays the difference in installed packages between two Docker images",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		images, err := getImageRefs(args)
		if err != nil {
			return fmt.Errorf("%s", err)
		}

		packages1 := listPackages(images[0])
		packages2 := listPackages(images[1])

		diffPackages := pmm.PkgMissingFromA(packages2, packages1)
		fmt.Printf("Packages only in %s\n", images[0].String())
		fmt.Println("==============================")
		for _, pkg := range diffPackages {
			fmt.Printf("%s\t\t%s\t%s\n", pkg.Name, pkg.Version, pkg.Manager)
		}

		fmt.Printf("\n")

		diffPackages = pmm.PkgMissingFromA(packages1, packages2)
		fmt.Printf("Packages only in %s\n", images[1].String())
		fmt.Println("==============================")
		for _, pkg := range diffPackages {
			fmt.Printf("%s\t\t%s\t%s\n", pkg.Name, pkg.Version, pkg.Manager)
		}

		return nil
	},
}

func init() {
	packagesCmd.AddCommand(packagesDiffCmd)
}
