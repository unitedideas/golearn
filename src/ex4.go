package main

import "fmt"

type shane int

var a shane

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	a = 42
	fmt.Println(a)
}
