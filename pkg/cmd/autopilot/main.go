package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autopilot",
	Short: "AutoPilot is a lightweight runbook executor.",
	Long: `AutoPilot allows you to execute runbooks with minimal setup and maximum flexibility.
Inspired by Do Nothing Scripting.`,
}

func main() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Configure default logger
	// TODO: make this configurable
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}
