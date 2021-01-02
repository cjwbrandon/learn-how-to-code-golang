package main

import "fmt"

func main() {
	year := 1997
	for {
		year++

		if year >= 2020 {
			break
		}
		fmt.Println(year)
	}
}
