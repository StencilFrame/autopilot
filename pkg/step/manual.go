package step

import (
	"autopilot/pkg/core"
	"bufio"
	"fmt"
	"os"
)

// ManualStep waits for user confirmation before proceeding.
type ManualStep struct {
	BaseStep
	Instructions string // Instructions for the user.
	uis          map[UIType]RenderFunc
}

// NewManualStep creates a new ManualStep instance.
func NewManualStep(id, name, instructions string) *ManualStep {
	m := &ManualStep{
		BaseStep: BaseStep{
			IDValue:   id,
			NameValue: name,
		},
		Instructions: instructions,
	}
	// Register render functions for supported UI types.
	m.uis = map[UIType]RenderFunc{
		UITypeCLI: m.renderCLI,
		UITypeWeb: m.renderWeb,
	}
	return m
}

// Run executes the manual step logic.
// TODO: Refactor to support different UI types.
func (m *ManualStep) Run(run *core.Run) error {
	// Display instructions to the user.
	fmt.Println("\n--------------------------------------------------")
	fmt.Printf("[Manual Step] %s\n\n", m.Name())
	if m.Instructions != "" {
		fmt.Println(m.Instructions)
	}
	fmt.Print("\nPress Enter when done...")

	// Wait for user confirmation.
	_, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to confirm manual step: %w", err)
	}

	run.Log(m.ID(), "Manual step completed by user.")
	return nil
}

// SupportsUI indicates if the manual step supports a specific UI type.
func (m *ManualStep) SupportsUI(ui UIType) bool {
	return m.uis[ui] != nil
}

// Render renders the manual step for a specific UI type.
func (m *ManualStep) Render(ui UIType) string {
	if f, ok := m.uis[ui]; ok {
		return f(m)
	}

	return fmt.Sprintf("Manual Step: %s\n%s", m.Name(), m.Instructions)
}

// renderCLI renders the manual step for the command-line interface.
func (m *ManualStep) renderCLI(step Step) string {
	s := fmt.Sprintf("[Manual Step] %s", m.Name())
	if m.Instructions != "" {
		s += fmt.Sprintf("\n%s", m.Instructions)
	}
	return s
}

// renderWeb renders the manual step for the web interface.
func (m *ManualStep) renderWeb(step Step) string {
	return fmt.Sprintf("<h2>Manual Step: %s</h2><p>%s</p>", m.Name(), m.Instructions)
}
