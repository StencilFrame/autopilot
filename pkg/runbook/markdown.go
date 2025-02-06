package runbook

import (
	"autopilot/pkg/step"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
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
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	rootNode := p.Parse(data)

	// Walk through the AST nodes to extract the first ordered list items.
	ast.WalkFunc(rootNode, func(node ast.Node, entering bool) ast.WalkStatus {
		// Check if the node is a List and it's an ordered list.
		if list, ok := node.(*ast.List); ok && entering && list.ListFlags&ast.ListTypeOrdered != 0 {
			// Iterate over each list item in the ordered list.
			for _, item := range list.Children {
				// Extract the text content of the list item.
				text, code := extractText(item)
				if code != "" {
					m.AddCodeStep(text, code)
				} else {
					m.AddManualStep(text)
				}
			}
			// Stop processing the AST because we found the first ordered list.
			return ast.Terminate
		}
		return ast.GoToNext
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
func extractText(node ast.Node) (text string, code string) {
	var buffer bytes.Buffer
	var codeBlock bytes.Buffer
	ast.WalkFunc(node, func(n ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}

		literal := bytes.TrimSpace([]byte(getContent(n)))
		if block, ok := n.(*ast.CodeBlock); ok {
			if block.Info != nil && bytes.Equal(block.Info, []byte("sh")) {
				codeBlock.Write(literal)
				return ast.GoToNext
			}
		}
		if len(literal) > 0 {
			buffer.Write(literal)
			buffer.WriteString("\n")
		}
		return ast.GoToNext
	})
	if buffer.Len() > 0 {
		buffer.Truncate(buffer.Len() - 1) // Remove the trailing newline
	}
	return buffer.String(), codeBlock.String()
}

func getContent(node ast.Node) string {
	if c := node.AsContainer(); c != nil {
		return contentToString(c.Literal, c.Content)
	}
	leaf := node.AsLeaf()
	return contentToString(leaf.Literal, leaf.Content)
}

func contentToString(d1 []byte, d2 []byte) string {
	if d1 != nil {
		return string(d1)
	}
	if d2 != nil {
		return string(d2)
	}
	return ""
}
