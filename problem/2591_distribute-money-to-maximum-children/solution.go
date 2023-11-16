package main

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	if money < children+7 {
		return 0
	}
	if children == 2 {
		switch money {
		case 12:
			return 0
		case 16:
			return 2
		default:
			return 1
		}
	}

	return distMoney(money-8, children-1) + 1
}

func distMoney2(money int, children int) int {
	if money < children {
		return -1
	}
	if money < children+7 {
		return 0
	}

	var dfs func(money int, children int) int
	dfs = func(money int, children int) int {
		if children == 1 {
			if money <= 0 || money == 4 {
				return -1
			}
			if money == 8 {
				return 1
			}
			return 0
		}

		return dfs(money-8, children-1) + 1
	}

	return dfs(money, children)
}
