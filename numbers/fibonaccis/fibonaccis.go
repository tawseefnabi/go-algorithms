package fibonaccis

//Fibonacci returns fibonacci number
// The Fibonacci numbers are the numbers in the following integer sequence.
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ……..

func Fibonacci(number int) int {
	if number <= 0 {
		return 0
	}
	n1 := 0
	n2 := 0
	current := 1
	for i := 1; i < number; i++ {
		n2 = n1
		n1 = current
		current = n2 + n1
	}
	return current
}
