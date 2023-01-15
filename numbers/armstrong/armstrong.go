package armstrong

import (
	"math"

	"github.com/tawseefnabi/go-interview/numbers/countdigits"
)

// IsArmstrong determines if a number is an armstrong number
// Given a number x, determine whether the given number is Armstrong
// number or not. An integer of n digits is called an Armstrong number if
// abcd... = pow(a,n) + pow(b,n) + pow(c,n) + pow(d,n) + ....
// Input : 153
// Output : Yes
// 153 is an Armstrong number.
// 1*1*1 + 5*5*5 + 3*3*3 = 153
func IsArmstrong(number int) bool {
	digits := float64(countdigits.CountDigits(number))
	current := number
	sum := 0
	for current != 0 {
		sum += int(math.Pow(float64(current%10), digits))
		current /= 10
	}
	return sum == number
}
