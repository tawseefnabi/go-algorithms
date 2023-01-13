package palindromes

import "github.com/tawseefnabi/go-algorithms/numbers/reverses"

// IsPalindrome determines if the input is a palindrome
func IsPalindrome(number int) bool {
	return number == reverses.Reverse(number)
}
