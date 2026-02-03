package mathematic

func Sum(a ...int) (sum int) {
	for _, v := range a {
		sum = sum + v
	}

	return sum
}
