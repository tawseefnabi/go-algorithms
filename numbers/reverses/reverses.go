package reverses

func Reverse(number int) int {
	reverse := 0
	for number != 0 {
		reverse *= 10
		reverse += number % 10
		number /= 10
	}
	return reverse

}
