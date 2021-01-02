package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	x = append(x, 11)
	fmt.Println(x)
	x = append(x, 12, 13, 14)
	fmt.Println(x)

	y := []int{15, 16, 17, 18, 19}
	x = append(x, y...)
	fmt.Println(x)
}
