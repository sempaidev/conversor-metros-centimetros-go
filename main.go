package main

import "fmt"

func main() {
	fmt.Println("Bienvenido escoja una opcion de conversion: ")
	fmt.Println("1 Para convertir de metros a centimetros")
	fmt.Println("2 Para convertir de centimetro a metros")
	// Variable opcion que almacenara el valor escojido
	var opcion int
	// Capturamos el valor introducido por el usuario y se lo pasamos a variable opcion
	fmt.Scanln(&opcion)

	// Declaramos switch para manejar las opciones
	switch opcion {
	case 1:
		fmt.Println("Cuantos metros desea convertir a cm?")
		var metros float32
		fmt.Scanln(&metros)
		centimetros := metros * 100
		fmt.Printf("%.2f metros equivale a %.2f centimetros", metros, centimetros)

	case 2:
		fmt.Println("Cuantos centimetros desea convertir a metros?")
		var centimetros float32
		fmt.Scanln(&centimetros)
		metros := centimetros / 100
		fmt.Printf("%.2f centimetros equivale a %.2f metros", centimetros, metros)

	default:
		fmt.Println("Opcion no valida!")
	}
}
