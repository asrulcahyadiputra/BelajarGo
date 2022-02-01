package main 

import "fmt"

func main(){

    // tipe data array
    var names[3]string

    names[0] = "Asrul"
    names[1] = "Cahyadi"
    names[2] = "Putra"

    fmt.Println(names[0])
    fmt.Println(names[1])
    fmt.Println(names[2])

    var values = [3]int{
        1,
        2,
        3,
    }
    fmt.Println(values)
    //panjang data bukan jumlah
    fmt.Println(len(values))
    
    var arrayKosong [10]string
    fmt.Println(len(arrayKosong))
}