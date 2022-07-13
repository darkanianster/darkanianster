package main

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func main() {

	var numero1, numero2 int

	fmt.Print("numero1: ")
	fmt.Scanln(&numero1)

	fmt.Print("numero2: ")
	fmt.Scanln(&numero2)

	suma := numero1 + numero2

	fmt.Printf("%d + %d = %d", numero1, numero2, suma)

}
