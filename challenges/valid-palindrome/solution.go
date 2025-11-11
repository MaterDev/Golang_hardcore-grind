// Package validpalindrome provides the Valid Palindrome challenge implementation point.
//
// Problem: Given a string s, return true if it is a palindrome, or false otherwise,
// after converting all uppercase letters into lowercase and removing all
// non-alphanumeric characters.
//
// Requirements
// - Consider only letters and digits; ignore spaces, punctuation, and symbols.
// - Case-insensitive comparison (treat 'A' and 'a' the same).
// - Aim for O(n) time using a two-pointer technique and O(1) extra memory.
//
// Examples
//
//	s = "A man, a plan, a canal: Panama" => true
//	s = "race a car" => false
//	s = "" => true
//
// Approach Hints
//   - Use two indices i (from start) and j (from end), move inward while skipping
//     non-alphanumeric runes. Compare lowercase forms.
//   - In Go, `unicode.IsLetter`, `unicode.IsNumber`, and `unicode.ToLower` are
//     convenient helpers for classification and case-folding.
//
// Running tests
//   - All tests for this challenge live in solution_test.go.
//   - From repo root you can run just this challenge via the harness:
//     ./scripts/run.sh -challenge valid-palindrome -v
//   - Run a specific subtest with:
//     ./scripts/run.sh -challenge valid-palindrome -run '^TestIsPalindrome/mixed-case$'
package validpalindrome

// isPalindrome reports whether s is a palindrome after stripping non-alphanumeric
// characters and lowercasing. Implement this function to satisfy the tests.

// ! Unicode Optimized

import (
	"fmt"
	"unicode"
)

// isPalindrome returns true if s is a palindrome after:
// - keeping only letters/digits (Unicode-aware)
// - lowercasing using Unicode rules
// Then it two-pointers over the normalized []rune.
func isPalindrome(s string) bool {
	// 1) Normalize: filter to letters/digits and lowercase

	// create slice of runes, with 0 items, and capacity equal to the length of incoming string.
	buf := make([]rune, 0, len(s))

	// Looping over all characters in teh string.
	for _, r := range s {
		// If the character is a unicode letter or number, add to the buffer as lowercase.
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			buf = append(buf, unicode.ToLower(r))
		}
		fmt.Printf("🔥 Current Buffer: %v \n", buf)
	}

	// 2) Two-pointer over normalized runes
		// Two variables, starting at both ends of the slice.
	left, right := 0, len(buf)-1

	// While the pointers dont overlap. At the point where the structure has an odd number of elements they will arrive at the center element together and skip this for loop
	for left < right {
		// If the pointers are not equal, then return false.
		if buf[left] != buf[right] {
			return false
		}

		// Progress the pointers
		left++
		right--
	}
	return true
}

// ! Original Attempt
// func isPalindrome(s string) bool {

// 	// Create two numbers to track the indexes at beginning/end of string
// 	i, j := 0, len(s) - 1

// 	// will loop while the two indexes dont overlap. Will exclude loops for odd numbered words since both indexes at tdhe middle character are ok.
// 	for i < j {
// 		// skip non-alphanumeric characters on the left
// 		for i < j && !isAlnumASCII(s[i]) {
// 			i++
// 		}
// 		// skip non-alphanumeric characters on the right
// 		for i < j && !isAlnumASCII(s[j]) {
// 			j--
// 		}
		
// 		// if the indexes are equal or cross eachother, then break out (not a palindrome)
// 		if i >= j {
// 			break
// 		}
		
// 		// Case-insensitive comparison using ASCII lowercasing. (Comparing the lower case version of both characters.)
// 		if toLowerASCII(s[i]) != toLowerASCII(s[j]) {
// 			return false
// 		}

// 		// move the indexes in towards eachother.
// 		i++
// 		j--
// 	}

// 	return true

// }

// // Report if  b is ASCII letter or digit
// func isAlnumASCII(b byte) bool {
	
// 	// Return true if the character is including or between a-z, A-Z, 0-9
// 	return (b >= 'a' && b <= 'z') ||
// 		(b >= 'A' && b <= 'Z') ||
// 		(b >= '0' && b <= '9')
// }

// func toLowerASCII(b byte) byte {

// 	// If the character is a capital A-Z, return the lowercase
// 	if b >= 'A' && b <= 'Z' {

// 		// 'a' - "A" is 32, meaning add 32 to a unicode number to turn it in to the corresponding capital. This express describes an offset, without a magic number.
// 		return b + ('a' - 'A')
// 	}
// 	return b
// }
