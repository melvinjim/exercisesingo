package main

import (
	"fmt"
)

func main() {
	var respuesta string
	var firstNumber int
	var secondNumber int

	fmt.Print("desea realizar una suma, resta, multiplicacion o division?-->")
	fmt.Scanln(&respuesta)

	if respuesta == "suma" {
		fmt.Print("¿Cuál es el primer numero desea sumar? ")
		fmt.Scanln(&firstNumber)
		fmt.Print("¿Cuál es el segundo numero que desea sumar? ")
		fmt.Scanln(&secondNumber)

		var suma = firstNumber + secondNumber
		fmt.Printf("el resultado es: %d\n", suma)
	}
}
