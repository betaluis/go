package main

import "fmt"

func main() {
    // Integers

    var ageOne int = 20
    var ageTwo = 30
    ageThree := 40

    fmt.Println(ageOne, ageTwo, ageThree)

    // Bits and memory
    var numOne int8 = 25 // 8 bits = 1 byte - Range: -128 to 127
    var numTwo int8 = -128

    // You can use int8, int16, int32, int64

    /////////////////////////////////////////////
    var numThree uint8 = 255
    
    // 8 bits = 1 byte - Range: 0 to 255 
    // unsigned int means it can't be negative
    // Range: 0 to 255
    
    // for uint you can use uint8, uint16, uint32, uint64

    /////////////////////////////////////////////

    // Float
    var scoreOne float32 = 25.98
    var scoreTwo float64 = 356789.123456789
    scoreThree := 1.5 // infers float64

}

