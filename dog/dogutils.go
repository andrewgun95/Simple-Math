// Package dog is about a dog.
package dog

// DogYears is 1 human year equal to 7 dog years.
const (
	DogYears = 7
)

// Years take a human year and turn into dog years.
func Years(n int) int {
	return n * DogYears
}
