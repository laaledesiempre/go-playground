package main

import "fmt"

func main() {
    // Infered type
    var a = "initial"
    fmt.Println(a)

    // Explicit type
    var b,c int = 1,2
    fmt.Println(b, c)

    // implicit boolean
    var d = true
    fmt.Println(d)
    // Zero value default
    var e int
    fmt.Println(e)

    f := "7.0"
    fmt.Println(f)
}
