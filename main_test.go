package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	if err := run([]string{"go-reloaded"}); err == nil {
		t.Fatal("expected error when called with wrong argument count")
	}

	tmpDir := t.TempDir()
	inFile := filepath.Join(tmpDir, "input.txt")
	outFile := filepath.Join(tmpDir, "output.txt")

	if err := os.WriteFile(inFile, []byte("1E (hex) files were added"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"go-reloaded", inFile, outFile}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "30 files were added" {
		t.Fatalf("expected %q, got %q", "30 files were added", string(got))
	}
}
