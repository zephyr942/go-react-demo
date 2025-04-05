package main

import "fmt"

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 10. Pointers ====")

    // Declare a variable
    x := 10
    fmt.Println("Initial value of x:", x)

    // Declare a pointer to x
    var p *int = &x
    fmt.Println("Pointer p holds address of x:", p)
    fmt.Println("Value at address p points to (*p):", *p)

    // Change value via pointer
    *p = 20
    fmt.Println("New value of x after *p = 20:", x)

    // Pass by reference example
    y := 5
    fmt.Println("Before double(): y =", y)
    double(&y)
    fmt.Println("After double(): y =", y)
}

// Function that modifies value using pointer
func double(n *int) {
    *n = *n * 2
}
