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

import (
	"fmt"
)

// TwoSum returns indices of two numbers in nums that add up to target.
// If no such pair exists, it returns nil. The returned indices are in
// ascending order (i < j).

func TwoSum(nums []int, target int) []int {
	fmt.Printf("💧 Incoming Numbers: %v \n🎯 target: %d \n", nums, target)

	// Use a hash map to track index-values that have already been seen.
	seen := make(map[int]int, len(nums))	
	

	// Iterate over numbers and check if the value at the current index has a match within the seen-map. Looking for a match for target - x, which is the required pair number.
	for i, x := range nums {
		fmt.Printf("\t⚡️ Matching for: %d \n", x)
		fmt.Printf("\t🔥 Seen: %d \n", seen)

		if j, ok := seen[target - x]; ok {
			fmt.Printf("Found match at index: %d \n", j)
			return []int{j, i}
		}
		
		// This will over ride and store only one index for the same value. So if the value is a match for the current number will only show the most recently seen one.
			// ! This is actually a good thing because it reduces space complexity.
		seen[x] = i

	}

	return nil
}


// ! My attempt O(n^2)
// func TwoSum(nums []int, target int) []int {
// 	fmt.Printf("🔥 TwoSum running with Nums: %v, Target: %v \n", nums, target)

// 	for i1, v1 := range nums {
// 		fmt.Printf("Matching for %v \n", v1)
// 		currentSlice := nums[i1 + 1:]

// 			for i2, v2 := range nums[i1 + 1:] {
// 				fmt.Printf("💧current slice: %v , at i2: %v \n", currentSlice, i2)

// 				originalIndexMatcher := i1 + i2 + 1
// 				fmt.Printf("⚡️ original index of v2: %v, in  %v ... is: %v \n", v2, nums, originalIndexMatcher)
				
// 				if target - v1 == v2 {
// 					fmt.Printf("👉🏾 Found a match: %v \n", v2)

// 					result := []int{min(i1, originalIndexMatcher), max(i1, originalIndexMatcher)}
// 					fmt.Printf("returning slice: %v", result)

// 					return result
// 				}
// 			}
// 	}

// 	return nil
// }
