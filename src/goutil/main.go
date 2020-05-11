package main

import "C"
import (
    "fmt"
)

// gotchas:
// - package name must be main
// - must import pkg C
// - export a func by prepending `//export ...` (no space aka //e...)
// - for .h signature to be C native type, func should use C.* types
// - see README for build instructions

//export Add
func Add(a int, b int) int {
    return a + b
}

//export Minus
func Minus(a C.int, b C.int) C.int {
    return a - b
}

//export Greet
func Greet() {
    fmt.Println("hello from go");
}

func main() {
    // do not remove!
    // cgo needs main() to make C shared lib
}
