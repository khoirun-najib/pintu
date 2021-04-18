package pintu

func RangeFactor(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		factorial := Factorial(i)
		if factorial == 6 {
			count++
		}
	}
	return count
}

func Factorial(n int) int {
	c := 1
	hb := n
	count := 0

	if n == 1 {
		return 1
	}

	for c < hb {
		hb = n / c
		if n%c == 0 {
			if c == hb {
				count++
			} else {
				count += 2
			}
		}
		if count > 6 {
			return 7
		}
		c++
	}
	return count
}
