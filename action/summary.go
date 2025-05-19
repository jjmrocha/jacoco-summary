package action

import (
	"fmt"

	"github.com/jjmrocha/jacoco-summary/jacoco"
)

func MarkdownReport(report *jacoco.Report) string {
	var markdown string

	markdown += "# JaCoCo Report\n"
	markdown += fmt.Sprintf("- Coverage: %s\n", percentageToString(report.Coverage))
	markdown += fmt.Sprintf("- Branches: %s\n", percentageToString(report.BranchCoverage))
	markdown += "\n"
	markdown += "## Class Coverage\n"

	markdown += "Class | Coverage | Branches\n"
	markdown += "------|----------|---------\n"

	for _, class := range report.Details {
		coverage := percentageToString(class.Coverage)
		branchCoverage := percentageToString(class.BranchCoverage)

		markdown += fmt.Sprintf("%s | %s | %s\n", class.ClassName, coverage, branchCoverage)
	}

	return markdown
}

func percentageToString(value int) string {
	if value == -1 {
		return "N/A"
	}

	return fmt.Sprintf("%d%%", value)
}
