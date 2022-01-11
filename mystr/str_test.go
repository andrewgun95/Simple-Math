package mystr

import (
	"testing"
	"fmt"
)

func TestJoin(t *testing.T) {
	actual := Join([]string{"hello", "andreas"})
	expect := "hello andreas"
	if actual != expect {
		t.Errorf("Expected output %v, but got %v", expect, actual)
	}
}

func TestCat(t *testing.T) {
	actual := Cat([]string{"hello", "andreas"})
	expect := "hello andreas"
	if actual != expect {
		t.Errorf("Expected output %v, but got %v", expect, actual)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i:=0; i < b.N; i++ {
		Join([]string{"hello", "andreas"})
	}
}

func BenchmarkCat(b *testing.B) {
	for i:=0; i < b.N; i++ {
		Cat([]string{"hello", "andreas"})
	}
}

func ExampleJoin() {
	r := Join([]string{"abc", "def", "ghi"})
	fmt.Println(r)
	// Output: abc def ghi
}

func ExampleCat() {
	r := Cat([]string{"abc", "def", "ghi"})
	fmt.Println(r)
	// Output: abc def ghi
}