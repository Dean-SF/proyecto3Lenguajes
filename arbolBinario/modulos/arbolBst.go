package bst

import (
	"fmt"
	"math"
)

// -------------------------------------------
// --	Structure of a Binary Search Tree	--
// -------------------------------------------

// Struct: Node of a Binary Tree
type BT_Node struct {
	Key        int
	Counter    int
	Left_node  *BT_Node
	Right_node *BT_Node
}

// Struct: Root of Binary Tree
type BinaryTree struct {
	Root *BT_Node
}

// ---------------------------
// --	Structure of Tuple	--
// ---------------------------

type Tuple struct {
	Found          bool
	Num_Comparison int
}

// -----------------------------------------------
// --	Find Functions of a Binary Search Tree	--
// -----------------------------------------------

// Accesible function to search a given key and return the result and the amount of comparisons
func (this *BinaryTree) Find(key int) (tuple *Tuple) {
	// Initialize values
	num_Comp := 0
	foundBool := false

	// Tuple with final result
	tupla := Tuple{Found: foundBool, Num_Comparison: num_Comp}

	if this.Root == nil {
		tupla.Found = false
		tupla.Num_Comparison += 1
		return &tupla
	} else {
		findAux(&this.Root, key, &tupla)
	}
	return &tupla
}

// Recursive auxiliary function that updates the tuple's values while searching for the key
func findAux(this **BT_Node, key int, tupla *Tuple) (tuple *Tuple) {
	if (*this) == nil {
		(*&tupla.Num_Comparison) += 1
		(*&tupla.Found) = false

	} else if key == (*this).Key {
		(*&tupla.Num_Comparison) += 2
		(*&tupla.Found) = true

	} else if key < (*this).Key {
		(*&tupla.Num_Comparison) += 3
		findAux(&((*this).Left_node), key, tupla)

	} else {
		(*&tupla.Num_Comparison) += 3
		findAux(&((*this).Right_node), key, tupla)
	}
	return
}

// ---------------------------------------------------
// --	Insertion Functions of a Binary Search Tree	--
// ---------------------------------------------------

// Accesible function to insert element in the tree -> Return the number of comparison made
func (this *BinaryTree) Insert(key int) (num_Comparison int) {
	num_Comparison = 1
	if this.Root == nil {
		this.Root = &BT_Node{key, 1, nil, nil}
	} else {
		insertNode(&this.Root, key, &num_Comparison)
	}
	return
}

// Inside function to perform the recursive process
func insertNode(this **BT_Node, new_key int, num_Comparison *int) {
	if (*this) == nil {
		(*num_Comparison) += 1
		*this = &BT_Node{new_key, 1, nil, nil}

	} else if new_key == (*this).Key {
		(*num_Comparison) += 2
		(*this).Counter++

	} else if new_key < (*this).Key {
		(*num_Comparison) += 3
		insertNode(&((*this).Left_node), new_key, num_Comparison)

	} else {
		(*num_Comparison) += 3
		insertNode(&((*this).Right_node), new_key, num_Comparison)
	}
}

// -----------------------------------
// --	Height Binary Search Tree	--
// -----------------------------------

func (this *BinaryTree) GetHeight() int {
	return this.Root.height();
}

func (this *BT_Node) height() int {
	if (this == nil) {
		return 0;
	}

	var H_Left int = this.Left_node.height() + 1;
	var H_Right int = this.Right_node.height() + 1;

	if (H_Left > H_Right) {
		return H_Left;
	} else {
		return H_Right;
	}
}

// ---------------------------------------------------
// --	Number of nodes of a  Binary Search Tree	--
// ---------------------------------------------------

func (this *BinaryTree) GetNumberNodes() int {
	return this.Root.numberNodes();
}

func (this *BT_Node) numberNodes() int {
	if (this == nil) {
		return 0;
	}
	return this.Left_node.numberNodes() + this.Right_node.numberNodes() + 1;
}

// -----------------------------------
// --	Print a Binary Search Tree	--
// -----------------------------------

// Test function: Show every value of every node in the tree
func (this *BT_Node) Print_Inorder() {
	if this != nil {
		var L_Value int = -1
		var R_Value int = -1

		if this.Left_node != nil {
			L_Value = this.Left_node.Key
		}
		if this.Right_node != nil {
			R_Value = this.Right_node.Key
		}

		this.Left_node.Print_Inorder()
		fmt.Printf("%v\t%v\t%v\t%v\n", this.Key, this.Counter, L_Value, R_Value)
		this.Right_node.Print_Inorder()
	}
}

// -------------------------------------------------------
// --	Implementation of Day-Stout-Warren Algorithm	--
// -------------------------------------------------------
/*
	Tomado de:
	> Wikipedia(2022) Day–Stout–Warren algorithm [Pseudo code]: https://en.wikipedia.org/wiki/Day%E2%80%93Stout%E2%80%93Warren_algorithm
	> jayshilbuddhadev-GeeksForGeeks(2022) Day-Stout-Warren algorithm to balance given Binary Search Tree [Source code]:
		https://www.geeksforgeeks.org/day-stout-warren-algorithm-to-balance-given-binary-search-tree/
*/

// Execute the process of DSW algorithm
func (this *BinaryTree) DSW_Algorithm() {
	// Create a 'PseudoRoot' whose right child is the original root
	var PsedoRoot *BT_Node = &BT_Node{ 0, 0, nil, this.Root};

	// Make a Only-Right-Childs Tree (Linked list)
	var num_Nodes = TreeToVine(PsedoRoot);

	// Create a Complete Binary Tree (Route Balance BT) out of the pseudo root
	VineToTree(PsedoRoot, num_Nodes);

	// Make the right child of the pseudo root, the actual tree root
	this.Root = PsedoRoot.Right_node;
}

// Apply right-rotations on the tree to make it a Linked List
func TreeToVine(node *BT_Node) (num_Nodes int) {
	var tail *BT_Node = node;
	var rest *BT_Node = node.Right_node;
	var temp *BT_Node;

	for ; rest != nil; {
		if (rest.Left_node == nil) {
			tail = rest;
			rest = rest.Right_node;
			num_Nodes++;
		} else {
			temp = rest.Left_node;
			rest.Left_node = temp.Right_node;
			temp.Right_node = rest;
			rest = temp;
			tail.Right_node = temp;
		}
	}
	return;
}

// Transform the linked list into a Balanced Binary Tree
func VineToTree(node *BT_Node, num_Nodes int) {

	// Get the max depth that a tree can have with n amount of nodes
	var max_Depth = math.Trunc(math.Log2(float64(num_Nodes + 1)));

	// Get the amount of leaves the tree will have after the first compression
	var m = int(math.Pow(2, max_Depth) - 1);

	Compress(node, num_Nodes - m);

	for m = m / 2; m > 0; m /= 2{
		Compress(node, m)
	}
}

// Compress process: Make the odd nodes the left child of the even nodes
func Compress(root *BT_Node, count int) {
	var temp *BT_Node = root.Right_node;
	var oldTemp *BT_Node;

	for i := 0; i < count; i++ {
		oldTemp = temp;
		temp = temp.Right_node;
		root.Right_node = temp;
		oldTemp.Right_node = temp.Left_node;
		temp.Left_node = oldTemp;
		root = temp;
		temp = temp.Right_node;
	}
}