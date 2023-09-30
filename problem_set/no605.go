package problemset

// 贪心算法，能种就种，i种的条件就是 i-1, i,i+1 都为0
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if (i == 0 || flowerbed[i-1] == 0) && flowerbed[i] == 0 && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
