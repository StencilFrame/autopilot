package runbook

import (
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
	}

	require.Equal(t, len(expectedSteps), len(steps))

	for i, step := range steps {
		assert.Equal(t, expectedSteps[i], step.Render("CLI"))
	}
}
