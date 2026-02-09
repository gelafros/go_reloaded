package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	// Setup temporary files for testing
	inputFile, _ := os.CreateTemp("", "input_*.txt")
	outputFile, _ := os.CreateTemp("", "output_*.txt")
	defer os.Remove(inputFile.Name())
	defer os.Remove(outputFile.Name())

	tests := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			name:      "Not enough arguments",
			args:      []string{"cmd"},
			expectErr: true,
		},
		{
			name:      "Input file missing",
			args:      []string{"cmd", "missing.txt", outputFile.Name()},
			expectErr: true,
		},
		{
			name:      "Valid arguments",
			args:      []string{"cmd", inputFile.Name(), outputFile.Name()},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := run(tt.args)
			if (err != nil) != tt.expectErr {
				t.Errorf("run() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
