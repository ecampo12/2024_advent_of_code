package utils

import (
	"fmt"
	"reflect"
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

// MapContains is a function that takes a map and a value and returns true if the value is in the map
func MapContainsValue[t comparable, u any](m map[t]u, value u) bool {
	for _, v := range m {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Repeat is a function that takes a value and a number and returns a slice of the value repeated n times
func Repeat[t any](m t, n int) []t {
	x := Abs(n)
	res := make([]t, x)
	for i := range res {
		res[i] = m
	}
	return res
}

// Int is a function that takes a string and returns the integer value of that string
func Int(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

// SliceMin is a function that takes a slice of values and returns the minimum value depending on the given function
func SliceMin[t any](s []t, f func(t) int) t {
	min := s[0]
	for _, v := range s {
		if f(v) < f(min) {
			min = v
		}
	}
	return min
}

// SliceMax is a function that takes a slice of values and returns the maximum value depending on the given function
func SliceMax[t any](s []t, f func(t) int) t {
	max := s[0]
	for _, v := range s {
		if f(v) > f(max) {
			max = v
		}
	}
	return max
}

// Apply is a function that takes a slice of values and a function and returns a slice of the values after applying the function
func Apply[t any, u any](s []t, f func(t) u) []u {
	res := make([]u, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}

// ApplySum is a function that takes a slice of values and a function and returns the sum of the values after applying the function
func ApplySum[t any](s []t, f func(t) int) int {
	sum := 0
	for _, v := range s {
		sum += f(v)
	}
	return sum
}
