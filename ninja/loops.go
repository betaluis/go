package main

import "fmt"

func main() {
    // x := 0
    //
    // for x < 5 {
    //     fmt.Println(x)
    //     x++
    // }

    // for i := 0; i < 5; i++ {
    //     fmt.Println(i)
    // }

    names := []string{"luis", "lauren", "lillie"}
    //
    // for i := 0; i < len(names); i++ {
    //     fmt.Println(names[i])
    // }

    for index, value := range names {
        fmt.Printf("index: %v, value: %v \n", index, value)
    }

    // If we don't want the index
    // for _, value := range names {
    //     //...
    // }
}

