package main 

import "fmt" 

func main(){
    // counter := 1

    // for counter <= 10  {
    //     fmt.Println("Perulangan Ke ", counter)
    //     counter++
    // }

    // for counter := 1; counter <= 10; counter++  {
    //     fmt.Println("Perulangan Ke ", counter)
    // }

    slice := []string{"Asrul","Cahyadi","Putra"}

    // for i := 0; i < len(slice); i++{
    //     fmt.Println(slice[i])
    // }

    for i, val := range slice {
        fmt.Println("index",i,"=",val)
    }

    for _, val := range slice {
        fmt.Println(val)
    }

    person := make(map[string]string)

    person["title"] = "Programmer"
    person["name"] = "Asrul"

    for key, value := range person{
        fmt.Println(key,"=",value)
    }
}