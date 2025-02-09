package search

func FindFibonacci(n uint) uint {
	if n < 2 {
		return n
	}
	
	return FindFibonacci(n-1) + FindFibonacci(n-2)
}
