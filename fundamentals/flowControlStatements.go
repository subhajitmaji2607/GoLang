/*
Go has only one looping construct, the for loop.
*/
package main

import (
	"fmt"
	"math"
	"runtime"
)

// If statements
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Like for, the if statement can start with a short statement to execute before the condition
// first it evaluate the value of "v" then compare the value of "v" with lim
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//The init and post statements are optional.(behave like while loop)
	sub := 11
	for sub > 10 {
		sub = sub - sub
	}
	fmt.Println(sub)

	//If statement example
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//Switch case
	fmt.Print("Go runs on ")
	// os := runtime.GOOS
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	/*A defer statement defers the execution of a function until the surrounding function returns.
	The deferred call's arguments are evaluated immediately,
	but the function call is not executed until the surrounding function returns.
	https://go.dev/blog/defer-panic-and-recover
	*/
	defer fmt.Println("world")

	fmt.Println("hello")
}
