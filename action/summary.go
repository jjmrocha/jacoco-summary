package action

import (
	"fmt"
	"strings"

	"github.com/jjmrocha/jacoco-summary/jacoco"
)

func MarkdownReport(report *jacoco.Report) string {
	var b strings.Builder

	b.WriteString("# JaCoCo Report\n")
	b.WriteString(fmt.Sprintf("- Coverage: %s\n", percentageToString(report.Coverage)))
	b.WriteString(fmt.Sprintf("- Branches: %s\n", percentageToString(report.BranchCoverage)))
	b.WriteString("\n")

	b.WriteString("## Class Coverage\n")
	b.WriteString("Class | Coverage | Branches\n")
	b.WriteString("------|----------|---------\n")

	for _, class := range report.Details {
		coverage := percentageToString(class.Coverage)
		branchCoverage := percentageToString(class.BranchCoverage)
		b.WriteString(fmt.Sprintf("%s | %s | %s\n", class.ClassName, coverage, branchCoverage))
	}

	return b.String()
}

func percentageToString(value int) string {
	if value == jacoco.Coverage_NA {
		return "N/A"
	}

	return fmt.Sprintf("%d%%", value)
}
