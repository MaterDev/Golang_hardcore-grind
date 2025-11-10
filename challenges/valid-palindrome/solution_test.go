// Package validpalindrome contains tests for the Valid Palindrome challenge.
// Implement the function in solution.go so these tests pass.
package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name string
		in   string
		ok   bool
	}{
		{"basic-panama", "A man, a plan, a canal: Panama", true},
		{"race-a-car", "race a car", false},
		{"empty", "", true},
		{"only-punct", ".,,", true},
		{"numeric", "1221", true},
		{"mixed-case", "No 'x' in Nixon", true},
		{"underscore-ignored", "ab_a", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.in)
			if got != tc.ok {
				t.Fatalf("%q: got %v want %v", tc.in, got, tc.ok)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	in := "A man, a plan, a canal: Panama"
	for i := 0; i < b.N; i++ {
		_ = isPalindrome(in)
	}
}
