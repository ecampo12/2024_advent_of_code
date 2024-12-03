package utils

import (
	"fmt"
	"time"
)

// Timer is a function that takes a function and returns the time it took to run that function
func Timer(f func(), funcName string) {
	start := time.Now()
	f()

	fmt.Printf("%s took: %v\n", funcName, time.Since(start))
}

// Abs is a function that takes an integer and returns its absolute value
// the standard library does not have a function for integer absolute value *facepalm*
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Sum is a function that takes a slice of integers and returns the sum of all the elements
func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Counter is a function that takes a slice of any type and returns a map of the count of each element
func Counter[t comparable](nums []t) map[t]int {
	m := make(map[t]int)
	for _, n := range nums {
		m[n]++
	}
	return m
}
