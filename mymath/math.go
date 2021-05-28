// Package math provides a basic mathematical functions, such as : calculate average, sum, find min and max, or etc.
package mymath

// Average take a numbers as arguments and calculate the average of that numbers.
func Average(n ...float64) float64 {
	s := len(n)
	var r float64
	for _, v := range n {
		r += v
	}
	return r / float64(s)
}

// IntMax take a numbers as arguments and find the max number of that numbers.
func IntMax(n ...int) int {
	max := -1
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return max
}

// IntMin take a numbers as arguments and find the min number of that numbers.
//
// BUG(who): IntMin(n) when all the numbers is equally the same will return the first of that numbers 
func IntMin(n ...int) int {
	min := n[0]
	for _, v := range n {
		if v < min {
			min = v
		}
	}
	return min
}

// Sum take a numbers as arguments and calculate the sum of that numbers.
func Sum(n ...float64) float64 {
	var r float64
	for _, v := range n {
		r += v
	}
	return r
}

// PlusTwo return a number, added by two.
//
// Deprecated:This function is no longer neccessary, used basic operator instead.
func PlusTwo(n int) int {
	return n + 2
}

// Vector a type represent a point in cartisian coordinate with x, y, z value
type Vector struct {
	x, y, z int
}

// Add operate addition of a value to the vector
func (v *Vector) Add(value Vector) {
	v.x += value.x
	v.y += value.y
	v.z += value.z
}