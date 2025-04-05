package main

import "fmt"

// Basic function
func greet(name string) {
    fmt.Println("Hello,", name)
}

// Function with return value
func square(n int) int {
    return n * n
}

// Function with multiple return values
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}


// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 09. Functions ====")

    // Call basic function
    greet("Zephyr")

    // Call function with return value
    result := square(5)
    fmt.Println("5 squared is:", result)

    // Call function with multiple return values
    q, r := divide(17, 4)
    fmt.Println("17 divided by 4 => Quotient:", q, "Remainder:", r)

    // Anonymous function (declared and called immediately)
    func(msg string) {
        fmt.Println("Message from anonymous function:", msg)
    }("This is inline!")
}
