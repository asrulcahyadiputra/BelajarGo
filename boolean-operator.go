package main

import "fmt"

func main(){
    // Operator boolean
    var ujian = 80
    var absensi = 60

    var lulusUjian = ujian >= 80
    var lulusAbsensi = absensi >= 75

    var lulus bool = lulusUjian && lulusAbsensi 

    fmt.Println(lulusUjian)
    fmt.Println(lulusAbsensi)
    fmt.Println(lulus)
}