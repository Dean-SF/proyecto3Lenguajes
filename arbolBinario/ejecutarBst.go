package main

import (
	"fmt"
	bst "lenguajes/arbolBinario/modulos"
)

func main() {
	// Create instance of BinaryTree
	BT_Root := &bst.BinaryTree{}

	// Insert elements
	// fmt.Printf("\nComparisons made until %d was inserted: %d", 8, BT_Root.Insert(8));
	// fmt.Printf("\nComparisons made until %d was inserted: %d", 3, BT_Root.Insert(3));
	// fmt.Printf("\nComparisons made until %d was inserted: %d", 6, BT_Root.Insert(6));
	// fmt.Printf("\nComparisons made until %d was inserted: %d", 4, BT_Root.Insert(4));
	// fmt.Printf("\nComparisons made until %d was inserted: %d", 1, BT_Root.Insert(1));
	// fmt.Printf("\nComparisons made until %d was inserted: %d", 10, BT_Root.Insert(10));
	// fmt.Printf("\nComparisons made until %d was inserted: %d", 14, BT_Root.Insert(14));
	// fmt.Printf("\nComparisons made until %d was inserted: %d", 13, BT_Root.Insert(13));
	// fmt.Printf("\nComparisons made until %d was inserted: %d", 7, BT_Root.Insert(7));

	BT_Root.Insert(24)
	BT_Root.Insert(20)
	BT_Root.Insert(6)
	BT_Root.Insert(5)
	BT_Root.Insert(7)
	BT_Root.Insert(8)
	BT_Root.Insert(9)
	BT_Root.Insert(29)
	BT_Root.Insert(27)
	BT_Root.Insert(33)

	// Print inorder (Solo pa probar)
	fmt.Println("\n\nKey\tValue\tL_Node\tR_Node")
	BT_Root.Root.Print_Inorder()

	// Execute TreeToVine
	BT_Root.DSW_Algorithm()
	fmt.Println("\n\nKey\tValue\tL_Node\tR_Node")
	BT_Root.Root.Print_Inorder()

	// Test Find function with a key that exists
	tupla := BT_Root.Find(27)
	fmt.Print("Resultado de Find para Key = 27 : (")
	fmt.Print(tupla.Found)
	fmt.Print(",")
	fmt.Print(tupla.Num_Comparison)
	fmt.Print(")")
	fmt.Println()

	// Test Find function with a key that doesn't exist
	tupla = BT_Root.Find(50)
	fmt.Print("Resultado de Find para Key = 50 : (")
	fmt.Print(tupla.Found)
	fmt.Print(",")
	fmt.Print(tupla.Num_Comparison)
	fmt.Print(")")
}
