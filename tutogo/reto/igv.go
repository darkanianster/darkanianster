package main

import "fmt"

//funcion que  calcula el igv

func calcularIgv(vv, igv float64) float64 {
	return igv * vv
}

//func calcula el precio de venta

func precioVenta(vv, igv float64) float64 {
	return vv + igv
}

//funcion principal

func main() {

	//variables
	var vv, igv, pv float64

	//entrada de datos
	fmt.Print("Ingrese valor de venta: ")
	fmt.Scanln(&vv)

	//igv
	igv = calcularIgv(vv, 0.18)
	//precio de venta
	pv = precioVenta(vv, igv)

	//salida de datos

	fmt.Println("====================")
	fmt.Println("Valor de Venta:", vv)
	fmt.Println("IGV:", igv)
	fmt.Println("Precio de venta:", pv)
	fmt.Println("======================")

}
