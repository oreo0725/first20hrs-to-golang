package main

import (
	"fmt"
	// "strconv"
	// "zentest.io/day2/test"
)

func intTest() {
	var i int8
	i = 10
	fmt.Println("i = ", i)
}

func getResult() (result string, errCode int) {
	defer fmt.Println("fallback 1")
	// defer fmt.Println("fallback 2")
	// defer fmt.Println("fallback 3")
	// defer fmt.Println("fallback 4")

	fmt.Println("fallback 0")
	return "Hello", 11
}

func main() {
	// var a, b = 4, 5
	// var c int

	// c, b = a, c

	// fmt.Println(a, b, c)

	// var s1, s2 string = "", "1"
	// fmt.Println(s1, s2)

	// //anoymous function
	// f := func(input int) (string, string) {
	// 	//string to int: type conversion
	// 	return strconv.Itoa(input + 1), strconv.Itoa(input - 1)
	// }
	// fmt.Println(f(99))

	// intTest()

	// // fmt.Println("Pi", test.PI, ", Monday:", test.Monday)

	// test.PrintConstants()

	result, errCode := getResult()
	fmt.Println("result => ", result, errCode)

}
