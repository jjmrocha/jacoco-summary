package action

import (
	"strings"
	"testing"

	"github.com/jjmrocha/jacoco-summary/jacoco"
)

func TestMarkdownReport(t *testing.T) {
	// given
	report := &jacoco.Report{
		Coverage:       80,
		BranchCoverage: 70,
		Details: []jacoco.ClassCoverage{
			{
				ClassName:      "foo.Bar",
				Coverage:       80,
				BranchCoverage: 70,
			},
			{
				ClassName:      "foo.Baz",
				Coverage:       90,
				BranchCoverage: jacoco.Coverage_NA,
			},
		},
	}
	// when
	md := MarkdownReport(report)
	// then
	if !strings.Contains(md, "# JaCoCo Report") {
		t.Error("expected report header")
	}

	if !strings.Contains(md, "- Coverage: 80%") {
		t.Error("expected overall coverage")
	}

	if !strings.Contains(md, "- Branches: 70%") {
		t.Error("expected branch coverage")
	}

	if !strings.Contains(md, "foo.Bar | 80% | 70%") {
		t.Error("expected class foo.Bar row")
	}

	if !strings.Contains(md, "foo.Baz | 90% | N/A") {
		t.Error("expected class foo.Baz row")
	}
}

func TestPercentageToString(t *testing.T) {
	if got := percentageToString(85); got != "85%" {
		t.Errorf("expected 85%%, got %q", got)
	}

	if got := percentageToString(jacoco.Coverage_NA); got != "N/A" {
		t.Errorf("expected N/A, got %q", got)
	}
}
