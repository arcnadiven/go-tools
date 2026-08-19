package tools

const (
	baseStr    = "abcdefghijklmnopqrstuvwxyz"
	diskPrefix = "/dev/sd"
)

func AllocateDisk(n int) string {
	return diskPrefix + generate(n)
}

func generate(n int) string {
	x, min, max, result := 1, 0, len(baseStr), ""
	for {
		if n >= min && n < max {
			n -= min
			e := power(len(baseStr), x-1)
			for x != 0 {
				result += string(baseStr[n/e])
				n %= e
				e /= len(baseStr)
				x--
			}
			break
		}
		min = max
		x++
		max += power(len(baseStr), x)
	}
	return result
}

func power(x, y int) int {
	p := 1
	for y > 0 {
		p *= x
		y--
	}
	return p
}
