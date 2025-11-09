// Package twosum provides the Two Sum challenge implementation point.
//
// Problem: Given an integer slice nums and an integer target, return the
// indices of the two numbers such that they add up to target.
//
// Requirements
// - Return a slice of exactly two indices [i, j] such that nums[i] + nums[j] == target.
// - If no such pair exists, return nil (or an empty slice). The tests accept either.
// - Indices are zero-based.
// - Return indices in ascending order (i < j).
// - The input may contain negative numbers and duplicate values.
//
// Examples
//   nums = [2, 7, 11, 15], target = 9    => [0, 1]
//   nums = [3, 2, 4], target = 6         => [1, 2]
//   nums = [-1, -2, -3, -4, -5], -8      => [2, 4]
//   nums = [1, 2, 3], target = 7         => nil
//
// Approach Hints
// - Aim for O(n) time using a hash map that tracks values to indices as you scan.
// - When at index i with value x, look for target - x in the map; if present at j,
//   return [min(i, j), max(i, j)]. Otherwise record x -> i and continue.
// - Keep memory usage O(n).
//
// Running tests
// - All tests for this challenge live in solution_test.go.
// - From repo root you can run just this challenge via the harness:
//     ./scripts/run.sh -challenge two-sum -v
// - Run a specific subtest with:
//     ./scripts/run.sh -challenge two-sum -run '^TestTwoSum/negatives$'
package twosum

// TwoSum returns indices of two numbers in nums that add up to target.
// If no such pair exists, it returns nil. The returned indices are in
// ascending order (i < j).
func TwoSum(nums []int, target int) []int {
	// TODO: Implement using an O(n) hash map approach as described above.
	// Remember to return indices in ascending order. Return nil if no pair exists.
	return nil
}
