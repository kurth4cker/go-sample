package sample

func removeFactor(number, factor int) int {
	for number%factor == 0 {
		number /= factor
	}
	return number
}

func IsUgly(number int) bool {
	factors := [...]int{2, 3, 5}
	for _, factor := range factors {
		number = removeFactor(number, factor)
	}
	return number == 1
}

func nthUglyNumber(n int) int {
	count := 0
	number := 0
	for count != n {
		number++
		if IsUgly(number) {
			count++
		}
	}
	return number
}
