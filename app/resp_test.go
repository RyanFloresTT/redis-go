package main

import "testing"

func TestFromRESP(t *testing.T) {
	tests := []struct {
		name     string
		input    command
		expected string
	}{
		{
			name:     "PONG Response",
			input:    "+PONG\r\n",
			expected: "PONG",
		},
		{
			name:     "SET Command",
			input:    "*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$5\r\nAlice\r\n",
			expected: "SET name Alice",
		},
		{
			name:     "GET Command",
			input:    "*2\r\n$3\r\nGET\r\n$4\r\nname\r\n",
			expected: "GET name",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.FromRESP()

			if result != tt.expected {
				t.Errorf("[%s] : Actual: %q. Expected: %q.", tt.name, result, tt.expected)
			}
		})
	}
}

func TestToString(t *testing.T) {

}
