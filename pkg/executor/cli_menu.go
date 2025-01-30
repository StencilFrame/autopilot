package executor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type (
	// CLIMenuOption represents a menu option.
	CLIMenuOption struct {
		Key         string
		Description string
	}

	// CLIMenu represents a CLI menu.
	CLIMenu struct {
		Options []CLIMenuOption
	}
)

// NewCLIMenu creates a new CLIMenu instance.
func NewCLIMenu() *CLIMenu {
	return &CLIMenu{
		Options: []CLIMenuOption{
			{Key: "y", Description: "Yes, execute step"},
			{Key: "n", Description: "No, skip step"},
			{Key: "c", Description: "Continue to execute current step"},
			{Key: "s", Description: "Skip step"},
			{Key: "b", Description: "Go Back to previous step"},
			{Key: "q", Description: "Quit the runbook"},
			{Key: "h", Description: "print Help"},
		},
	}
}

// WaitForEnter waits for the user to press Enter.
func (m *CLIMenu) WaitForEnter() error {
	fmt.Print("\nPress Enter when continue...")

	// Wait for user confirmation.
	_, err := bufio.NewReader(os.Stdin).ReadString('\n')
	return err
}

// WaitForOption waits for the user to select an option.
func (m *CLIMenu) WaitForOption() (string, error) {
	fmt.Println()
	fmt.Print(`Continue? (enter "h" for help) `)
	options := []string{}
	for _, o := range m.Options {
		options = append(options, o.Key)
	}
	fmt.Printf("[%s] ", strings.Join(options, ","))

	// Wait for user input.
	var option string
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	option = strings.TrimSpace(input)

	option = strings.ToLower(option)
	if option == "h" {
		fmt.Println("Options:")
		for _, opt := range m.Options {
			fmt.Printf("  %s - %s\n", opt.Key, opt.Description)
		}
		return m.WaitForOption()
	}

	return option, nil
}
