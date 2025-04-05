package main

import "fmt"

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 02. Arithmetic Operators ====")

    // Addition (+)
    a := 10
    b := 5
    sum := a + b
    fmt.Println("Addition (a + b):", sum)

    // Subtraction (−)
    diff := a - b
    fmt.Println("Subtraction (a - b):", diff)

    // Multiplication (×)
    product := a * b
    fmt.Println("Multiplication (a * b):", product)

    // Division (÷)
    quotient := a / b
    fmt.Println("Integer Division (a / b):", quotient)

    // Modulo (%)
    remainder := a % b
    fmt.Println("Modulo (a % b):", remainder)

    // Float division (if needed)
    x := 10.0
    y := 4.0
    result := x / y
    fmt.Println("Float Division (x / y):", result)
}
