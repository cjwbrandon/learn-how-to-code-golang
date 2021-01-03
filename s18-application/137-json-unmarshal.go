package main

import (
	"encoding/json"
	"fmt"
)

// Field Type == Field Type `json:Field`
// Field Type `json:"myName"` -> Different names
// Field Type `json:"myName,omitempty"` -> Ignore empty fields
// Field Type `json:",omitempty"` == Field Type `json:Field,omitempty`
// Field Type `json:"-"` -> Ignore

type person struct {
	First string `json:"First`
	Last  string `json:"last"`
	Age   int
}

func main() {
	s := `[{"First":"James","last":"Bond","Age":32},{"First":"Miss","last":"Moneypenny","Age":27}]`
	bs := []byte(s)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
