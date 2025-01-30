package executor

import (
	"autopilot/pkg/core"
	"autopilot/pkg/step"
	"fmt"
	"time"
)

// LocalExecutor is responsible for running the steps in a runbook on the local machine.
type LocalExecutor struct {
	Run     *core.Run
	Runbook step.Runbook
	CLIMenu *CLIMenu
}

// NewLocalExecutor creates a new LocalExecutor instance.
func NewLocalExecutor(run *core.Run, runbook step.Runbook) *LocalExecutor {
	return &LocalExecutor{
		Run:     run,
		Runbook: runbook,
		CLIMenu: NewCLIMenu(),
	}
}

// Execute runs all steps in the runbook sequentially on the local machine.
func (e *LocalExecutor) Execute() error {
	e.Run.Status = core.StatusInProgress
	e.Run.StartTime = time.Now()
	e.Run.Log("", "Run started.")

	stepCount := len(e.Runbook.Steps())
	for e.Run.CurrentStepIndex < stepCount {
		s := e.Runbook.Steps()[e.Run.CurrentStepIndex]

		// Display the step to the user.
		fmt.Println("--------------------------------------------------")
		fmt.Printf("Step %d/%d ", e.Run.CurrentStepIndex+1, stepCount)
		fmt.Println(s.Render(step.UITypeCLI))

		// Wait for user input.
		option, err := e.CLIMenu.WaitForOption()
		if err != nil {
			e.Run.Log(s.ID(), fmt.Sprintf("Error waiting for user input: %s", err))
			return fmt.Errorf("error waiting for user input: %w", err)
		}

		if option == "q" {
			// Quit the runbook.
			e.Run.Status = core.StatusAborted
			e.Run.EndTime = time.Now()
			e.Run.Log("", "Run aborted.")
			return nil
		}

		switch option {
		case "y", "c":
			// Execute the step.
			err := s.Run(e.Run)
			if err != nil {
				e.Run.Status = core.StatusAborted
				e.Run.EndTime = time.Now()
				return fmt.Errorf("error in step %s: %w", s.ID(), err)
			}
		case "n", "s":
			// Skip the step.
			e.Run.Log(s.ID(), "Step skipped.")
		case "b":
			// Go back to the previous step.
			if e.Run.CurrentStepIndex > 0 {
				e.Run.CurrentStepIndex -= 2
			} else {
				e.Run.CurrentStepIndex--
			}
		}

		// Mark step as complete
		e.Run.Log(s.ID(), "Step completed successfully.")
		// Advance to the next step.
		e.Run.CurrentStepIndex++
	}

	// Mark the run as completed.
	e.Run.Status = core.StatusCompleted
	e.Run.EndTime = time.Now()
	e.Run.Log("", "Run completed successfully.")

	return nil
}
