package main

import (
	"lenguajes/arbolAvl/modulos"
	"fmt"
	"strconv"
)

func main(){
	var arbol = avl.AVLTree[int]{Root: nil}
	fmt.Println("Insercion llave -> 6 Comparaciones -> " + strconv.Itoa(arbol.Insert(6)))
	fmt.Println("Insercion llave -> 2 Comparaciones -> " + strconv.Itoa(arbol.Insert(2)))
	fmt.Println("Insercion llave -> 3 Comparaciones -> " + strconv.Itoa(arbol.Insert(3)))
	fmt.Println("Insercion llave -> 5 Comparaciones -> " + strconv.Itoa(arbol.Insert(5)))
	fmt.Println("Insercion llave -> 8 Comparaciones -> " + strconv.Itoa(arbol.Insert(8)))
	fmt.Println("Insercion llave -> 9 Comparaciones -> " + strconv.Itoa(arbol.Insert(9)))
	fmt.Println("Insercion llave -> 6 Comparaciones -> " + strconv.Itoa(arbol.Insert(6)))
	fmt.Println(arbol.Find(2))
	fmt.Println(arbol.Find(1))
	fmt.Println(arbol.Find(29))
	fmt.Println(arbol.Find(3))
	arbol.PrintInorder()
	
}