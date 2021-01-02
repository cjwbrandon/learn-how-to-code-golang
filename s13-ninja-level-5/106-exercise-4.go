package main

import "fmt"

func main() {
	p1 := struct {
		name     map[string]string
		children []string
	}{
		name: map[string]string{
			"first": "Bobby",
			"last":  "Kim",
		},
		children: []string{
			"Peter", "Katie",
		},
	}
	fmt.Println(p1)
}
