package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

// Helper function to launch the editor
func LaunchEditor(editor string, content string, pos int) (string, error) {
	args := []string{}
	if editor == "vim" || editor == "vi" { // TODO: add support for other editors. Make it configurable
		if pos > 0 {
			// move cursor n bytes
			args = append(args, fmt.Sprintf("+normal %dgo", pos))
		}
	}

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "autopilot-editor-input")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}
	defer os.Remove(tmpFile.Name())
	args = append(args, tmpFile.Name())

	// add description to the file
	if _, err := tmpFile.WriteString(content + "\n"); err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	// Open the editor
	cmd := exec.Command(editor, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: Failed to launch the editor")
		return "", err
	}

	// Read the content from the temporary file
	out, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	// Trim the trailing newline
	out = bytes.Trim(out, " \n")

	return string(out), nil
}

// Parse the content
func ParseContent(content string, obj interface{}) error {
	return yaml.Unmarshal([]byte(content), obj)
}
