package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo ¡"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Jose"
	edad := 36

	fmt.Printf("Hola %s, su edad es %d \n", nombre, edad)
	fmt.Printf("Hola %v, su edad es %v \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola %s, su edad es %d ", nombre, edad)
	fmt.Println(mensaje)

	fmt.Printf("nombre: %T \n", nombre)
	fmt.Printf("edad: %T \n", edad)

	fmt.Print("ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("otro Nombre: ", nombre)
}
