package main

import "fmt"

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 07. Maps ====")

    // Declare and initialize a map
    scores := map[string]int{
        "Alice": 90,
        "Bob":   85,
        "Carol": 92,
    }
    fmt.Println("Initial scores:", scores)

    // Access value by key
    fmt.Println("Bob's score:", scores["Bob"])

    // Add or update a value
    scores["David"] = 88
    scores["Alice"] = 95
    fmt.Println("Updated scores:", scores)

    // Delete a key
    delete(scores, "Bob")
    fmt.Println("After deleting Bob:", scores)

    // Check if a key exists
    val, exists := scores["Carol"]
    if exists {
        fmt.Println("Carol's score is:", val)
    } else {
        fmt.Println("Carol not found")
    }

    // Iterate over the map
    fmt.Println("All scores:")
    for name, score := range scores {
        fmt.Println(name, "->", score)
    }
}
