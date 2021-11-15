package problemset

import "math"

//求奇数的质因数
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}
