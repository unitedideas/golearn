package main

import (
	"fmt" 
)

// pointers
func main() {
	for z := 0; z < 100; z++{
		prtMem(&z)
		fmt.Println(z)
	}
}

func prtMem(x *int){
	*x++
}

// // Struct
// type human struct{
// 	name string
// 	age int
// }

// func main() {
// 	person := human{name: "steve", age: 32}
// 	fmt.Println(person)
// 	fmt.Println(person.age)
// }


// // functions
// // go doesn't have exceptions
// func main() {
// 	result, err := sqrt(-16)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("Sqrt undefined on negative number")
// 	}
// 	return math.Sqrt(x), nil
// }

// func main() {
// 	result := sum(5,11)
// 	fmt.Println(result)

// 	// or
// 	fmt.Println(sum(5, 11))
// }

// func sum(x int, y int) int {
// 	return x + y
// }

// // loops
// func main() {
// 	m := make(map[int]string)
// 	m[10] = "Ten"
// 	m[20] = "Twenty"
// 	for key, value := range m{
// 		fmt.Println("key:", key, "value:", value)
// 	}
// // array
// arr := []string{"a", "b", "c"}
// for index, value := range arr{
// 	fmt.Println("index:", index, "value:", value)
// }
// // number
// for i := 0; i < 5; i++{
// 	fmt.Println(i)
// }
// }
// // map (dict)
// func main() {
// 	verticies := make(map[string]int)
// 	verticies["tri"] = 3
// 	verticies["sqr"] = 4
// 	verticies["oct"] = 8
// 	fmt.Println(verticies["tri"])
// 	delete(verticies, "sqr")
// 	fmt.Println(verticies)
// }

// // Array
// func main() {
// 	// // Array
// 	// var a [5]int
// 	// a:= [5]int{1,2,3,4,5}

// 	//Slice
// 	a := []int{1, 2, 3, 4, 5}
// 	a[2] = 7
// 	a = append(a, 99)
// 	fmt.Println(a)
// }

// // if
// func main() {
// 	x := 6
// 	if x > 6 {
// 		fmt.Printf(strconv.Itoa(x) + " is less than 5")
// 	} else if x < 6 {
// 		fmt.Printf(strconv.Itoa(x) + " is greater than 5")
// 	} else {
// 		fmt.Printf(strconv.Itoa(x)+" = "+"%v", x)
// 	}
// }

// // if
// func main() {
// 	x:= 6
// 	if x > 5 {
// 		fmt.Printf(strconv.Itoa(x) + " is less than 5")
// 	}
// }
// Const
// func main() {
// 	const myConst int = 42
// 	fmt.Printf("%v, %T\n", myConst, myConst)
// }

// // Runes
// func main() {
// 	r := 'a'
// 	fmt.Printf("%v, %T\n", r, r)
// }

// // Boolean
// func main() {
// 	r := false
// 	r = true
// 	fmt.Printf("%v, %T\n", r, r)
// }

// // Text Types
// func main() {
// 	s := "this is a string"
// 	b := []byte(s)
// 	fmt.Printf("%v, %T\n", b, b)
// }

// // Numaric Types
// func main() {
// 	var n uint = 42
// 	fmt.Printf("%v, %T\n", n, n)
// 	var m bool
// 	fmt.Printf("%v, %T\n", m, m)
// }

// var i int = 42

// var (
// // actorName string = 	'Liz Sladen'
// // companion string = ' Sarah Smith'
// // docNum int = 3
// // season int = 2
// )

// func main() {
// 	var i int
// 	i = 42
// 	fmt.Printf("%v, %T\n", i, i)

// 	// var j float32 = 27
// 	// k :=99
// 	var j string
// 	j = strconv.Itoa(i)
// 	fmt.Printf("%v, %T\n", j, j)
// }
