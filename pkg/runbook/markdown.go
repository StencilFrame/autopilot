package runbook

import (
	"bytes"
	"donot/pkg/step"
	"fmt"
	"os"
	"strings"

	"github.com/russross/blackfriday/v2"
)

type (
	// Markdown represents a runbook in markdown format.
	Markdown struct {
		name  string      // Name of the runbook (optional)
		steps []step.Step // List of steps in the runbook
	}
)

// NewMarkdown creates a new Markdown instance.
func NewMarkdown() *Markdown {
	return &Markdown{
		steps: []step.Step{},
	}
}

// Parse reads a markdown file and extracts the steps.
func (m *Markdown) Parse(fileName string) []step.Step {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err) // TODO: Handle error
	}

	// Parse the markdown file and populate the steps
	parser := blackfriday.New(blackfriday.WithExtensions(blackfriday.CommonExtensions))
	ast := parser.Parse(data)

	// Walk through the AST nodes to extract the first ordered list items.
	ast.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		// Check if the node is a List and it's an ordered list.
		if node.Type == blackfriday.List && node.ListFlags&blackfriday.ListTypeOrdered != 0 {
			// Iterate over each list item in the ordered list.
			for item := node.FirstChild; item != nil; item = item.Next {
				// Extract the text content of the list item.
				text, code := extractText(item)
				if code != "" {
					m.AddCodeStep(text, code)
				} else {
					m.AddManualStep(text)
				}
			}
			// Stop processing the AST because we found the first ordered list.
			return blackfriday.Terminate
		}
		return blackfriday.GoToNext
	})

	return m.steps
}

// AddCodeStep adds a new code step to the runbook.
func (m *Markdown) AddCodeStep(name, code string) {
	stepId := fmt.Sprintf("step-%d", len(m.steps)+1)
	codeStep := step.NewShellStep(stepId, name, code)
	m.steps = append(m.steps, codeStep)
}

// AddStep adds a new step to the runbook.
func (m *Markdown) AddManualStep(raw string) {
	stepId := fmt.Sprintf("step-%d", len(m.steps)+1)
	s := strings.SplitN(raw, "\n", 2)
	name := s[0]
	instructions := ""
	if len(s) > 1 {
		instructions = s[1]
	}
	manualStep := step.NewManualStep(stepId, name, instructions)
	m.steps = append(m.steps, manualStep)
}

// Name returns the name of the runbook.
func (m *Markdown) Name() string {
	return m.name
}

// Steps returns the list of steps in the runbook.
func (m *Markdown) Steps() []step.Step {
	return m.steps
}

// extractText helper function for extracting plain text from a node
func extractText(node *blackfriday.Node) (text string, code string) {
	var buffer bytes.Buffer
	var codeBlock bytes.Buffer
	node.Walk(func(n *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		literal := bytes.TrimSpace(n.Literal)
		if n.Type == blackfriday.Code {
			codeBlock.Write(literal)
			return blackfriday.SkipChildren
		}
		if len(literal) > 0 {
			buffer.Write(literal)
			buffer.WriteString("\n")
		}
		return blackfriday.GoToNext
	})
	buffer.Truncate(buffer.Len() - 1) // Remove the trailing newline
	return buffer.String(), codeBlock.String()
}
