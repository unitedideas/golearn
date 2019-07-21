package chap1ex

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

// Identifiers are assigned the zero value

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
