package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autopilot",
	Short: "AutoPilot is a lightweight runbook executor.",
	Long: `AutoPilot allows you to execute runbooks with minimal setup and maximum flexibility.
Inspired by Do Nothing Scripting.`,
}

func main() {
	// Add subcommands
	rootCmd.AddCommand(runCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
