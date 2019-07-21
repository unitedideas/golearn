package main

import "fmt"

type xray int

var a5 xray

var y5 int

func main() {
	fmt.Println(a5)
	fmt.Printf("%T\n", a5)
	a5 = 42
	fmt.Println(a5)
	y5 = int(a5)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)

}
