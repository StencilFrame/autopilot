package main

import (
	"donot/pkg/core"
	"donot/pkg/executor"
	coreRunbook "donot/pkg/runbook"
	"donot/pkg/step"
	"fmt"
	"os"
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
			fmt.Printf("Unsupported runbook type: %s\n", ext)
		}

		// Create a new run
		run := core.NewRun("run-" + runbook.Name())

		// Set up the executor with a CLI observer
		executor := executor.NewExecutor(run, runbook)

		// Execute the runbook
		if err := executor.Execute(); err != nil {
			fmt.Printf("Run failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Run completed successfully.")
	},
}
