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
func isPalindrome(s string) bool {

	// Create two numbers to track the indexes at beginning/end of string
	i, j := 0, len(s) - 1

	// will loop while the two indexes dont overlap. Will exclude loops for odd numbered words since both indexes at tdhe middle character are ok.
	for i < j {
		// skip non-alphanumeric characters on the left
		for i < j && !isAlnumASCII(s[i]) {
			i++
		}
		// skip non-alphanumeric characters on the right
		for i < j && !isAlnumASCII(s[j]) {
			j--
		}
		
		// if the indexes are equal or cross eachother, then break out (not a palindrome)
		if i >= j {
			break
		}
		
		// Case-insensitive comparison using ASCII lowercasing. (Comparing the lower case version of both characters.)
		if toLowerASCII(s[i]) != toLowerASCII(s[j]) {
			return false
		}

		// move the indexes in towards eachother.
		i++
		j--
	}

	return true

}

// Report if  b is ASCII letter or digit
func isAlnumASCII(b byte) bool {
	
	// Return true if the character is including or between a-z, A-Z, 0-9
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}

func toLowerASCII(b byte) byte {

	// If the character is a capital A-Z, return the lowercase
	if b >= 'A' && b <= 'Z' {

		// 'a' - "A" is 32, meaning add 32 to a unicode number to turn it in to the corresponding capital. This express describes an offset, without a magic number.
		return b + ('a' - 'A')
	}
	return b
}
