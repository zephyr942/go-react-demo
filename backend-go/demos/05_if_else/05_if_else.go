package main

import "fmt"

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 05. if / else Conditions ====")

    number := -7

    if number > 0 {
        fmt.Println("The number is positive.")
    } else if number < 0 {
        fmt.Println("The number is negative.")
    } else {
        fmt.Println("The number is zero.")
    }

    // Example with boolean condition
    isRainy := true
    if isRainy {
        fmt.Println("Grab an umbrella.")
    } else {
        fmt.Println("No umbrella needed.")
    }

    // Inline condition (if statement with short variable declaration)
    if age := 18; age >= 18 {
        fmt.Println("Welcome to adulthood – bills, taxes, and coffee addiction!")
    } else {
        fmt.Println("Enjoy your youth while it lasts!")
    }
}
