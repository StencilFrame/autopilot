package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionInfo is the version of AutoPilot set during build
// It is set using the -ldflags parameter in the build command: go build -ldflags "-X main.versionInfo=1.0"
var versionInfo string

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of AutoPilot",
	Run: func(cmd *cobra.Command, args []string) {
		vi := versionInfo
		if vi == "" {
			vi = "development"
		}
		fmt.Println("AutoPilot version:", vi)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
