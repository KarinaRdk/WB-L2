package main

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	// Тестовые случаи
	testCases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			name:     "Single letter",
			input:    "a",
			expected: "a",
			err:      nil,
		},
		{
			name:     "Single digit",
			input:    "1",
			expected: "1",
			err:      errors.New("invalid input"),
		},
		{
			name:     "Mixed letters and digits",
			input:    "a4b2c1",
			expected: "aaaabbc",
			err:      nil,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
			err:      nil,
		},
		{
			name:     "Mixed letters and digits_1",
			input:    "a4b2c1d",
			expected: "aaaabbcd",
			err:      nil,
		},
		{
			name:     "Only letters",
			input:    "abcd",
			expected: "abcd",
			err:      nil,
		},
	}

	// Проходим по всем тестовым случаям
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := unpack(tc.input)
			if tc.err!= nil && err == nil {
				t.Errorf("Expected an error, but got none")
			}
			if tc.err == nil && err!= nil {
				t.Errorf("Expected no error, but got %v", err)
			}
			if result!= tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, result)
			}
		})
	}
}
