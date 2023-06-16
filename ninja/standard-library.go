package main

import (
	"fmt"
	"sort"
)

func main() {
    // greeting := "Hello there friends!"

    // fmt.Println(strings.Contains(greeting, "there")) // true
    // fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi")) // Hi there friends!
    // fmt.Println(strings.ToUpper(greeting)) // HELLO THERE FRIENDS!
    // fmt.Println(strings.Index(greeting, "ll")) // 2
    // fmt.Println(strings.Split(greeting, " ")) // [Hello there friends!]

    // Sort a slice of integers
    ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
    sort.Ints(ages)
    fmt.Println(ages)

    index := sort.SearchInts(ages, 30)
    fmt.Println(index)

    names := []string{"luis", "lauren", "lillie"}
    sort.Strings(names)
    fmt.Println(names)

    fmt.Println(sort.SearchStrings(names, "lillie"))

}
