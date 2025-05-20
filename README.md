# jacoco-summary
GitHub Action to add coverage from a JaCoCo report to the workflow job summary.

## Usage
Add the following step to your GitHub Actions workflow after your tests and JaCoCo report generation:

```yaml
- name: Add JaCoCo Coverage Summary
  uses: jjmrocha/jacoco-summary@v1
  with:
    jacoco-csv-file: path/to/jacoco.csv
```

- `jacoco-csv-file`: Path to your JaCoCo CSV report file (defaults to `build/reports/jacoco/test/jacocoTestReport.csv`).

## Example Workflow

```yaml
name: CI

on:
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up JDK
        uses: actions/setup-java@v4
        with:
          java-version: '17'
          distribution: 'temurin'

      - name: Build and test with JaCoCo
        run: ./gradlew test jacocoTestReport

      - name: Add JaCoCo Coverage Summary
        uses: jjmrocha/jacoco-summary@v1
```

## Output
The action parses the JaCoCo CSV report and adds a Markdown summary to the GitHub Actions job summary, showing overall and per-class coverage.

## License
MIT