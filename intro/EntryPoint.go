package main

import (
	"fmt"
	"github.com/andrewgun95/Simple-Math/dog"
)

func main() {
	// Unit Testing

	// What is it ?
	// 1. Unit is a small of logic of unit of the program
	// 2. Testing is a way to *ensure code quality and reliability*

	// It's important part after writing your program !

	// Command :
	// go test path
	// * will look for any tests in any of the files (*_test.go) and run of them

	// Best Practice :
	// 1. The same package between code and tests code
	// 2. Start with *Test* followed by function name and received an argument of type *testing.T

	// For Ex :
	// func Average() {}
	// func TestAverage() {}

	// 3. Repetitive test case can used idiomatic *table driven style*

	// 4. Test By Example

	// Execute Example
	// Capture data written to standard output and then compare the output againts the example "Output:" comment

	// Writing a Document

	// What is it ?
	// 1. Documentation is about a way to *ensure code accessible and maintainable*

	// Ideally, should be coupled to the code itself - documentation evolve along with the code

	// Command :
	// go doc - standard command for read documentation
	// godoc  - documentation tools, ex : can launch a local web server to access the doc

	// ---------------------------
	// 			godoc
	// ---------------------------

	// Parse Go source code with comments and produce documentation as HTML or plain text
	// It's similar to *javadoc* in Java or *docstring* in Python

	// Convention :
	//
	// 1. Insert a documentation to a type, variable, function, constant, or package
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

	// 2. Insert a documentation on a package declaration for general information - (what is the package about)
	//	  and for details information about the package can added in doc.go
	//	  First sentence of comment documentation will appear in godoc *package list*

	// 3. Must adjacent to a top-level declaration

	// For Ex :

	// Wrong !

	// func xyz() {} // xyz is a function .

	// Summary :
	// Godoc's minimal approach is how easy it is to use.
	// As a result, a lot of Go code, including all of the standard library, already follows the conventions.

	// Exercise :
	fmt.Printf("Human %v years is equal to dog %v years", 5, dog.Years(5))
}

// More Example :
// Revisit - Difference between godoc.org and golang.org
// godoc.org is standard and third party library
// golang.org is only standard library
