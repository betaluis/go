package main

import "fmt"

func main() {
    var ages [3]int = [3]int{20, 25, 30} // Length is fixed

    var names = [3]string{"John", "Doe", "Smith"}
    names[1] = "Doe Jr."

    grades := [3]int{1, 2, 3}

    fmt.Println(ages, len(ages))
    fmt.Println(names, len(names))
    fmt.Println(grades, len(grades))

    // Slices
    var scores = []int{100, 200, 300, 400, 500}
    scores[1] = 250

    // Append to slice
    scores = append(scores, 600)

    fmt.Println(scores, len(scores))
    
    // Slice ranges
    rangeOne := scores[2:5]  // up to but not including
    rangeTwo := scores[3:]   // up to the end
    rangeThree := scores[:3] // up to but not including

    fmt.Println(rangeOne)    // 300, 400
    fmt.Println(rangeTwo)    // 400, 500, 600
    fmt.Println(rangeThree)  // 100, 250, 300

    rangeOne = append(rangeOne, 350)
    fmt.Println(rangeOne)    

    evens := []int{2, 4, 6, 8, 10}
    rangeFour := evens[1:3]
    fmt.Println(rangeFour) // 4, 6
}
