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
			e.Run.Status = core.StatusFailed
			e.Run.EndTime = time.Now()
			return fmt.Errorf("error waiting for user input: %w", err)
		}

		switch option {
		case "y", "c", "":
			// Execute the step and continue to the next step
			err := s.Run(e.Run)
			if err != nil {
				e.Run.Status = core.StatusFailed
				e.Run.EndTime = time.Now()
				e.Run.Log(s.ID(), fmt.Sprintf("Error executing step: %s", err))
				return fmt.Errorf("error in step %s: %w", s.ID(), err)
			}
			// Mark step as complete
			e.Run.Log(s.ID(), "Step completed successfully.")
		case "n", "s":
			// Skip the step without executing it
			e.Run.Log(s.ID(), "Step skipped.")
		case "b":
			// Go back to the previous step without executing current step
			if e.Run.CurrentStepIndex > 0 {
				e.Run.CurrentStepIndex -= 2
			} else {
				e.Run.CurrentStepIndex--
			}
		case "q":
			// Quit the runbook.
			e.Run.Status = core.StatusAborted
			e.Run.EndTime = time.Now()
			e.Run.Log("", "Run aborted.")
			return nil
		}

		// Advance to the next step.
		e.Run.CurrentStepIndex++
	}

	// Mark the run as completed.
	e.Run.Status = core.StatusCompleted
	e.Run.EndTime = time.Now()
	e.Run.Log("", "Run completed successfully.")

	return nil
}
