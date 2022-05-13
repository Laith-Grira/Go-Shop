package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewReceipt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected receipt
	}{
		{
			name:  "Test with a valid client name 1",
			input: "James",
			expected: receipt{
				clientName: "James",
				items:      map[string]float64{},
				cash:       0.0,
			},
		},
		{
			name:  "Test with a valid client name 2",
			input: "Maria",
			expected: receipt{
				clientName: "Maria",
				items:      map[string]float64{},
				cash:       0.0,
			},
		},
		{
			name:  "Test with a numerical client name",
			input: "Bob123",
			expected: receipt{
				clientName: "Bob123",
				items:      map[string]float64{},
				cash:       0.0,
			},
		},
		{
			name:  "Test with a client full name",
			input: "Susan Flyn",
			expected: receipt{
				clientName: "Susan Flyn",
				items:      map[string]float64{},
				cash:       0.0,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, newReceipt(tc.input))
		})
	}
}
