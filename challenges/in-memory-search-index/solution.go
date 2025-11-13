// Package inmemorysearchindex provides the In-Memory Search Index challenge implementation point.
//
// Problem: Build a simple in-memory inverted index with the ability to add documents
// and search for documents containing a given query term.
//
// Requirements
// - Define a struct to hold index state.
// - Implement methods on that struct to add documents and search.
// - The search should return a slice of document IDs that match the query.
// - Keep the implementation details up to you; tests validate observable behavior.
//
// Examples
//   // Pseudocode usage
//   var idx Index
//   idx.Add("doc1", "go is fun")
//   idx.Add("doc2", "go go run")
//   got := idx.Search("go") // expected to include doc1 and doc2 (order defined by tests)
//
// Approach Hints
// - Consider an inverted index mapping terms to document IDs.
// - Ensure no duplicate doc IDs are returned for a term.
//
// Running tests
// - All tests for this challenge live in solution_test.go.
// - From repo root you can run just this challenge via the harness:
//     ./scripts/run.sh -challenge in-memory-search-index -v
// - Run a specific subtest with:
//     ./scripts/run.sh -challenge in-memory-search-index -run '^TestIndex/basic$'
package inmemorysearchindex

import (
	"fmt"
	"strings"
)

const debugAdd = true
const debugSearch = true

func dbgAddf(format string, args ...interface{}) {
	if debugAdd {
		fmt.Printf(format, args...)
	}
}

func dbgSearchf(format string, args ...interface{}) {
	if debugSearch {
		fmt.Printf(format, args...)
	}
}

// Index stores any state needed for the in-memory search index.
type Index struct {
	// implement fields as needed
	inv map[string]map[string]struct{}
}

// Add ingests the given document into the index.
func (idx *Index) Add(docID string, text string) {

	// Check if inv is empty, if so then initialize outer map for zero-value of the index
	if idx.inv == nil {
		idx.inv = make(map[string]map[string]struct{})
	}

	tokens := strings.Fields(text)

	// If there are no fields, then will end the method call.
	if len(tokens) == 0 {
		dbgAddf("🟡 Add(%s): no tokens\n", docID)
		return
	}
	dbgAddf("🌈 Add(%s): tokens=%v | 📚 terms=%d\n", docID, tokens, len(idx.inv))

	// Seen to deduping
	seen := make(map[string]struct{})

	// for each token will check the seen terms map, if it is found then will continue
	for _, tok := range tokens {
		//make the token lowercase
		term := strings.ToLower(tok)

		// If the term exists in the seen map, then continue to the next iteration
		if _, ok := seen[term]; ok {
			dbgAddf("⏭️  skip %q (seen)\n", term)
			continue
		}

		// ! If the term does not exist in the map, then will create a new entry for the seen term inside of the map
		seen[term] = struct{}{} // This will set the term in the seen map, using a zero value to indicate that it exists.

		// Trying to find the term in the inv, if not there it will ad the term after makign a n empty struct
		m := idx.inv[term]

		if m == nil {
			m = make(map[string]struct{})
			idx.inv[term] = m
		}

		// 
		m[docID] = struct{}{}
		dbgAddf("➕ index[%s] ← %s\n", term, docID)
	}
	dbgAddf("✅ Add(%s): done | 📚 terms=%d\n\n", docID, len(idx.inv))
}

// Search returns the document IDs that match the given query.
func (idx *Index) Search(query string) []string {
	// TODO: implement query against the index
	return nil
}
