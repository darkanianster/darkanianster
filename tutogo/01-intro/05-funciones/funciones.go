package main

import "fmt"

func saludar(nombre string) {
	//fmt.Println("Hola mundo")
	fmt.Println("Hola,", nombre)
}

func sumar(a, b int) int {
	return a + b

}

func main() {
	saludar("jose")
	r := sumar(10, 20)
	fmt.Println("La suma es:", r)
}
