package core

import (
	"fmt"
	"time"
)

// RunStatus represents the current status of the run.
type RunStatus string

const (
	StatusInProgress RunStatus = "in_progress"
	StatusCompleted  RunStatus = "completed"
	StatusAborted    RunStatus = "aborted"
	StatusFailed     RunStatus = "failed"
	StatusPaused     RunStatus = "paused"
)

// RunLog is used to track logs or output for each step execution.
type RunLog struct {
	StepID    string    // ID of the step associated with the log
	Message   string    // Log message
	Timestamp time.Time // Time when the log was recorded
}

// Run represents the current execution state of a runbook.
type Run struct {
	ID               string    // Unique identifier for the run
	CurrentStepIndex int       // Index of the step currently being executed
	Status           RunStatus // Current status of the run
	Logs             []RunLog  // Collection of logs for the run
	StartTime        time.Time // Start time of the run
	EndTime          time.Time // End time of the run
	Context          *Context  // Execution context for storing variables
	verbose          bool      // Verbose mode for additional output
}

// NewRun initializes a new run instance for a runbook.
func NewRun(id string, verbose bool) *Run {
	return &Run{
		ID:               id,
		CurrentStepIndex: 0,
		Status:           StatusInProgress,
		Logs:             []RunLog{},
		Context:          NewContext(),
		StartTime:        time.Now(),
		verbose:          verbose,
	}
}

// Log adds a new log message to the run.
func (r *Run) Log(stepID, message string) {
	log := RunLog{
		StepID:    stepID,
		Message:   message,
		Timestamp: time.Now(),
	}
	r.Logs = append(r.Logs, log)

	// Optionally print logs to stdout for immediate feedback
	if r.verbose {
		fmt.Printf("[%s] %s: %s\n", log.Timestamp.Format(time.RFC3339), stepID, message)
	}
}

// MarkStepComplete updates the run state after a step is successfully executed.
func (r *Run) MarkStepComplete() {
	r.CurrentStepIndex++
	if r.CurrentStepIndex >= len(r.Logs) {
		r.Status = StatusCompleted
		r.EndTime = time.Now()
	}
}

// Abort marks the run as aborted and records the reason.
func (r *Run) Abort(reason string) {
	r.Status = StatusAborted
	r.EndTime = time.Now()
	r.Log("system", fmt.Sprintf("Run aborted: %s", reason))
}

// Pause marks the run as paused.
func (r *Run) Pause() {
	r.Status = StatusPaused
	r.Log("system", "Run paused by user.")
}

// Resume resumes a paused run.
func (r *Run) Resume() {
	if r.Status == StatusPaused {
		r.Status = StatusInProgress
		r.Log("system", "Run resumed by user.")
	}
}

// IsComplete checks if the run has completed all steps.
func (r *Run) IsComplete() bool {
	return r.Status == StatusCompleted
}

// CurrentStepID returns the ID of the step currently being executed.
func (r *Run) CurrentStepID() string {
	if r.CurrentStepIndex < len(r.Logs) {
		return r.Logs[r.CurrentStepIndex].StepID
	}
	return ""
}
