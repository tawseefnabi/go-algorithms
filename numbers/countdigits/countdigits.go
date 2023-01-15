package countdigits

// CountDigits counts the number of digits
func Countdigits(number int) int {
	count := 0
	if number == 0 {
		return 1
	}
	for number != 0 {
		number /= 10
		count++
	}
	return count
}
