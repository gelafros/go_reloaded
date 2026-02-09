package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("usage: go run . <input_file> <output_file>")
	}

	inputPath := args[1]
	outputPath := args[2]

	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	result := function.Process(string(data))

	if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}