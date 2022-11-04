package main

import (
	"fmt"
	"strconv"
	bst "lenguajes/arbolBinario/modulos"
	gen "lenguajes/GeneradorPseudo-Aleatorios/modulos"
	avl "lenguajes/arbolAvl/modulos"
)

var NUMEROS_PRIMOS = [42]int{11, 13, 17, 19, 23, 29, 31, 37, 41, 
	43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func experimento(tamannio int, semilla int) {

	ArbolAVL := avl.AVLTree[int]{Root: nil}
	ArbolBST := bst.BinaryTree{}
	ArbolDSW := bst.BinaryTree{}

	generador := gen.Generador{Semilla: semilla, N: tamannio, X: 2040, Limit: tamannio}
	
	numeros := make([]int,tamannio)
	
	ArbolAVLComp := 0
	ArbolBSTComp := 0
	ArbolDSWComp := 0

	ArbolAVLFind := 0
	ArbolBSTFind := 0
	ArbolDSWFind := 0

	for i := 0; i < tamannio; i++ {
		numeros[i] = generador.Next()
		ArbolAVLComp += ArbolAVL.Insert(numeros[i]) 
		ArbolBSTComp += ArbolBST.Insert(numeros[i])
		ArbolDSWComp += ArbolDSW.Insert(numeros[i])
	}
	
	generador = gen.Generador{Semilla: semilla, N: 3001, X: 2040, Limit: 3001}
	for i,j := 0,0; i < 100000; i,j = i+1,j+1 {
		if j >= 3000 {
			generador = gen.Generador{Semilla: NUMEROS_PRIMOS[generador.Next()%42], N: 3001, X: 2040, Limit: 3001}
			j = 0
		}
		numAleatorio := generador.Next()
		_,AVLComp := ArbolAVL.Find(numAleatorio)
		BSTFResult := ArbolBST.Find(numAleatorio)
		DSWFResult := ArbolDSW.Find(numAleatorio)

		ArbolAVLFind += AVLComp
		ArbolBSTFind += BSTFResult.Num_Comparison
		ArbolDSWFind += DSWFResult.Num_Comparison
		
	}	
	fmt.Println(ArbolAVLComp)
	fmt.Println(ArbolBSTComp)
	fmt.Println(ArbolDSWComp)
	
	fmt.Println()

	fmt.Println(ArbolAVLFind)
	fmt.Println(ArbolBSTFind)
	fmt.Println(ArbolDSWFind)

	fmt.Println()

}

func main() {
	var tamannio = [5]int{500,1000,2000,3500,5000}
	for i, value := range tamannio {
		fmt.Println("Experimento #" + strconv.Itoa(i+1))
		experimento(value,NUMEROS_PRIMOS[value%42])
	}
}