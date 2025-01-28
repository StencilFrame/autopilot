package runbook

import (
	"autopilot/pkg/step"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	YAML struct {
		InternalName  string     `yaml:"name,omitempty"`        // Name of the runbook (optional)
		Description   string     `yaml:"description,omitempty"` // A brief description of what the runbook does (optional)
		InternalSteps []YAMLStep `yaml:"steps"`                 // An ordered list of steps to execute as part of the runbook (required)
	}

	YAMLStep struct {
		ID     string      `yaml:"id"`             // A unique identifier for the step within the runbook (required)
		Name   string      `yaml:"name,omitempty"` // Human-readable name of the step (optional)
		Type   string      `yaml:"type"`           // Type of the step (required)
		Fields interface{} // Step-specific configuration
	}

	YAMLManualStep struct {
		Instructions string `yaml:"instructions"` // Instructions for the user to perform before continuing (required)
	}

	YAMLShellStep struct {
		Command string `yaml:"command"` // The shell command to execute (required)
	}

	YAMLInputStep struct {
		Prompt    string `yaml:"prompt"`              // The prompt message to display to the user (required)
		Variable  string `yaml:"variable"`            // The name of the variable to store the user input (required)
		Sensitive bool   `yaml:"sensitive,omitempty"` // If true, the input will be hidden in logs and UI displays, and not preserved in the run history (optional)
	}
)

// NewYAML creates a new YAML instance.
func NewYAML() *YAML {
	return &YAML{
		InternalSteps: []YAMLStep{},
	}
}

// Parse reads a YAML file and extracts the steps.
func (y *YAML) Parse(fileName string) []step.Step {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err) // TODO: Handle error
	}

	// Parse the YAML file and populate the steps
	err = yaml.Unmarshal(data, y)
	if err != nil {
		panic(err) // TODO: Handle error
	}

	return y.Steps()
}

// Name returns the name of the runbook.
func (y *YAML) Name() string {
	return y.InternalName
}

// Steps returns the steps in the runbook.
func (y *YAML) Steps() []step.Step {
	steps := make([]step.Step, len(y.InternalSteps))
	for i, s := range y.InternalSteps {
		switch s.Type {
		case "manual":
			step := step.NewManualStep(s.ID, s.Name, s.Fields.(YAMLManualStep).Instructions)
			steps[i] = step
		case "shell":
			step := step.NewShellStep(s.ID, s.Name, s.Fields.(YAMLShellStep).Command)
			steps[i] = step
		case "input":
			// TODO: Implement input step
		default:
			// Add cases for additional step types here
			panic(fmt.Sprintf("unsupported step type: %s", s.Type))
		}
	}

	return steps
}

// UnmarshalYAML unmarshals a YAMLStep from a YAML representation.
func (s *YAMLStep) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type rawStep YAMLStep
	var raw rawStep
	if err := unmarshal(&raw); err != nil {
		return err
	}

	*s = YAMLStep(raw)

	switch s.Type {
	case "manual":
		var manualStep YAMLManualStep
		if err := unmarshal(&manualStep); err != nil {
			return err
		}
		s.Fields = manualStep
	case "shell":
		var shellStep YAMLShellStep
		if err := unmarshal(&shellStep); err != nil {
			return err
		}
		s.Fields = shellStep
	case "input":
		var inputStep YAMLInputStep
		if err := unmarshal(&inputStep); err != nil {
			return err
		}
		s.Fields = inputStep
	// Add cases for additional step types here
	default:
		return fmt.Errorf("unsupported step type: %s", s.Type)
	}

	return nil
}

// MarshalYAML marshals a YAMLStep into a YAML representation.
func (s YAMLStep) MarshalYAML() (interface{}, error) {
	switch v := s.Fields.(type) {
	case YAMLManualStep:
		return struct {
			ID             string `yaml:"id"`
			Name           string `yaml:"name,omitempty"`
			Type           string `yaml:"type"`
			YAMLManualStep `yaml:",inline"`
		}{
			ID:             s.ID,
			Name:           s.Name,
			Type:           s.Type,
			YAMLManualStep: v,
		}, nil
	case YAMLShellStep:
		return struct {
			ID            string `yaml:"id"`
			Name          string `yaml:"name,omitempty"`
			Type          string `yaml:"type"`
			YAMLShellStep `yaml:",inline"`
		}{
			ID:            s.ID,
			Name:          s.Name,
			Type:          s.Type,
			YAMLShellStep: v,
		}, nil
	case YAMLInputStep:
		return struct {
			ID            string `yaml:"id"`
			Name          string `yaml:"name,omitempty"`
			Type          string `yaml:"type"`
			YAMLInputStep `yaml:",inline"`
		}{
			ID:            s.ID,
			Name:          s.Name,
			Type:          s.Type,
			YAMLInputStep: v,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported step type: %T", v)
	}
}
