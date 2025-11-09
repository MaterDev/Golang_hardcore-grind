// Package twosum contains tests for the classic Two Sum problem. Implement the
// function in solution.go so these tests pass.
package twosum

import (
	"reflect"
	"testing"
)

// TestTwoSum verifies the expected indices for various input arrays and targets.
func TestTwoSum(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		got := TwoSum([]int{2, 7, 11, 15}, 9)
		want := []int{0, 1}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
	
	t.Run("duplicates", func(t *testing.T) {
		got := TwoSum([]int{3, 2, 4}, 6)
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
	
	t.Run("negatives", func(t *testing.T) {
		got := TwoSum([]int{-1, -2, -3, -4, -5}, -8)
		want := []int{2, 4}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
	
	t.Run("same-value", func(t *testing.T) {
		got := TwoSum([]int{3, 3}, 6)
		want := []int{0, 1}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
	
	t.Run("no-solution", func(t *testing.T) {
		got := TwoSum([]int{1, 2, 3}, 7)
		if len(got) != 0 {
			t.Fatalf("expected empty when no solution, got %v", got)
		}
	})
}

// BenchmarkTwoSum runs a simple benchmark of TwoSum with a fixed workload.
func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 20, 1, 5, 3, 8, 12}
	for i := 0; i < b.N; i++ {
		_ = TwoSum(nums, 26)
	}
}
