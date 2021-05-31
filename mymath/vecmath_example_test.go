package mymath

import (
	"fmt"
	"sort"
)

type ByMag []*Vector

func (a ByMag) Len() int           { return len(a) }
func (a ByMag) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMag) Less(i, j int) bool { return a[i].Magnitude() < a[j].Magnitude() }

func ExampleVector() { // Attach to Type
	vectors := []*Vector{
		&Vector{1, 3, 4},
		&Vector{5, 6, 7},
		&Vector{1, 1, 1},
	}

	// Before Sorting
	fmt.Println(vectors)
	sort.Sort(ByMag(vectors))
	// After  Sorting
	fmt.Println(vectors)

	// Output:
	// [(1,3,4) (5,6,7) (1,1,1)]
	// [(1,1,1) (1,3,4) (5,6,7)]
}
