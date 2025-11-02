package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bay")
	assert.Equal(t, result, "Hallo Bayu", "Seharusnya Hallo Bayu")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bayu")
	require.Equal(t, result, "Hallo Bayu", "Seharusnya Hallo Bayu")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bay")

	if result != "Hallo Bayu" {
		// Error
		t.Error("Seharusnya 'Hallo Bayu'")

	}

	fmt.Println("Tes Done")
}
func TestHelloWorldName(t *testing.T) {
	result := HelloWorld("Bayu")

	if result != "Hallo Silfi" {
		// Fatal
		t.Fatal("Seharusnya Hallo Silfi")
	}
}
