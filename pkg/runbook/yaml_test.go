package runbook

import (
	"testing"

	"autopilot/pkg/step"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestYAML_UnmarshalYAML(t *testing.T) {
	yamlContent := `
name: Example Runbook
description: This is an example runbook
steps:
  - id: step-1
    type: manual
    name: Initialize the environment
    instructions: Ensure all prerequisites are installed.
  - id: step-2
    type: shell
    name: Run setup script
    command: ./setup.sh
  - id: step-3
    type: input
    name: Get user input
    prompt: Please enter your name
    variable: user_name
    sensitive: true
`

	var runbook YAML
	err := yaml.Unmarshal([]byte(yamlContent), &runbook)
	require.NoError(t, err)

	assert.Equal(t, "Example Runbook", runbook.InternalName)
	assert.Equal(t, "This is an example runbook", runbook.Description)
	assert.Equal(t, 3, len(runbook.InternalSteps))

	step1 := runbook.InternalSteps[0]
	assert.Equal(t, "step-1", step1.ID)
	assert.Equal(t, "manual", step1.Type)
	assert.Equal(t, "Initialize the environment", step1.Name)
	require.IsType(t, YAMLManualStep{}, step1.Fields)
	assert.Equal(t, "Ensure all prerequisites are installed.", step1.Fields.(YAMLManualStep).Instructions)

	step2 := runbook.InternalSteps[1]
	assert.Equal(t, "step-2", step2.ID)
	assert.Equal(t, "shell", step2.Type)
	assert.Equal(t, "Run setup script", step2.Name)
	require.IsType(t, YAMLShellStep{}, step2.Fields)
	assert.Equal(t, "./setup.sh", step2.Fields.(YAMLShellStep).Command)

	step3 := runbook.InternalSteps[2]
	assert.Equal(t, "step-3", step3.ID)
	assert.Equal(t, "input", step3.Type)
	assert.Equal(t, "Get user input", step3.Name)
	require.IsType(t, YAMLInputStep{}, step3.Fields)
	assert.Equal(t, "Please enter your name", step3.Fields.(YAMLInputStep).Prompt)
	assert.Equal(t, "user_name", step3.Fields.(YAMLInputStep).Variable)
	assert.True(t, step3.Fields.(YAMLInputStep).Sensitive)

	marshalled, err := yaml.Marshal(runbook)
	require.NoError(t, err)
	assert.YAMLEq(t, yamlContent, string(marshalled))
}

func TestYAML_Steps(t *testing.T) {
	yamlContent := `
name: Example Runbook
description: This is an example runbook
steps:
  - id: step-1
    type: manual
    name: Initialize the environment
    instructions: Ensure all prerequisites are installed.
  - id: step-2
    type: shell
    name: Run setup script
    command: ./setup.sh
  - id: step-3
    type: input
    name: Get user input
    prompt: Please enter your name
    variable: user_name
    sensitive: true
`

	var runbook YAML
	err := yaml.Unmarshal([]byte(yamlContent), &runbook)
	require.NoError(t, err)

	steps := runbook.Steps()

	require.Equal(t, 3, len(steps))

	step1 := steps[0].(*step.ManualStep)
	assert.Equal(t, "step-1", step1.ID())
	assert.Equal(t, "Initialize the environment", step1.Name())
	assert.Equal(t, "Ensure all prerequisites are installed.", step1.Instructions)

	step2 := steps[1].(*step.ShellStep)
	assert.Equal(t, "step-2", step2.ID())
	assert.Equal(t, "Run setup script", step2.Name())
	assert.Equal(t, "./setup.sh", step2.Command)

	// TODO: Implement input step

	// step3 := steps[2].(*step.InputStep)
	// assert.Equal(t, "step-3", step3.ID())
	// assert.Equal(t, "Get user input", step3.Name())
	// assert.Equal(t, "Please enter your name", step3.Prompt)
	// assert.Equal(t, "user_name", step3.Variable)
	// assert.True(t, step3.Sensitive)
}
