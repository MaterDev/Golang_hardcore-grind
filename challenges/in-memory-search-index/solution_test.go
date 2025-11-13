package inmemorysearchindex

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestIndex_Basic(t *testing.T) {
	var idx Index
	idx.Add("doc1", "Go is fun and fast")
	idx.Add("doc2", "Go go run")
	idx.Add("doc3", "Rust and Go")

	got := idx.Search("go")
	want := []string{"doc1", "doc2", "doc3"}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("go: got %v want %v", got, want)
	}

	got = idx.Search("and")
	want = []string{"doc1", "doc3"}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("and: got %v want %v", got, want)
	}
}

func TestIndex_NoDuplicates(t *testing.T) {
	var idx Index
	idx.Add("docX", "go go go go")
	got := idx.Search("go")
	if !reflect.DeepEqual(got, []string{"docX"}) {
		t.Fatalf("dedup: got %v want [docX]", got)
	}
}

func TestIndex_CaseInsensitive(t *testing.T) {
	var idx Index
	idx.Add("doc1", "Go")
	idx.Add("doc2", "gO")
	idx.Add("doc3", "rust")
	got := idx.Search("GO")
	sort.Strings(got)
	if !reflect.DeepEqual(got, []string{"doc1", "doc2"}) {
		t.Fatalf("case: got %v want [doc1 doc2]", got)
	}
}

func TestIndex_UnknownTerm(t *testing.T) {
	var idx Index
	idx.Add("doc1", "go is great")
	if got := idx.Search("python"); len(got) != 0 {
		t.Fatalf("unknown term: got %v want []", got)
	}
}

func BenchmarkIndex_Add(b *testing.B) {
	var idx Index
	text := strings.Repeat("go rust ", 20)
	for i := 0; i < b.N; i++ {
		idx.Add("doc", text)
	}
}

func BenchmarkIndex_Search(b *testing.B) {
	var idx Index
	idx.Add("d1", "go go go go go")
	idx.Add("d2", "go rust go")
	for i := 0; i < b.N; i++ {
		_ = idx.Search("go")
	}
}
