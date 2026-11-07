package cpeversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageMagick(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "patch level with dash",
			input:    "7.1.2-34",
			expected: "7.1.2+34",
		},
		{
			name:     "patch level with dash (single digit)",
			input:    "7.1.2-1",
			expected: "7.1.2+1",
		},
		{
			name:     "patch level with dash (zero)",
			input:    "7.1.2-0",
			expected: "7.1.2+0",
		},
		{
			name:     "no patch level",
			input:    "7.1.2",
			expected: "7.1.2",
		},
		{
			name:     "non-numeric suffix preserved",
			input:    "7.1.2-alpha",
			expected: "7.1.2-alpha",
		},
		{
			name:     "multiple dashes, only last one is patch level",
			input:    "7.1.2-beta-34",
			expected: "7.1.2-beta+34",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ImageMagick(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestAlpine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with build suffix",
			input:    "1.2.3-r21",
			expected: "1.2.3",
		},
		{
			name:     "without build suffix",
			input:    "1.2.3",
			expected: "1.2.3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Alpine(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestJVM(t *testing.T) {
	tests := []struct {
		name      string
		version   string
		cpeUpdate string
		expected  string
	}{
		{
			name:      "pre-JEP-223 with update",
			version:   "1.8.0",
			cpeUpdate: "update351",
			expected:  "1.8.0_351",
		},
		{
			name:      "no update",
			version:   "1.8.0",
			cpeUpdate: "",
			expected:  "1.8.0",
		},
		{
			name:      "already has update",
			version:   "1.8.0_351",
			cpeUpdate: "update351",
			expected:  "1.8.0_351",
		},
		{
			name:      "not pre-JEP-223",
			version:   "11.0.1",
			cpeUpdate: "update",
			expected:  "11.0.1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := JVM(test.version, test.cpeUpdate)
			assert.Equal(t, test.expected, result)
		})
	}
}
