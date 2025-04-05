package main

import "fmt"

// This main.go is for demo purposes only.
// In real-world Go projects, we use a single main.go and import modules instead.
func main() {

    fmt.Println("==== 01: Variables ====")
    
    
    // Basic Variable Declaration
    var a int = 10
    b := 20                // shorthand declaration，strongly recommended
    var c string = "Hello World"
    var d float64 = 3.14159
    var e bool = true

    // Multiple variables in one line
    var x, y int = 1, 2
    m,n,o := "Golang","is", "great!" // Go compiler does not allow declared variables to be unused.


    // Print types
    fmt.Println("a =", a)
    fmt.Println("b =", b)
    fmt.Println("c =", c)
    fmt.Println("d =", d)
    fmt.Println("e =", e)
    fmt.Println("x =", x, "y =", y)
    fmt.Println( m, n, o) 
    // Print typess
    fmt.Printf("Type of a: %T\n", a)
    fmt.Printf("Type of c: %T\n", c)
    fmt.Printf("Type of d: %T\n", d)
    fmt.Printf("Type of e: %T\n", e)
}
