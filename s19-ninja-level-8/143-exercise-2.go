package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	s := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27}]`
	b := []byte(s)

	var people []user
	err := json.Unmarshal(b, &people)
	if err != nil {
		fmt.Println(err)
	}
}
