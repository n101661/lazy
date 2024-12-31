package lazy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicLoader(t *testing.T) {
	loader := NewAtomicLoader(func() string {
		return "hello world"
	})

	assert.False(t, loader.Loaded())
	assert.Equal(t, "hello world", loader.Load())
	assert.True(t, loader.Loaded())
}

func ExampleNewAtomicLoader() {
	loader := NewAtomicLoader(func() string {
		return "hello world"
	})
	fmt.Println(loader.Load())
	// Output: hello world
}

func TestAtomicELoader(t *testing.T) {
	loader := NewAtomicELoader(func() (string, error) {
		return "hello world", nil
	})

	assert.False(t, loader.Loaded())
	value, err := loader.Load()
	assert.Nil(t, err)
	assert.Equal(t, "hello world", value)
	assert.True(t, loader.Loaded())
}

func ExampleNewAtomicELoader() {
	loader := NewAtomicELoader(func() (string, error) {
		return "hello world", nil
	})
	fmt.Println(loader.Load())
	// Output: hello world <nil>
}
