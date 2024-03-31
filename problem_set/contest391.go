package problemset

func sumOfTheDigitsOfHarshadNumber(x int) int {
	numList := make([]int, 0)
	t := x
	for t != 0 {
		numList = append(numList, t%10)
		t /= 10
	}
	sum := 0
	for _, v := range numList {
		sum += v
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}

func maxBottlesDrunk(numBottles int, numExchange int) int {
	fullBottles := numBottles
	emptyBottles := 0
	sum := 0
	for fullBottles != 0 {
		sum += fullBottles
		emptyBottles += fullBottles
		fullBottles = 0
		if emptyBottles >= numExchange {
			fullBottles++
			emptyBottles = emptyBottles - numExchange
			numExchange++
		}
	}
	return sum
}
