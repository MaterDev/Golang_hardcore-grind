# Golang Hardcore Grind

A collection of algorithmic coding challenges with a central harness to run tests per challenge. You write solutions; tests are provided.

## Structure

- `go.mod`
- `cmd/harness/main.go`
- `challenges/<slug>/...`
- `scripts/run.sh`
- `CHALLENGES.md`

## Requirements

- Go 1.22+

## Usage

- List challenges:
  - `./scripts/run.sh -list`
- Run one challenge:
  - `./scripts/run.sh -challenge two-sum`
- Run all challenges:
  - `./scripts/run.sh -all`
- Run specific tests via -run (supports subtests):
  - `./scripts/run.sh -challenge two-sum -run '^TestTwoSum/basic$'`
  - `./scripts/run.sh -challenge two-sum -- -run '^TestTwoSum/basic$'`
- Forward more flags to `go test` either explicitly or after `--`:
  - `-v`, `-race`, `-cover`, `-count`, `-timeout`, `-failfast`, `-bench`
- Benchmark example (passthrough after `--`):
  - `./scripts/run.sh -challenge two-sum -- -bench . -benchmem`

## Challenges checklist

Track planned and completed challenges in `CHALLENGES.md`. We will add items as we go.

## Adding a solution

- Create `challenges/<slug>/solution.go` with the function(s) referenced by tests.

## Setup

- `chmod +x scripts/run.sh`
- `go run ./cmd/harness -list`
