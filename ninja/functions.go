package main

import (
	"fmt"
	"math"
)

func sayGreeting (name string) {
    fmt.Printf("Good morning, %v! \n", name);
}
func sayBye (name string) {
    fmt.Printf("Bye, %v! \n", name);
}

func cycleNames (names []string, f func(string)) {
    for i := 0; i < len(names); i++ {
        f(names[i])
    }
}

func circleArea (r float64) float64 {
    return math.Pi * r * r
}

func main () {
    // names := []string{"luis", "lauren", "lillie"}
    // cycleNames(names, sayGreeting)

    a1 := circleArea(10.5)
    a2 := circleArea(15)
    a3 := circleArea(2)

    fmt.Printf("Circle 1 is %0.1f,\ncircle 2 is %0.1f,\nand circle 3 is %0.1f.", a1, a2, a3)
}
