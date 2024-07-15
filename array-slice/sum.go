package arrayslice

func Sum(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum
}

func SumAll(arrToSum ...[]int) []int {
	var sums []int
	for _, numbers := range arrToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(arrToSum ...[]int) []int {
	var sums []int
	for _, numbers := range arrToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
