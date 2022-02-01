package main

import "fmt"

func main(){
    /**
    * konstanta
    * konstanta tidak dapat diubah, tidak error
    * jika tidak digunakan. Berbeda dengan variable
    */
    const firstName string = "Cahyadi"
    const lastName  = "Putra"
    const number  = 1000

    // multiple constant
    const (
        city = "Bulukumba"
        country = "Indonesia"
    )

    fmt.Println(firstName)
    fmt.Println(lastName)
    fmt.Println(number)
    fmt.Println(city)
    fmt.Println(country)
}