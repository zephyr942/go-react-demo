package main

import "fmt"

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 03. Arrays ====")

    // Declare an array with fixed size
    var nums [3]int
    nums[0] = 10
    nums[1] = 20
    nums[2] = 30

    fmt.Println("Array nums:", nums)

    // Shorter syntax
    letters := [3]string{"A", "B", "C"}
    fmt.Println("Array letters:", letters)

    // Using the '...' operator to infer size
    primes := [...]int{2, 3, 5, 7}
    fmt.Println("Array primes:", primes)

    // Access by index
    fmt.Println("Second prime number:", primes[1])

    // Array length
    fmt.Println("Length of primes array:", len(primes))
}
