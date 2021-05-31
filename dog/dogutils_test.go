package dog

import (
	"testing"
)

func TestYears(t *testing.T) {
	expect := 35
	actual := Years(5)
	if actual != expect {
		t.Errorf("Expected %v, but got %v", expect, actual)
	}
}
