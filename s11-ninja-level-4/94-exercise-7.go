package main

import "fmt"

func main() {
	x := []string{"james", "bond"}
	y := []string{"Miss", "Moneypenny"}
	z := [][]string{x, y}

	for i, v := range z {
		for j, vj := range v {
			fmt.Println(i, j, v, vj)
		}
	}

	a := [][]string{{"a", "b"}, {"c", "d"}}
	fmt.Println(a)
}
