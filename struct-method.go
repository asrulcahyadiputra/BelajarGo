package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) welcomeMessage(name string){
	fmt.Println("Welcome", name, "My Name is",customer.Name)
}

func main() {
	var customer Customer 
	customer.Name = "Cahyadi"
	customer.welcomeMessage("Asrul")
}