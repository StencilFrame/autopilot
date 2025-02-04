package library

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type (
	// Library is a collection of items
	Library struct {
		mutex sync.Mutex
		item  []Item
	}

	// Item is an item in the library
	Item struct {
		Description string `json:"description,omitempty" yaml:"description,omitempty"` // Description of the command
		Command     string `json:"command" yaml:"command"`                             // Command to execute
	}
)

var (
	// ErrEmptyCommand is returned when the command is empty
	ErrEmptyCommand = errors.New("empty command")
	// ErrCommandExists is returned when the command already exists
	ErrCommandExists = errors.New("command already exists")
)

// NewLibrary creates a new Library instance
func NewLibrary() *Library {
	return &Library{
		item: []Item{},
	}
}

// Add adds an item to the library
func (l *Library) Add(item Item) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// check for duplicates
	for _, i := range l.item {
		if i.Command == item.Command {
			return ErrCommandExists
		}
	}

	l.item = append(l.item, item)
	return nil
}

// Load loads the library from a file
func (l *Library) Load(fileName string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Clear the library
	l.item = []Item{}
	// Load the library from a file
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, &l.item)
}

// Save saves the library to a file
func (l *Library) Save(fileName string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Save the library to a file
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(l.item)
}

// Items returns concatenated fields of all items
func (l *Library) Items(fields []string) []string {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// TODO: Implement the fields
	// default fields
	// if len(fields) == 0 {
	// 	fields = []string{"description", "command"}
	// }

	var items []string
	for i, item := range l.item {
		// TODO: Implement the fields

		desc := trim(item.Description, 50)
		item := fmt.Sprintf("%-5d %-50s %s", i+1, desc, item.Command)
		items = append(items, strings.TrimSpace(item))
	}
	return items
}

// GetItemByCommand returns an item by command
func (l *Library) GetItemByCommand(command string) (*Item, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	re := regexp.MustCompile(`^\d+`)
	match := re.Find([]byte(command))
	if match == nil {
		return nil, fmt.Errorf("invalid command: %s", command)
	}
	i, err := strconv.Atoi(string(match))
	if err != nil {
		return nil, err
	}
	if i < 1 || i > len(l.item) {
		return nil, fmt.Errorf("invalid command: %s", command)
	}
	return &l.item[i-1], nil
}

func trim(s string, max int) string {
	if len(s) > max {
		return s[:max-2] + "··"
	}
	return s
}

// Validate validates the item
func (l *Item) Validate() error {
	if l.Command == "" {
		return ErrEmptyCommand
	}
	return nil
}
