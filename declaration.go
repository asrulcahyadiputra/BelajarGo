package main

import "fmt"

func main(){
    /** 
    * declaration
    * alias dari tipe data yang sudah dibuat
    */
    type noKTP string

    type Married bool
    
    var noKtpAsrul noKTP = "123456789"
    var marriedStatus Married = true
    fmt.Println(noKtpAsrul)
    fmt.Println(marriedStatus)

}