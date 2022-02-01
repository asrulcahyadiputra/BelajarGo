package main

import "fmt"

func main(){
    name := "Cahyadi"

    if name == "Asrul" {
        fmt.Println("Hello Asrul!")
    } else if(name == "Cahyadi"){
        fmt.Println("Hello Cahyadi!")
    }else{
        fmt.Println("Saya tidak mengenal anda!")
    }

    if length := len(name); length > 5{
        fmt.Println("Nama kepanjangan!")
    }else{
        fmt.Println("Nama sudah benar!")
    }
}