package main

import "fmt"

type Mahasiswa struct {
	Name, Major string
	Age         int
	Active 		bool
}

func main(){
	var mahasiswa Mahasiswa 
	mahasiswa.Name = "Asrul Cahyadi"
	mahasiswa.Major = "Accounting Information System"
	mahasiswa.Age = 23

	fmt.Println(mahasiswa.Name)
	fmt.Println(mahasiswa.Major)
	fmt.Println(mahasiswa.Age)

	asrul := Mahasiswa{
		Name : "asrul",
		Major : "SIA",
		Age : 23,
	}

	// error pada saat struct berubah
	// cahyadi := Mahasiswa{"Cahyadi","SIA",23} 

	fmt.Println(asrul)
	// fmt.Println(cahyadi)

}