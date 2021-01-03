package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string
	Age   int
}

type byFirst []user

func (a byFirst) Len() int           { return len(a) }
func (a byFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFirst) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	u1 := user{"James", 32}
	u2 := user{"Moneypenny", 27}
	users := []user{u1, u2}

	sort.Sort(byFirst(users))
	fmt.Println(users)
}
