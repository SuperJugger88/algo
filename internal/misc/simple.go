package misc

import "math"

func EvenOrOdd(num int) bool {
	return num%2 == 0
}

func EvenOrOddInArr(arr ...int) []bool {
	res := make([]bool, len(arr), len(arr))

	for i, num := range arr {
		res[i] = num%2 == 0
	}
	return res
}

func QuarterOfYear(month int) int {
	return int(math.Ceil(float64(month) / 3))
}
