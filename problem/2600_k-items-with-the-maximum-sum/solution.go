package main

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	sub := k - numOnes
	if sub <= 0 {
		return k
	}

	if sub <= numZeros {
		return numOnes
	}

	return numOnes - (sub - numZeros)
}
