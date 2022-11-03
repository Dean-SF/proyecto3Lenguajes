package main

import (
	"fmt"
	gen "lenguajes/GeneradorPseudo-Aleatorios/modulos"
)

func main() {
	// Generador struct = {semilla, n, x, limite}
	// n se utiliza para ir restando de 1 en 1 a la cantidad de numeros a generar y parar el trabajo cuando se llega a 0
	// x será una constante

	/*
		limite es una constante con la cantidad de numeros a generar, se utiliza para identificar cuándo se está por generar
		el primer número, para así inicializar x con la semilla.
		Se debe ingresar el mismo número para n y límite
	*/

	gen := gen.Generador{11, 500, 2040, 500}

	// Cada vez que se hace Next se consigue el siguiente numero generado
	for {
		result := gen.Next()
		if result == nil {
			return
		}
		fmt.Printf("%v ", *result)
	}
}
