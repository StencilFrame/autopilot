package main

import (
	"autopilot/pkg/core"
	"autopilot/pkg/executor"
	coreRunbook "autopilot/pkg/runbook"
	"autopilot/pkg/step"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [runbook-file]",
	Short: "Execute a runbook",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse the runbook file
		runbookFile := args[0]

		// Determine the runbook type
		ext := filepath.Ext(runbookFile)
		var runbook step.Runbook
		switch ext {
		case ".md":
			runbookMd := coreRunbook.NewMarkdown()
			_ = runbookMd.Parse(runbookFile)
			runbook = runbookMd
		case ".yml", ".yaml":
			runbookYaml := coreRunbook.NewYAML()
			_ = runbookYaml.Parse(runbookFile)
			runbook = runbookYaml
		default:
			log.Fatalf("Unsupported runbook type: %s", ext)
		}

		// Create a new run
		verbose, _ := cmd.Flags().GetBool("v")
		run := core.NewRun("run-"+runbook.Name(), verbose)

		// Set up the executor with a CLI observer
		executor := executor.NewLocalExecutor(run, runbook)

		// Execute the runbook
		if err := executor.Execute(); err != nil {
			log.Fatalf("Run failed: %v", err)
		}
	},
}

func init() {
	runCmd.Flags().BoolP("v", "v", false, "Enable verbose output")
	rootCmd.AddCommand(runCmd)
}
