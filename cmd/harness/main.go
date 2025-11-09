// Package main provides a small CLI harness to discover and run tests for
// individual coding challenges contained under the ./challenges directory.
//
// Usage examples:
//   go run ./cmd/harness -list
//   go run ./cmd/harness -challenge two-sum -v
//   go run ./cmd/harness -challenge two-sum -run '^TestTwoSum/basic$'
//   go run ./cmd/harness -all -- -race -cover
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// main parses flags, discovers requested challenge packages, and shells out to
// `go test` with the appropriate arguments. Extra arguments after "--" are
// forwarded unchanged to `go test`.
func main() {
	// Split out passthrough args after "--" so we can forward them to `go test`.
	raw := os.Args[1:]
	passthrough := []string{}
	if i := indexOf("--", raw); i >= 0 {
		passthrough = append(passthrough, raw[i+1:]...)
		raw = raw[:i]
	}

	fs := flag.NewFlagSet("harness", flag.ExitOnError)
	list := fs.Bool("list", false, "list challenges")
	all := fs.Bool("all", false, "run all challenge tests")
	challenge := fs.String("challenge", "", "challenge slug (directory under challenges/)")

	// Common go test flags we explicitly support (in addition to passthrough).
	run := fs.String("run", "", "regex for go test -run")
	bench := fs.String("bench", "", "regex for go test -bench")
	v := fs.Bool("v", false, "verbose")
	race := fs.Bool("race", false, "race detector")
	cover := fs.Bool("cover", false, "coverage")
	count := fs.Int("count", 1, "go test -count")
	timeout := fs.String("timeout", "", "go test -timeout")
	failfast := fs.Bool("failfast", false, "fail fast")

	_ = fs.Parse(raw)

	if *list {
		slugs, err := listChallenges("challenges")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, s := range slugs {
			fmt.Println(s)
		}
		return
	}

	if *all && *challenge != "" {
		fmt.Fprintln(os.Stderr, "cannot use -all with -challenge")
		os.Exit(2)
	}

	var slugs []string
	if *all {
		var err error
		slugs, err = listChallenges("challenges")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else if *challenge != "" {
		slugs = []string{*challenge}
	} else {
		fmt.Fprintln(os.Stderr, "usage: -list or -all or -challenge <slug>")
		os.Exit(2)
	}

	// Build the base `go test` arguments once, then reuse per challenge.
	baseArgs := []string{"test"}
	if *v {
		baseArgs = append(baseArgs, "-v")
	}
	if *race {
		baseArgs = append(baseArgs, "-race")
	}
	if *cover {
		baseArgs = append(baseArgs, "-cover")
	}
	if *failfast {
		baseArgs = append(baseArgs, "-failfast")
	}
	if *count != 1 {
		baseArgs = append(baseArgs, "-count", fmt.Sprint(*count))
	}
	if *timeout != "" {
		baseArgs = append(baseArgs, "-timeout", *timeout)
	}
	if *run != "" {
		baseArgs = append(baseArgs, "-run", *run)
	}
	if *bench != "" {
		baseArgs = append(baseArgs, "-bench", *bench)
	}

	// Execute tests for each selected challenge. Non-zero exit if any fail.
	overallFail := 0
	for _, slug := range slugs {
		pkg := "./challenges/" + slug
		fmt.Printf("==> %s\n", pkg)

		args := append([]string{}, baseArgs...)
		args = append(args, pkg)
		args = append(args, passthrough...)

		cmd := exec.Command("go", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = os.Environ()
		if err := cmd.Run(); err != nil {
			overallFail++
		}
	}

	if overallFail > 0 {
		os.Exit(1)
	}
}

// listChallenges returns the set of immediate subdirectory names under root
// (excluding dot-prefixed directories), sorted alphabetically.
func listChallenges(root string) ([]string, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("challenges directory not found: %s", root)
		}
		return nil, err
	}
	slugs := []string{}
	for _, e := range entries {
		if e.IsDir() && !strings.HasPrefix(e.Name(), ".") {
			slugs = append(slugs, e.Name())
		}
	}
	sort.Strings(slugs)
	return slugs, nil
}

// indexOf returns the index of needle in hay, or -1 if not present.
func indexOf(needle string, hay []string) int {
	for i, s := range hay {
		if s == needle {
			return i
		}
	}
	return -1
}
