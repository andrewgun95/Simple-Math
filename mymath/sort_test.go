package mymath

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s:%d", p.Name, p.Age)
}

// ByName implements sort.Interface for []Person based on
// the Name field.
type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// It's a whole file example
func Example() { // Attach to Package
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

	// Output:
	// [Bob:31 John:42 Michael:17 Jenny:26]
	// [Bob:31 Jenny:26 John:42 Michael:17]
}
