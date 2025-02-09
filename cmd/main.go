package main

import (
	"algo/internal/search"
	"fmt"
)

func main() {
	//var n int
	//arr := rand.Perm(20)
	//fmt.Println(arr)
	//fmt.Println(misc.EvenOrOddInArr(arr...))
	//q := misc.QuarterOfYear(10)
	//fmt.Println(q)
	//fmt.Println(misc.EvenOrOdd(q))
	//
	//fmt.Println("Enter a number to check if it is prime..")
	//_, err := fmt.Scan(&n)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(misc.IsPrime(n))
	//fmt.Println("Enter a number to check if it's in a slice or not")
	//if _, err := fmt.Scan(&n); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(search.BinarySearch(n, slices.Sort(arr)))
	fmt.Println(search.FindFibonacci(50))
}
