package step

import (
	"autopilot/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStep struct {
	BaseStep
}

func (m *MockStep) Run(run *core.Run) error {
	return nil
}

func (m *MockStep) Render(ui UIType) string {
	return "mock render"
}

func (m *MockStep) SupportsUI(ui UIType) bool {
	return ui == UITypeCLI
}

func TestBaseStep_ID(t *testing.T) {
	step := &BaseStep{IDValue: "step-1"}
	assert.Equal(t, "step-1", step.ID())
}

func TestBaseStep_Name(t *testing.T) {
	step := &BaseStep{NameValue: "Test Step"}
	assert.Equal(t, "Test Step", step.Name())
}

func TestMockStep_Run(t *testing.T) {
	step := &MockStep{}
	err := step.Run(nil)
	assert.NoError(t, err)
}

func TestMockStep_Render(t *testing.T) {
	step := &MockStep{}
	assert.Equal(t, "mock render", step.Render(UITypeCLI))
}

func TestMockStep_SupportsUI(t *testing.T) {
	step := &MockStep{}
	assert.True(t, step.SupportsUI(UITypeCLI))
	assert.False(t, step.SupportsUI(UITypeWeb))
}
