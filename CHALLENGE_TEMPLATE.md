# Challenge Template

Use this template to add a new challenge that fits the repo’s flow and style. Copy this file’s sections into the new challenge’s `solution.go` header comment and into the tests as needed.

---

## Guide

### Metadata

- Title: `TITLE_HERE`
- Slug (kebab-case): `slug-here`
- Package name (drop hyphens): `slughere`
- Directory: `challenges/slug-here/`

### Problem

- Description: Briefly state the task. Avoid revealing the solution.
- Signature (Go):
  - Example: `func solve(input ...type) outputType`
- Constraints:
  - Input size/limits
  - Edge cases to consider
  - Return value rules
- Time/Space targets (non-binding but directional):
  - Time: `O(...)`
  - Space: `O(...)`

### Examples

- Example 1
  - Input: `...`
  - Output: `...`
- Example 2
  - Input: `...`
  - Output: `...`

Keep examples illustrative but not exhaustive. Tests will contain more.

### Hints (non-revealing)

- Nudge toward the right data structure or pattern (e.g., two-pointers, map, heap) without giving the algorithm away.
- Mention pitfalls (off-by-one, duplicate handling) if relevant.

---

### Files to create

Create the following in `challenges/slug-here/`:

- `solution.go`
- `solution_test.go`

### solution.go

- Package: `slughere`
- Contains only the function signature(s) and a header comment with this template content adapted.
- Should compile but NOT solve the problem (keep the body as a stub returning zero-values or TODO).
- Do not add prints. Do not import unused packages.

Header boilerplate to copy (example):

```go
// Package slughere provides the <TITLE_HERE> challenge implementation point.
//
// Problem: <short description>
//
// Requirements
// - <bullet 1>
// - <bullet 2>
//
// Examples
//   <short examples>
//
// Approach Hints
// - <non-revealing hint>
//
// Running tests
// - All tests for this challenge live in solution_test.go.
// - From repo root you can run just this challenge via the harness:
//     ./scripts/run.sh -challenge slug-here -v
// - Run a specific subtest with:
//     ./scripts/run.sh -challenge slug-here -run '^TestName/subcase$'
package slughere

// TODO: implement the required function signature(s) below so tests pass.
```

### solution_test.go

- Package: `slughere`
- Cover typical, edge, and negative cases.
- Use subtests via `t.Run` with readable names.
- Add a small `Benchmark...` where it makes sense.
- Keep tests non-revealing. Do not include algorithmic hints or step-by-step guidance.

Test skeleton to copy:

```go
package slughere

import "testing"

func TestSolve(t *testing.T) {
 cases := []struct {
  name string
  in   any // replace with proper fields
  want any // replace with proper fields
 }{
  {"basic", /* in */ nil, /* want */ nil},
  // add more cases
 }
 for _, tc := range cases {
  t.Run(tc.name, func(t *testing.T) {
   got := /* call your function */
   if /* compare got vs want */ false {
    t.Fatalf("got %v want %v", got, tc.want)
   }
  })
 }
}

func BenchmarkSolve(b *testing.B) {
 in := /* representative input */
 for i := 0; i < b.N; i++ {
  _ = /* call your function */
 }
}
```

---

### Authoring rules

- Naming
  - Directory slugs in kebab-case: `slug-here`.
  - Package name is the slug without hyphens: `slughere`.
  - Test names: `TestXxx`, subtests `t.Run("case-name", ...)`.
- Style
  - No solution hints in comments.
  - Keep stubs minimal and compiling.
  - Lints: avoid unused imports/vars; prefer table-driven tests; meaningful subtest names.
- Difficulty
  - Ensure the initial `solution.go` does not already pass the tests.
  - Tests should clearly communicate expectations through inputs/outputs and error messages.

---

### Usage with harness

- List challenges: `./scripts/run.sh -list`
- Run one challenge: `./scripts/run.sh -challenge slug-here -v`
- Run a subtest: `./scripts/run.sh -challenge slug-here -run '^TestSolve/basic$'`
- Passthrough to `go test` after `--` (e.g., benchmarks):
  - `./scripts/run.sh -challenge slug-here -- -bench . -benchmem`
