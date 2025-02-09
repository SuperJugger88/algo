package misc

import "math"

func IsPrime(n int) bool {
	switch n < 2 {
	case true:
		return false
	default:
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}
