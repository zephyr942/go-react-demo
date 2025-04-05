package main

import "fmt"

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
    fmt.Println("==== 04. Slices ====")

    // Create a slice from an array
    arr := [5]int{10, 20, 30, 40, 50}
    slice1 := arr[1:4] // includes index 1 to 3 (not 4)
    fmt.Println("Slice from array:", slice1)

    // Declare a slice directly
    nums := []int{1, 2, 3}
    fmt.Println("Original slice:", nums)

    // Append elements (creates a new slice)
    nums = append(nums, 4)
    nums = append(nums, 5, 6)
    fmt.Println("After append:", nums)

    // Slice length and capacity
    fmt.Println("Length:", len(nums))
    fmt.Println("Capacity:", cap(nums))

    // Create an empty slice and fill it
    var fruits []string
    fruits = append(fruits, "apple", "banana", "cherry")
    fmt.Println("Fruit slice:", fruits)

    // Modify slice element
    fruits[1] = "blueberry"
    fmt.Println("Modified fruits:", fruits)
}
