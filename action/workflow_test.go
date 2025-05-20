package action

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFileNameWithWorkspace(t *testing.T) {
	// given
	tmpDir := t.TempDir()
	os.Setenv("GITHUB_WORKSPACE", tmpDir)
	defer os.Unsetenv("GITHUB_WORKSPACE")

	filename := "foo/bar.txt"
	expected := filepath.Join(tmpDir, filename)
	// when
	got := GetFileName(filename)
	// then
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestGetFileNameWithoutWorkspace(t *testing.T) {
	// given
	os.Unsetenv("GITHUB_WORKSPACE")
	filename := "foo/bar.txt"
	// when
	got := GetFileName(filename)
	// then
	if got != filename {
		t.Errorf("expected %q, got %q", filename, got)
	}
}

func TestWriteJobSummaryAppendsToFile(t *testing.T) {
	// given
	tmpFile, err := os.CreateTemp("", "summary-*.md")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	os.Setenv("GITHUB_STEP_SUMMARY", tmpFile.Name())
	defer os.Unsetenv("GITHUB_STEP_SUMMARY")

	summary := "test summary\n"

	// when
	err = WriteJobSummary(summary)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	// then
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if string(content) != summary {
		t.Errorf("expected file content %q, got %q", summary, string(content))
	}
}
