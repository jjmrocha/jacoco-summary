package action

import (
	"fmt"
	"os"
	"path"
)

func GetFileName(filename string) string {
	workspace := os.Getenv("GITHUB_WORKSPACE")

	if workspace == "" {
		return filename
	}

	return path.Join(workspace, filename)
}

func WriteJobSummary(summary string) error {
	env := os.Getenv("GITHUB_STEP_SUMMARY")

	if env == "" {
		fmt.Println(summary)
		return nil
	}

	file, err := os.OpenFile(env, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open GITHUB_STEP_SUMMARY file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(summary); err != nil {
		return fmt.Errorf("failed to write to GITHUB_STEP_SUMMARY file: %v", err)
	}

	return nil
}
