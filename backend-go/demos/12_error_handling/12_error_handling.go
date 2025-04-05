package main

import (
	"errors"
	"fmt"
)

// divide returns the result of a / b, or an error if b is zero
func safeDivide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {
	fmt.Println("==== 12. Error Handling ====")

	// Case 1: valid division
	result1, err1 := safeDivide(10, 2)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("10 divided by 2 =", result1)
	}

	// Case 2: division by zero
	result2, err2 := safeDivide(8, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("8 divided by 0 =", result2)
	}

	// Case 3: Custom error with fmt.Errorf
	name := ""
	if name == "" {
		customErr := fmt.Errorf("name cannot be empty")
		fmt.Println("Custom error:", customErr)
	}
}
