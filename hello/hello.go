package main

import (
	"fmt"
)

func main() {
	var num1 = 5.6
	var num2 = 9.5
	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "There"
	fmt.Println(multiple(w1, w2))



}

func multiple(a, b string) (string, string) {
	return a, b
}

func add(x float64, y float64) float64 {
	return x + y
}


