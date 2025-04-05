package main

import "fmt"

// Define an interface
type Speaker interface {
    Speak()
}

// Define a struct type
type Human struct {
    Name string
}

// Human implements Speaker by defining Speak()
func (h Human) Speak() {
    fmt.Println(h.Name, "says: Hello!")
}

// Another type
type Robot struct {
    ID int
}

func (r Robot) Speak() {
    fmt.Printf("Robot #%d says: Beep boop.\n", r.ID)
}

// A function that takes an interface
func announce(s Speaker) {
    fmt.Println("Announcing:")
    s.Speak()
}

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 11. Interfaces ====")

    // Create a Human and Robot
    alice := Human{Name: "Alice"}
    r2d2 := Robot{ID: 101}

    // Use interface function
    announce(alice)
    announce(r2d2)

    // Store multiple types in a slice of interfaces
    speakers := []Speaker{alice, r2d2}
    fmt.Println("All speakers:")
    for _, s := range speakers {
        s.Speak()
    }
}
