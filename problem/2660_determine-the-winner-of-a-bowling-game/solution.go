package main

func isWinner(player1 []int, player2 []int) int {
	s1, s2 := score(player1), score(player2)
	if s1 > s2 {
		return 1
	} else if s1 < s2 {
		return 2
	} else {
		return 0
	}
}

func score(player []int) int {
	sum, flag := 0, 0
	for _, v := range player {
		if flag > 0 {
			flag--
			sum += 2 * v
		} else {
			sum += v
		}

		if v == 10 {
			flag = 2
		}
	}
	return sum
}

func score2(player []int) int {
	sum := 0
	for i, v := range player {
		if (i >= 1 && player[i-1] == 10) || (i >= 2 && player[i-2] == 10) {
			sum += 2 * v
		} else {
			sum += v
		}
	}
	return sum
}
