package tools

func Max(data ...int) int {
	max := 0
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

func Pow(x, y int) int {
	pow := 1
	count := y
	for count != 0 {
		pow *= x
		count--
	}
	return pow
}
