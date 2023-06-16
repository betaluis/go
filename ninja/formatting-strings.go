package main

import "fmt"

func main () {

    age := 25
    name := "John Doe"

    // Print
    // fmt.Print("Hello,")
    // fmt.Print("World! \n")
    // fmt.Print("New line \n")

    // Println
    // Same thing but prints a new line char at the end
    // fmt.Println("Hello, World.")
    // fmt.Println("Goodbye, World.")

    // Printing variables
    // fmt.Println("My name is", name, "and I am", age, "years old.")

    // Formatting strings
    // This is a way to format strings with variables in them
    // fmt.Printf("My name is %v and I am %v years old.\n", age, name)
    // fmt.Printf("My name is %q and I am %q years old.\n", age, name) // %q prints the value in double quotes
    // fmt.Printf("age is of type %T\n", age) // %T prints the type of the variable
    // fmt.Printf("You scored %0.1f points!\n", 225.55) // %0.1f formats the float to 1 decimal place


    // Sprintf
    // This is a way to format strings and store them in a variable
    var str = fmt.Sprintf("My name is %v and I am %v years old.\n", age, name)
    fmt.Println("The str variable is:", str)
}

