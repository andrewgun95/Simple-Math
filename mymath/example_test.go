package mymath

import (
	"fmt"
)

func ExampleAverage() { // Attach to Function
	// Success
	fmt.Println(Average(1, 2))
	// Output: 1.5
}

func ExampleAverage_2() { // Attach to Function
	// Failed
	fmt.Println(Average(2, 4))
	// Output: 2.5
}

func ExampleAverage_3() { // Attach to Function
	// Success
	fmt.Println(Average(3, 5))
	// Output:
	// 4
}

func ExampleAverage_4() { // Attach to Function
	// Success
	fmt.Println(Average(4, 4))
	// Output:
	//
	// 4
}

func ExampleAverage_5() { // Attach to Function
	// Compiled, but Not Executed
	fmt.Println(Average(5, 6))
}

func ExampleVector_Add() { // Attach to Type_Method
	v1 := &Vector{1, 2, 3}
	v1.Add(&Vector{3, 2, 1})
	fmt.Println(*v1)
	// Output: {4 4 4}
}
