package runbook

import (
	"autopilot/pkg/step"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarkdown_Parse(t *testing.T) {
	// Create a temporary markdown file for testing
	content := `
# Runbook

1. Step 1: Initialize the environment
2. Step 2: Do something
   Additional information about step 2
3. Step 3: Do something else
   Additional information about step 3
4. Step 4: Do something automatically
   ` + "```" + `
   echo "Hello, world!"
   ` + "```" + `

Additional information about the runbook

1. Step 3: Do something else
2. Step 4: Do something else
`
	file, err := os.CreateTemp("", "runbook*.md")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(content))
	require.NoError(t, err)
	file.Close()

	// Parse the markdown file
	md := NewMarkdown()
	steps := md.Parse(file.Name())

	// Validate the parsed steps
	expectedSteps := []string{
		"[Manual Step] Step 1: Initialize the environment",
		"[Manual Step] Step 2: Do something\nAdditional information about step 2",
		"[Manual Step] Step 3: Do something else\nAdditional information about step 3",
		"[Shell Step] Step 4: Do something automatically\nCommand: echo \"Hello, world!\"",
	}

	require.Equal(t, len(expectedSteps), len(steps))

	for i, step := range steps {
		assert.Equal(t, expectedSteps[i], step.Render("CLI"))
	}
}

func TestMarkdown_Steps(t *testing.T) {
	// Create a temporary markdown file for testing
	content := `
# Runbook

1. Step 1: Initialize the environment
   Ensure all prerequisites are installed.
2. Step 2: Do something
   Additional information about step 2
3. Step 3: Do something else
   Additional information about step 3
4. Step 4: Do something automatically
   ` + "```" + `
   echo "Hello, world!"
   ` + "```" + `
`
	file, err := os.CreateTemp("", "runbook*.md")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(content))
	require.NoError(t, err)
	file.Close()

	// Parse the markdown file
	md := NewMarkdown()
	md.Parse(file.Name())

	// Validate the parsed steps
	steps := md.Steps()
	require.Equal(t, 4, len(steps))

	step1, ok := steps[0].(*step.ManualStep)
	require.True(t, ok, "step is not a ManualStep")
	assert.Equal(t, "step-1", step1.ID())
	assert.Equal(t, "Step 1: Initialize the environment", step1.Name())
	assert.Equal(t, "Ensure all prerequisites are installed.", step1.Instructions)

	step2, ok := steps[1].(*step.ManualStep)
	require.True(t, ok, "step is not a ManualStep")
	assert.Equal(t, "step-2", step2.ID())
	assert.Equal(t, "Step 2: Do something", step2.Name())
	assert.Equal(t, "Additional information about step 2", step2.Instructions)

	step3, ok := steps[2].(*step.ManualStep)
	require.True(t, ok, "step is not a ManualStep")
	assert.Equal(t, "step-3", step3.ID())
	assert.Equal(t, "Step 3: Do something else", step3.Name())
	assert.Equal(t, "Additional information about step 3", step3.Instructions)

	step4, ok := steps[3].(*step.ShellStep)
	require.True(t, ok, "step is not a ShellStep")
	assert.Equal(t, "step-4", step4.ID())
	assert.Equal(t, "Step 4: Do something automatically", step4.Name())
	assert.Equal(t, "echo \"Hello, world!\"", step4.Command)
}
