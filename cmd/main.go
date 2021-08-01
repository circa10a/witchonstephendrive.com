package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var frontendAssetsFS embed.FS
var apiDocAssetsFS embed.FS

var rootCmd = &cobra.Command{
	Use:   "witch",
	Short: "CLI for halloween decoration automation",
	Long:  "View docs at https://github.com/circa10a/witchonstephendrive.com",
}

func Execute(frontendAssets, apiDocAssets embed.FS) {
	// Go only allows reading embed files from the root of the project
	// So we just pass them here and set them as global vars in the cmd package
	frontendAssetsFS = frontendAssets
	apiDocAssetsFS = apiDocAssets
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
