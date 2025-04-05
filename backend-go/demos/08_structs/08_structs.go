package main

import "fmt"

// Define a struct type
type Person struct {
    Name string
    Age  int
    IsStudent bool
}

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 08. Structs ====")

    // Create a struct instance
    var p1 Person
    p1.Name = "Alice"
    p1.Age = 21
    p1.IsStudent = true

    // Print full struct
    fmt.Println("Person 1:", p1)

    // Access fields
    fmt.Println("Name:", p1.Name)
    fmt.Println("Age:", p1.Age)
    fmt.Println("Is student:", p1.IsStudent)

    // Another way: struct literal
    p2 := Person{
        Name: "Bob",
        Age:  30,
        IsStudent: false,
    }
    fmt.Println("Person 2:", p2)

    // Modify fields
    p2.Age = 31
    fmt.Println("Updated age for Bob:", p2.Age)
}
