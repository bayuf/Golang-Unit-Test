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

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		funcName string
		request  string
		expected string
	}{
		{
			funcName: "HelloWorld(Bayu)",
			request:  "Bayu",
			expected: "Hallo Bayu",
		},
		{
			funcName: "HelloWorld(Firman)",
			request:  "Firman",
			expected: "Hallo Firman",
		},
		{
			funcName: "HelloWorld(Silfi)",
			request:  "Silfi",
			expected: "Hallo Silfi",
		},
		{
			funcName: "HelloWorld(Joko)",
			request:  "Joko",
			expected: "Hallo Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.funcName, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
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
