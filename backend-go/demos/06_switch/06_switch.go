package main

import "fmt"

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 06. switch Statement ====")

    day := 3

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    case 5:
        fmt.Println("Friday")
    case 6:
        fmt.Println("Saturday")
    case 7:
        fmt.Println("Sunday")
    default:
        fmt.Println("Invalid day in this universe.")
    }

    // Switch with string
    language := "go"

    switch language {
    case "go":
        fmt.Println("Go is awesome!")
    case "python":
        fmt.Println("Python is cool too.")
    default:
        fmt.Println("Unknown language.")
    }

    // Switch without a condition
    score := 85

    switch {
    case score >= 90:
        fmt.Println("Grade: A+")
    case score >= 80:
        fmt.Println("Grade: A-")
    case score >= 70:
        fmt.Println("Grade: B")
    default:
        fmt.Println("Grade: C or below")
    }
}
