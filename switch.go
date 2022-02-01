package main 

import "fmt"

func main(){
    name := "Asrul"

    switch name {
        case "Asrul":
            fmt.Println("Hello Asrul")
        case "Cahyadi": 
            fmt.Println("Hello Cahyadi")
        default:
            fmt.Println("Tidak ada Nama")
    }

    // short statement
    // switch length := len(name); length > 5 {
    //     case true:
    //         fmt.Println("Nama Kepanjangan")
    //     case false: 
    //         fmt.Println("Nama benar")
    // }

    // tanpa kondisi
    length := len(name)
    switch {
        case length > 10 :
            fmt.Println("Nama Kepanjangan")
        case length > 5 : 
            fmt.Println("Nama Lumayan Panjang!")
        default: 
            fmt.Println("Nama benar!")
    }
}