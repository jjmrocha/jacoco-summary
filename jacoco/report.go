package jacoco

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

const Coverage_NA = -1

type Report struct {
	Details        []ClassCoverage
	Coverage       int
	BranchCoverage int
}

type ClassCoverage struct {
	ClassName       string
	Missed          int
	Covered         int
	MissedBranches  int
	CoveredBranches int
	Coverage        int
	BranchCoverage  int
}

func ReadReport(fileName string) (*Report, error) {
	coverageRows, err := readReportFile(fileName)
	if err != nil {
		return nil, err
	}

	report := buildReport(coverageRows)
	return report, nil
}

func readReportFile(fileName string) ([]ClassCoverage, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading report file %s: %v", fileName, err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	classRows := make([]ClassCoverage, 0)
	line := 0

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading report file %s: %v", fileName, err)
		}

		if line == 0 {
			if len(row) < 13 {
				return nil, fmt.Errorf("not enough fields in file %s: %v", fileName, err)
			}
		} else {
			classCoverage, err := parseRow(row)
			if err != nil {
				return nil, fmt.Errorf("error parsing line %d in report file %s: %v", line, fileName, err)
			}

			classRows = append(classRows, classCoverage)
		}

		line++
	}

	return classRows, nil
}

func parseRow(row []string) (ClassCoverage, error) {
	missed, err := strconv.Atoi(row[3])
	if err != nil {
		return ClassCoverage{}, fmt.Errorf("error converting INSTRUCTION_MISSED to int: %v", err)
	}

	covered, err := strconv.Atoi(row[4])
	if err != nil {
		return ClassCoverage{}, fmt.Errorf("error converting INSTRUCTION_COVERED to int: %v", err)
	}

	missedBranches, err := strconv.Atoi(row[5])
	if err != nil {
		return ClassCoverage{}, fmt.Errorf("error converting BRANCH_MISSED to int: %v", err)
	}

	coveredBranches, err := strconv.Atoi(row[6])
	if err != nil {
		return ClassCoverage{}, fmt.Errorf("error converting BRANCH_COVERED to int: %v", err)
	}

	classCoverage := ClassCoverage{
		ClassName:       row[1] + "." + row[2],
		Missed:          missed,
		Covered:         covered,
		MissedBranches:  missedBranches,
		CoveredBranches: coveredBranches,
		Coverage:        percentage(missed, covered),
		BranchCoverage:  percentage(missedBranches, coveredBranches),
	}

	return classCoverage, nil
}

func percentage(missed, covered int) int {
	if missed+covered == 0 {
		return Coverage_NA
	}

	return int(float32(covered) / float32(missed+covered) * 100)
}

func buildReport(coverageRows []ClassCoverage) *Report {
	totalMissed := 0
	totalCovered := 0
	totalMissedBranches := 0
	totalCoveredBranches := 0

	for _, coverage := range coverageRows {
		totalMissed += coverage.Missed
		totalCovered += coverage.Covered
		totalMissedBranches += coverage.MissedBranches
		totalCoveredBranches += coverage.CoveredBranches
	}

	report := Report{
		Details:        coverageRows,
		Coverage:       percentage(totalMissed, totalCovered),
		BranchCoverage: percentage(totalMissedBranches, totalCoveredBranches),
	}

	return &report
}
