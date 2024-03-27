package utility

func Add(x, y int) (res int) {
	return x + y
}

func Sign(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}
