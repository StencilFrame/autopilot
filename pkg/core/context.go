package core

import (
	"fmt"
	"sync"
)

// Context manages variables and state during a runbook execution.
type Context struct {
	mu        sync.RWMutex
	variables map[string]interface{} // Map to store variables
}

// NewContext initializes a new execution context.
func NewContext() *Context {
	return &Context{
		variables: make(map[string]interface{}),
	}
}

// Set sets a variable in the context.
func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.variables[key] = value
}

// Get retrieves a variable from the context. Returns nil if the key does not exist.
func (c *Context) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.variables[key]
	return value, exists
}

// GetString retrieves a string variable from the context. Returns an error if the key is missing or not a string.
func (c *Context) GetString(key string) (string, error) {
	value, exists := c.Get(key)
	if !exists {
		return "", fmt.Errorf("key '%s' not found in context", key)
	}

	strValue, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("key '%s' is not a string", key)
	}

	return strValue, nil
}

// GetInt retrieves an integer variable from the context. Returns an error if the key is missing or not an integer.
func (c *Context) GetInt(key string) (int, error) {
	value, exists := c.Get(key)
	if !exists {
		return 0, fmt.Errorf("key '%s' not found in context", key)
	}

	intValue, ok := value.(int)
	if !ok {
		return 0, fmt.Errorf("key '%s' is not an integer", key)
	}

	return intValue, nil
}

// Delete removes a variable from the context.
func (c *Context) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.variables, key)
}

// Keys returns a list of all variable names stored in the context.
func (c *Context) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]string, 0, len(c.variables))
	for key := range c.variables {
		keys = append(keys, key)
	}
	return keys
}

// Clear removes all variables from the context.
func (c *Context) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.variables = make(map[string]interface{})
}
