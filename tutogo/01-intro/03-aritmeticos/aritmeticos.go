package main

import "fmt"

func main() {

	var a int
	a = 20
	var b int
	b = 10

	//suma
	var result int
	result = a + b
	fmt.Println("suma:", result)

	//Resta
	result = a - b
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = a * b
	fmt.Println("Multiplicacion:", result)

	//Division
	result = a / b
	fmt.Println("Division:", result)

	var div float64 = 3.0 / 2.0
	fmt.Println("Division:", div)

	//Modulo
	result = 3 % 2
	fmt.Println("modulo:", result)

}
