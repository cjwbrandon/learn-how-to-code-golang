package main

import (
	"errors"
	"fmt"
	"log"
)

// Errors with info
// Using Struct to add to Error inteface
type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

// Assigning to a variable
var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	// Example 1
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Example 1
func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("norgate math: square root of negative number")
		// return 0, ErrNorgateMath
		nme := fmt.Errorf("norgate math: square root of negative number") // Returns an error
		return 0, norgateMathError{"12 N", "91 W", nme}
	}
	return 42, nil
}
