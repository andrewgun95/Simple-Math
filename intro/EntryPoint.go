package main

import (
	"fmt"
	"github.com/andrewgun95/Simple-Math/dog"
)

func main() {
	// A. Unit Testing

	// 1. What is it ?
	// (i)  Unit is a small of logic of unit of the program
	// (ii) Testing is a way to *ensure code quality and reliability*

	// 1.b It's important part after writing your program !

	// 1.c Command :
	// go test path
	// * will look for any tests in any of the files (*_test.go) and run of them

	// 1.d Best Practice : The same package between code and tests code
	
	// 1.e Must start with *Test* followed by package-level declaration (function, type, const, etc) 
	//	   and received an argument of type *testing.T

	// 1.f For Ex :
	// func Average() {}
	// func TestAverage() {}

	// 2. Repetitive test case can used idiomatic *table driven style*

	// 3. Test By Example
	// Snippets of Go code - displayed as package documentation and verified by running as a test

	// 3.a It's executable documentation ?
	// Guarantee the information *documentation will not out of date*, caused :
	// (i)  Compiled
	// (ii) Tested

	// 3.b Must start with *Example* followed by package-level declaration (function, type, const, etc) 
	//     and take no argument

	// 3.c For Ex :
	// In example_test.go :
	// func Example() {}
	// func ExampleAverage() {}
	// func ExampleAverage_2() {}
	// func ExampleAverage_third() {}
	// * Multiple example by adding a suffix in the end of identifier follow by underscore

	// 3.d Behind Execute Example
	// Example can be executed as a Test and it will,
	// capture data written to standard output and then compare the output againts the example "Output:" comment

	// B. Writing a Document

	// 1. What is it ?
	// (i) Documentation is about a way to *ensure code accessible and maintainable*

	// 1.a Ideally, should be coupled to the code itself - documentation evolve along with the code !

	// 1.b Command :
	// go doc - standard command for read documentation
	// godoc  - documentation tools, ex : can launch a local web server to access the doc

	// ---------------------------
	// 			godoc
	// ---------------------------

	// 2. Parse Go source code with comments and produce documentation as HTML or plain text
	// It's similar to *javadoc* in Java or *docstring* in Python

	// Convention :
	//
	// 2.a Insert a documentation to a type, variable, function, constant, or package
	//    by directly give a comment on it's declaration - and must begin with the identifier
	//
	// For Ex :
	//
	// // xyz is a function .
	// func xyz() {}
	//
	// // abc is a variable .
	// var abc int
	//
	// // Package math
	// package math

	// 2.b Insert a documentation on a package declaration for general information - (what is the package about)
	//	  and for details information about the package can added in doc.go
	//	  First sentence of comment documentation will appear in godoc *package list*

	// 2.c Must adjacent to a top-level declaration

	// For Ex :

	// Wrong !

	// func xyz() {} // xyz is a function .

	// 2.d Summary :
	// Godoc's minimal approach is how easy it is to use.
	// As a result, a lot of Go code, including all of the standard library, already follows the conventions.

	// 2.e Exercise :
	fmt.Printf("Human %v years is equal to dog %v years", 5, dog.Years(5))

	// C. Benchmarking
	// 1. What is it ?
	// (i) Benchmarking is a way to *ensure code performance*

	// 1.a It's important that performance tweaking should typically be done after the system is up and running !
	// Remember : Premature optimization is the root of the evil - Donald Knuth

	// 1.b Ideally conjunction with your standard unit tests - BET (Benchmarking, Example, and Test)

	// 1.c Must start with *Benchmark* followed by package-level declaration (function, type, const, etc) 
	//	   and received an argument of type *testing.B

	// 1.d Syntax :
	// go test path -bench .
	// * . current package

	// go test path -run Bench -bench .
	// * will test only performance issues

	// go test path -run filter -bench .
	// * will test according to filter
	// For Ex :
	// go test path -run Average -bench .

	// 1.e In real life, your program will accept a variety of distinct inputs
	
}

// More Example :
// Revisit - Difference between godoc.org and golang.org
// godoc.org is standard and third party library
// golang.org is only standard library
