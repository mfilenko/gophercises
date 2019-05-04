package arrays

func Equilibrium(arr []int) int {
	sum := 0
	leftsum := 0

	for _, el := range arr {
		sum += el
	}

	for i, el := range arr {
		sum -= el

		if leftsum == sum {
			return i
		}

		leftsum += el
	}

	return -1
}
