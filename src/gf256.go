package main

import (
    "fmt"
)

type Element struct {
    val int
}

type Field interface {
    add(x Element, y Element) Element
    mul(x Element, y Element) Element
    div(x Element, y Element) Element
    inv(e Element) Element
}

func main() {
    fmt.Println("Test.")
}
