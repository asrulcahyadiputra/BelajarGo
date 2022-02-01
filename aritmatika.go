package main 

import "fmt"

func main(){
    var a = 10
    var b = 10

    var c = a + b 
    var d = a * b 
    var e = a - b 
    var f = a / b 

    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)

    // augmented assigments
    var i = 10 
    i += 10
    fmt.Println(i)
    // unery operator
    i++ 

    fmt.Println(i)

    var negative = -100
    fmt.Println(negative)
}