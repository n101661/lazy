package lazy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoader(t *testing.T) {
	loader := NewLoader(func() string {
		return "hello world"
	})

	assert.False(t, loader.Loaded())
	assert.Equal(t, "hello world", loader.Load())
	assert.True(t, loader.Loaded())
}

func ExampleNewLoader() {
	loader := NewLoader(func() string {
		return "hello world"
	})
	fmt.Println(loader.Load())
	// Output: hello world
}

func TestELoader(t *testing.T) {
	loader := NewELoader(func() (string, error) {
		return "hello world", nil
	})

	assert.False(t, loader.Loaded())
	value, err := loader.Load()
	assert.Nil(t, err)
	assert.Equal(t, "hello world", value)
	assert.True(t, loader.Loaded())
}

func ExampleNewELoader() {
	loader := NewELoader(func() (string, error) {
		return "hello world", nil
	})
	fmt.Println(loader.Load())
	// Output: hello world <nil>
}
