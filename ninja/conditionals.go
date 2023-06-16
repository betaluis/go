package main

import "fmt"

func main () {
    age := 28

    if age < 30 {
        fmt.Println("Age is less than 30")
    } else if age < 40 {
        fmt.Println("Age is less than 40")
    } else {
        fmt.Println("Age is not less than 45")
    }

    names := []string{"lauren", "luis", "jessica", "jon"}

    for index, value := range names {
        if index == 1 {
            fmt.Println("continuing at pos ", index)
            continue
        }

        if index > 2 {
            fmt.Println("Breaking at pos", index)
            break
        }

        fmt.Printf("the value at position %v is %v \n", index, value)
    }
}

