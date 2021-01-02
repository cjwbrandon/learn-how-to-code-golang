package main

import (
	"fmt"
	"runtime"
)

// Specifying type
var x int
var y float64

var z int8 = -128

// var z int8 = -129 // overflows

func main() {
	x = 42
	y = 42.34546354235234523523
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)

	// GO Operating System
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
