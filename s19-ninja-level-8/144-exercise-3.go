package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{"James", 32}
	u2 := user{"Moneypenny", 27}
	users := []user{u1, u2}

	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}
