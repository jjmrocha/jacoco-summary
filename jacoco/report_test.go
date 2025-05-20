package jacoco

import (
	"os"
	"testing"
)

func TestReadReportWithSuccess(t *testing.T) {
	// given
	csvContent := `GROUP,PACKAGE,CLASS,INSTRUCTION_MISSED,INSTRUCTION_COVERED,BRANCH_MISSED,BRANCH_COVERED,LINE_MISSED,LINE_COVERED,COMPLEXITY_MISSED,COMPLEXITY_COVERED,METHOD_MISSED,METHOD_COVERED
group,package,Class1,10,90,0,0,1,9,1,9,1,9
group,package,Class2,20,80,4,6,2,8,2,8,2,8
`
	tmpFile := buildTeamFile(t, csvContent)
	defer os.Remove(tmpFile)

	// when
	report, err := ReadReport(tmpFile)
	if err != nil {
		t.Fatalf("ReadReport returned error: %v", err)
	}

	// then
	if len(report.Details) != 2 {
		t.Errorf("expected 2 class coverages, got %d", len(report.Details))
	}

	if report.Details[0].ClassName != "package.Class1" {
		t.Errorf("expected class name 'package.Class1', got '%s'", report.Details[0].ClassName)
	}

	if report.Details[0].Coverage != 90 {
		t.Errorf("expected coverage 90, got %d", report.Details[0].Coverage)
	}

	if report.Details[0].BranchCoverage != Coverage_NA {
		t.Errorf("expected branch coverage %d, got %d", Coverage_NA, report.Details[0].BranchCoverage)
	}

	if report.Details[1].ClassName != "package.Class2" {
		t.Errorf("expected class name 'package.Class2', got '%s'", report.Details[1].ClassName)
	}

	if report.Details[1].Coverage != 80 {
		t.Errorf("expected coverage 80, got %d", report.Details[1].Coverage)
	}

	if report.Details[1].BranchCoverage != 60 {
		t.Errorf("expected branch coverage 60, got %d", report.Details[1].BranchCoverage)
	}

	if report.Coverage != 85 {
		t.Errorf("expected total coverage 85, got %d", report.Coverage)
	}

	if report.BranchCoverage != 60 {
		t.Errorf("expected branch coverage 70, got %d", report.BranchCoverage)
	}
}

func buildTeamFile(t *testing.T, csvContent string) string {
	tmpFile, err := os.CreateTemp("", "jacoco-*.csv")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	if _, err := tmpFile.Write([]byte(csvContent)); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	return tmpFile.Name()
}
