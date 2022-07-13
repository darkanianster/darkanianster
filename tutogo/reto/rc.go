package main

import "fmt"

func cociente(a, b int) int {
	return a / b
}

func residuo(a, b int) int {
	return a % b

}

func main() {

	var a, b, c, r int

	fmt.Print("ingrese n01: ")
	fmt.Scanln(&a)

	fmt.Print("ingrese n02: ")
	fmt.Scanln(&b)

	//cociente
	c = cociente(a, b)

	//residuo
	r = residuo(a, b)

	//resultados
	fmt.Println("cociente:", c)
	fmt.Println("residuo:", r)

}
