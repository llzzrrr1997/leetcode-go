package problemset

func isWinner(player1 []int, player2 []int) int {
	double1 := 0
	double2 := 0
	point1 := 0
	point2 := 0
	n := len(player1)
	for i := 0; i < n; i++ {
		point1 += player1[i]
		point2 += player2[i]
		if double1 > 0 {
			point1 += player1[i]
			double1--
		}
		if double2 > 0 {
			point2 += player2[i]
			double2--
		}
		if player1[i] == 10 {
			double1 = 2
		}
		if player2[i] == 10 {
			double2 = 2
		}
	}
	if point1 > point2 {
		return 1
	}
	if point2 > point1 {
		return 2
	}
	return 0
}
