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
	Short: "Generate metafile.json file",
	Long: `Generate a metafile.json file optionally using a seed.yml file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")

		mf := metafile.Generate()

		file, _ := json.MarshalIndent(mf, "", "")
		_ = ioutil.WriteFile("test.json", file, 0644)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
