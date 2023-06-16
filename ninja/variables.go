package main

import "fmt"

func main() {
    var nameOne string = "Mario"
    var nameTwo = "Luigi"
    var nameThree string

    fmt.Println(nameOne, nameTwo, nameThree)

    nameOne = "Peach"
    nameThree = "Bowser"

    fmt.Println(nameOne, nameTwo, nameThree)

    nameFour := "Yoshi"
    fmt.Println(nameFour)
}

