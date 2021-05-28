package mymath

import (
	"testing"
	"fmt"
)

func TestAverage(t *testing.T) {
	// Code test in here ...

	r := Average(1, 2)
	if r != 1.5 {
		// Print the error log and mark this test as failed
		t.Error("Expected 1.5, but got", r)
	}
}

// Repetitive test cases, can use idiomatic *table-driven* style
type test struct {
	inputs []int // input args of IntMax
	want   int   // expected output of returning IntMax
}

func TestIntMax(t *testing.T) {
	// Code test in here ...
	
	tests := []test{
		{[]int{2001, 2000, 100, 101}, 2001},
		{[]int{5, 6, 3, 2}, 6},
		{[]int{89, 100, 70, 1}, 100},
	}

	for _, v := range tests {
		r := IntMax(v.inputs...)
		if r != v.want {
			// t.Error - report test failure (log and mark) but continue execution
			// t.Errorf("Expected %v, but got %v", v.want, r)

			// t.Fatal - report test failure (log and mark) but stop the execution (called runtime.Goexit)
			t.Fatalf("Expected %v, but got %v", v.want, r)
		}
	}

	// will not run when t.Fatal
	defer func() {
		t.Log("Closing testing!")
	}()
	t.Log("Done testing!")
}

// runtime.Goexit ?
// Only terminate the goroutine that calls it - defer will still run, but different between panic
// Others goroutines still continue execution

func TestIntMin(t *testing.T) {
	// Code test in here ...

	tests := []test{
		{[]int{2001, 2000, 100, 101}, 100},
		{[]int{5, 6, 3, 2}, 2},
		{[]int{89, 100, 70, 1}, 1},
	}

	for i, v := range tests {
		testName := fmt.Sprintf("Input %v : %v", (i + 1), v.inputs)
		// Create a subtest with name
		t.Run(testName, func(t *testing.T){
			r := IntMin(v.inputs...)
			if r != v.want {
				t.Fatalf("Expected %v, but got %v", v.want, r)
			}
		})
	}

	// will still run even when t.Fatal (caused different go-routine)
	t.Log("Done testing!")
}