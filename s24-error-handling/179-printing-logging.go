package main

import (
	"fmt"
	"log"
	"os"
)

// fmt.Println()
// log.Println()
// log.Fatalln() -> os.Exit()
// log.Panicln() -> deferred functions run, can use "recover"
// panic()

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	log.SetOutput(f)

	defer foo()
	_, err = os.Open("no-file.txt")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		log.Fatalln("Fatal", err)
		log.Panicln("Panic", err)
	}
}

func foo() {
	fmt.Println("deferred functions will run with panic and not with fatal")
}
