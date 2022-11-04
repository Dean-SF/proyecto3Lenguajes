package main

import (
	"fmt"
	gen "lenguajes/GeneradorPseudo-Aleatorios/modulos"
	avl "lenguajes/arbolAvl/modulos"
	bst "lenguajes/arbolBinario/modulos"
	"math"
	"strconv"
)

var NUMEROS_PRIMOS = [42]int{11, 13, 17, 19, 23, 29, 31, 37, 41,
	43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func experimento(tamannio int, semilla int) {

	ArbolAVL := avl.AVLTree[int]{Root: nil}
	ArbolBST := bst.BinaryTree{}
	ArbolDSW := bst.BinaryTree{}

	generador := gen.GeneradorParametros{Semilla: semilla, N: tamannio, X: 2040, Limit: tamannio}

	numeros := make([]int, tamannio)

	ArbolAVLComp := 0
	ArbolBSTComp := 0
	ArbolDSWComp := 0

	ArbolAVLFind := 0
	ArbolBSTFind := 0
	ArbolDSWFind := 0

	for i := 0; i < tamannio; i++ {
		numeros[i] = generador.Generador()
		ArbolAVLComp += ArbolAVL.Insert(numeros[i])
		ArbolBSTComp += ArbolBST.Insert(numeros[i])
		ArbolDSWComp += ArbolDSW.Insert(numeros[i])
	}

	ArbolDSW.DSW_Algorithm()

	generador = gen.GeneradorParametros{Semilla: semilla, N: 3001, X: 2040, Limit: 3001}
	for i, j := 0, 0; i < 100000; i, j = i+1, j+1 {
		if j >= 3000 {
			generador = gen.GeneradorParametros{Semilla: NUMEROS_PRIMOS[generador.Generador()%42], N: 3001, X: 2040, Limit: 3001}
			j = 0
		}
		numAleatorio := generador.Generador()
		_, AVLComp := ArbolAVL.Find(numAleatorio)
		BSTFResult := ArbolBST.Find(numAleatorio)
		DSWFResult := ArbolDSW.Find(numAleatorio)

		ArbolAVLFind += AVLComp
		ArbolBSTFind += BSTFResult.Num_Comparison
		ArbolDSWFind += DSWFResult.Num_Comparison

	}

	heightAVL, _ := ArbolAVL.GetHeight()

	fmt.Printf("\nAltura AVL: %v", heightAVL)
	fmt.Printf("\nAltura ABB: %v", ArbolBST.GetHeight())
	fmt.Printf("\nAltura DSW: %v\n", ArbolDSW.GetHeight())

	fmt.Printf("\nProfundidad AVL: %f", math.Log2(float64(ArbolAVL.GetNumberNodes())))
	fmt.Printf("\nProfundidad ABB: %f", math.Log2(float64(ArbolBST.GetNumberNodes())))
	fmt.Printf("\nProfundidad DSW: %f\n", math.Log2(float64(ArbolDSW.GetNumberNodes())))

	fmt.Printf("\nDensidad AVL: %v", (ArbolAVL.GetNumberNodes() / heightAVL))
	fmt.Printf("\nDensidad ABB: %v", (ArbolBST.GetNumberNodes() / ArbolBST.GetHeight()))
	fmt.Printf("\nDensidad DSW: %v\n", (ArbolDSW.GetNumberNodes() / ArbolDSW.GetHeight()))

	fmt.Printf("\nComparaciones AVL(insercion): %v", ArbolAVLComp)
	fmt.Printf("\nComparaciones BST(insercion): %v", ArbolBSTComp)
	fmt.Printf("\nComparaciones DSW(insercion): %v\n", ArbolDSWComp)

	fmt.Printf("\nComparaciones AVL(find): %v", ArbolAVLFind)
	fmt.Printf("\nComparaciones BST(find): %v", ArbolBSTFind)
	fmt.Printf("\nComparaciones DSW(find): %v\n", ArbolDSWFind)
}

func main() {
	var tamannio = [5]int{500, 1000, 2000, 3500, 5000}
	for i, value := range tamannio {
		fmt.Println("\n\n >>>> Experimento #" + strconv.Itoa(i+1))
		experimento(value, NUMEROS_PRIMOS[value%42])
	}
}
