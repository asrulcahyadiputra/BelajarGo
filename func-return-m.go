package main

import "fmt"

// function return multiple value

func getFullName() (string,string){
	return "Asrul Cahyadi", "Putra"
}

func main(){
	// firstName, lastName := getFullName()
	// fmt.Println(firstName,lastName)
	firstName, _ := getFullName()
	fmt.Println(firstName)
}