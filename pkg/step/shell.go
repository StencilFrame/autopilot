package step

import (
	"donot/pkg/core"
	"fmt"
	"os/exec"
	"strings"
)

// ShellStep runs a shell command.
type ShellStep struct {
	BaseStep
	Command string // Shell command to execute.
	uis     map[UIType]RenderFunc
}

// NewShellStep creates a new ShellStep instance.
func NewShellStep(id, name, command string) *ShellStep {
	s := &ShellStep{
		BaseStep: BaseStep{
			IDValue:   id,
			NameValue: name,
		},
		Command: command,
	}
	// Register render functions for supported UI types.
	s.uis = map[UIType]RenderFunc{
		UITypeCLI: s.renderCLI,
		UITypeWeb: s.renderWeb,
	}
	return s
}

// Run executes the shell step logic.
func (s *ShellStep) Run(run *core.Run) error {
	fmt.Println("\n--------------------------------------------------")
	fmt.Printf("[Shell Step] %s\n", s.Name())
	fmt.Printf("Running command: %s\n", s.Command)

	// Execute the shell command.
	cmd := exec.Command("sh", "-c", s.Command)
	output, err := cmd.CombinedOutput()
	outStr := strings.TrimSpace(string(output))

	// Print the command output.
	fmt.Println(outStr)

	// Log the command output.
	run.Log(s.ID(), fmt.Sprintf("Command output: %s", outStr))
	if err != nil {
		return fmt.Errorf("shell command failed: %w", err)
	}

	return nil
}

// SupportsUI indicates if the shell step supports a specific UI type.
func (s *ShellStep) SupportsUI(ui UIType) bool {
	return s.uis[ui] != nil
}

// Render renders the shell step for a specific UI type.
func (s *ShellStep) Render(ui UIType) string {
	if f, ok := s.uis[ui]; ok {
		return f(s)
	}

	return fmt.Sprintf("Shell Step: %s\nCommand: %s", s.Name(), s.Command)
}

// renderCLI renders the shell step for the command-line interface.
func (s *ShellStep) renderCLI(step Step) string {
	return fmt.Sprintf("[Shell Step] %s\nCommand: %s", s.Name(), s.Command)
}

// renderWeb renders the shell step for the web interface.
func (s *ShellStep) renderWeb(step Step) string {
	return fmt.Sprintf("<h2>Shell Step: %s</h2><p>Command: <pre>%s</pre></p>", s.Name(), s.Command)
}
