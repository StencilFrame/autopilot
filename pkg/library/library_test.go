package library

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLibrary(t *testing.T) {
	lib := NewLibrary()
	assert.NotNil(t, lib)
	assert.Empty(t, lib.item)
}

func TestAdd(t *testing.T) {
	lib := NewLibrary()
	item := Item{Description: "Test command", Command: "echo test"}

	err := lib.Add(item)
	assert.Nil(t, err)
	assert.Len(t, lib.item, 1)

	// Test adding duplicate command
	err = lib.Add(item)
	assert.ErrorAs(t, err, &ErrCommandExists)
}

func TestLoadAndSave(t *testing.T) {
	lib := NewLibrary()
	item := Item{Description: "Test command", Command: "echo test"}
	err := lib.Add(item)
	assert.Nil(t, err)

	fileName := "test_library.json"
	defer os.Remove(fileName)

	err = lib.Save(fileName)
	assert.Nil(t, err)

	newLib := NewLibrary()
	err = newLib.Load(fileName)
	assert.Nil(t, err)
	assert.Len(t, newLib.item, 1)
	assert.Equal(t, item.Description, newLib.item[0].Description)
	assert.Equal(t, item.Command, newLib.item[0].Command)
}

func TestItems(t *testing.T) {
	lib := NewLibrary()
	item1 := Item{Description: "Test command 1", Command: "echo test1"}
	item2 := Item{Description: "Test command 2", Command: "echo test2"}
	lib.Add(item1)
	lib.Add(item2)

	items := lib.Items(nil)
	assert.Len(t, items, 2)
	assert.Contains(t, items[0], item1.Command)
	assert.Contains(t, items[0], item1.Description+"")
	assert.Contains(t, items[1], item2.Command)
	assert.Contains(t, items[1], item2.Description)
}
