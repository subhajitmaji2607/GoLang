/*
	The var statement declares a list of variables;
	A var statement can be at package or function level

	A var declaration can include initializers, one per variable.
	If an initializer is present, the type can be omitted;
	the variable will take the type of the initializer.
	Variables declared without an explicit initial value are given their zero value.
	The zero value is:
	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.


	Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available

	Go's basic types are:
	----------------------------
	bool,string,
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32
		// represents a Unicode code point
	float32 float64
	complex64 complex128
	------------------------------

	Constants are declared like variables, but with the const keyword.
	*Constants cannot be declared using the := syntax.
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var c, java, python bool
var j, k int = 1, 2

func main() {
	var i int
	var javaScript = "awesome"
	rust := "new in the town"
	const constant = 10

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	var typeConversion int = 42
	var convertedType float64 = float64(typeConversion)

	fmt.Println(i, c, java, python)
	fmt.Println(j, k, javaScript, rust)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println(convertedType)
	fmt.Println("constant: %v\n", constant)
}
