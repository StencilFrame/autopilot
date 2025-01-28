package step

import "autopilot/pkg/core"

// UIType represents the type of user interface supported by a step.
type UIType string

type RenderFunc func(step Step) string

const (
	// UITypeCLI represents a command-line interface.
	UITypeCLI UIType = "CLI"
	// UITypeWeb represents a web interface.
	UITypeWeb UIType = "Web"
)

// Step defines the interface that all step types must implement.
type Step interface {
	ID() string                // Unique identifier for the step.
	Name() string              // Human-readable name of the step.
	Run(run *core.Run) error   // Executes the step logic.
	Render(ui UIType) string   // Renders the step for a specific UI type, but does not execute it.
	SupportsUI(ui UIType) bool // Indicates if the step supports a specific UI type.
}

// Runbook defines the structure of a workflow with multiple steps.
type Runbook interface {
	Name() string  // Name of the runbook if any
	Steps() []Step // Steps in the runbook
}

// BaseStep is a reusable struct for common step fields.
type BaseStep struct {
	IDValue   string
	NameValue string
}

// ID returns the step's unique identifier.
func (b *BaseStep) ID() string {
	return b.IDValue
}

// Name returns the human-readable name of the step.
func (b *BaseStep) Name() string {
	return b.NameValue
}
