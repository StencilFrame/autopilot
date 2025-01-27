package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContext_SetAndGet(t *testing.T) {
	ctx := NewContext()
	ctx.Set("key1", "value1")

	value, exists := ctx.Get("key1")
	assert.True(t, exists)
	assert.Equal(t, "value1", value)
}

func TestContext_GetString(t *testing.T) {
	ctx := NewContext()
	ctx.Set("key1", "value1")

	value, err := ctx.GetString("key1")
	require.NoError(t, err)
	assert.Equal(t, "value1", value)

	_, err = ctx.GetString("key2")
	assert.Error(t, err)

	ctx.Set("key3", 123)
	_, err = ctx.GetString("key3")
	assert.Error(t, err)
}

func TestContext_GetInt(t *testing.T) {
	ctx := NewContext()
	ctx.Set("key1", 123)

	value, err := ctx.GetInt("key1")
	require.NoError(t, err)
	assert.Equal(t, 123, value)

	_, err = ctx.GetInt("key2")
	assert.Error(t, err)

	ctx.Set("key3", "value1")
	_, err = ctx.GetInt("key3")
	assert.Error(t, err)
}

func TestContext_Delete(t *testing.T) {
	ctx := NewContext()
	ctx.Set("key1", "value1")
	ctx.Delete("key1")

	_, exists := ctx.Get("key1")
	assert.False(t, exists)
}

func TestContext_Keys(t *testing.T) {
	ctx := NewContext()
	ctx.Set("key1", "value1")
	ctx.Set("key2", "value2")

	keys := ctx.Keys()
	expectedKeys := []string{"key1", "key2"}

	assert.Equal(t, len(expectedKeys), len(keys))

	for _, key := range expectedKeys {
		found := false
		for _, k := range keys {
			if k == key {
				found = true
				break
			}
		}
		assert.True(t, found)
	}
}

func TestContext_Clear(t *testing.T) {
	ctx := NewContext()
	ctx.Set("key1", "value1")
	ctx.Set("key2", "value2")
	ctx.Clear()

	assert.Equal(t, 0, len(ctx.Keys()))
}
