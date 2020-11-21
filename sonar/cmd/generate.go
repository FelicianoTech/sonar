package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/felicianotech/sonar/sonar/metafile"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")

		mf := metafile.Metafile{
			Version: "0.1.0",
		}

		file, _ := json.MarshalIndent(mf, "", "")
		_ = ioutil.WriteFile("test.json", file, 0644)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
