package main

import (
	"donot/pkg/core"
	"donot/pkg/executor"
	"donot/pkg/runbook"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [runbook-file]",
	Short: "Execute a runbook",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse the runbook file
		runbookFile := args[0]
		runbookMd := runbook.NewMarkdown()
		_ = runbookMd.Parse(runbookFile)

		// Create a new run
		run := core.NewRun("run-" + runbookMd.Name())

		// Set up the executor with a CLI observer
		executor := executor.NewExecutor(run, runbookMd)

		// Execute the runbook
		if err := executor.Execute(); err != nil {
			fmt.Printf("Run failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Run completed successfully.")
	},
}
