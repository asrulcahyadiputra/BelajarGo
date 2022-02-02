package main

import "fmt"

func main() {
	name := "Asrul"
	increment := func() {
		fmt.Println("Increment")
		name = "Cahyadi"
	}

	increment()
	increment()
	fmt.Println(name)
}