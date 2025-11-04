package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Before test
	fmt.Println("Jalankan code sebelum tes")

	// Jalankan semua tes di package yang sama
	m.Run()

	// After Test
	fmt.Println("Jalankan code setelah tes selesai dilakukan")
}

func TestSubTest(t *testing.T) {
	t.Run("Bayu", func(t *testing.T) {
		result := HelloWorld("Bayu")
		assert.Equal(t, result, "Hallo Bayu", "Seharusnya Hallo Bayu")
	})
	t.Run("Firman", func(t *testing.T) {
		result := HelloWorld("Firman")
		assert.Equal(t, result, "Hallo Firman", "Seharusnya Hallo Firman")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("tidak bisa di jalankan di OS Windows")
	}

	result := HelloWorld("Bayu")
	assert.Equal(t, result, "Hallo Bayu", "Seharusnya Hallo Bayu")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bayu")
	assert.Equal(t, result, "Hallo Bayu", "Seharusnya Hallo Bayu")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bayu")
	require.Equal(t, result, "Hallo Bayu", "Seharusnya Hallo Bayu")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bayu")

	if result != "Hallo Bayu" {
		// Error
		t.Error("Seharusnya 'Hallo Bayu'")

	}

	fmt.Println("Tes Done")
}
func TestHelloWorldName(t *testing.T) {
	result := HelloWorld("Silfi")

	if result != "Hallo Silfi" {
		// Fatal
		t.Fatal("Seharusnya Hallo Silfi")
	}
}
