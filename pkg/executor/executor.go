package executor

import (
	"donot/pkg/core"
	"donot/pkg/step"
	"fmt"
	"time"
)

// Executor is responsible for running the steps in a runbook.
type Executor struct {
	Run     *core.Run
	Runbook step.Runbook
}

// NewExecutor creates a new Executor instance.
func NewExecutor(run *core.Run, runbook step.Runbook) *Executor {
	return &Executor{
		Run:     run,
		Runbook: runbook,
	}
}

// Execute runs all steps in the runbook sequentially.
func (e *Executor) Execute() error {
	e.Run.Status = core.StatusInProgress
	e.Run.StartTime = time.Now()

	for e.Run.CurrentStepIndex < len(e.Runbook.Steps()) {
		step := e.Runbook.Steps()[e.Run.CurrentStepIndex]

		// Execute the step.
		err := step.Run(e.Run)
		if err != nil {
			e.Run.Status = core.StatusAborted
			e.Run.EndTime = time.Now()
			return fmt.Errorf("error in step %s: %w", step.ID(), err)
		}

		// Mark step as complete, notify observers.
		e.Run.Log(step.ID(), "Step completed successfully.")

		// Advance to the next step.
		e.Run.CurrentStepIndex++
	}

	// Mark the run as completed.
	e.Run.Status = core.StatusCompleted
	e.Run.EndTime = time.Now()

	return nil
}
