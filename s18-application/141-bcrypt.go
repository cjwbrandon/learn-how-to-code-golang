package main

import "fmt"

func main() {
	s := "password123"

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println(`YOU CAN'T LOGIN`)
		return
	}

	fmt.Println(`You're logged in!`)
}
