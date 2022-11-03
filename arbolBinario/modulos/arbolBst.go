package bst

import (
	"fmt"
	"math"
)

// Struct: Node of a Binary Tree
type BT_Node struct {
	Key int;
	Counter int;
	Left_node  *BT_Node;
	Right_node *BT_Node;
}

// Struct: Root of Binary Tree
type BinaryTree struct {
	Root *BT_Node;
}

// Accesible function to insert element in the tree -> Return the number of comparison made
func (this *BinaryTree) Insert(key int) (num_Comparison int){
	num_Comparison = 1;
	if (this.Root == nil) {
		this.Root = &BT_Node{ key, 1, nil, nil };
	} else {
		insertNode(&this.Root, key, &num_Comparison);
	}
	return;
}

// Inside function to perform the recursive process
func insertNode(this **BT_Node, new_key int, num_Comparison *int) {
	if ((*this) == nil) {
		(*num_Comparison) += 1;
		*this = &BT_Node{ new_key, 1, nil, nil };

	} else if(new_key == (*this).Key) {
		(*num_Comparison) += 2;
		(*this).Counter++;

	} else if(new_key < (*this).Key) {
		(*num_Comparison) += 3;
		insertNode( &((*this).Left_node), new_key, num_Comparison );

	} else {
		(*num_Comparison) += 3;
		insertNode( &((*this).Right_node), new_key, num_Comparison );
	}
}

func (this *BinaryTree) DSW_Algorithm() {
	var tempRoot *BT_Node = &BT_Node{ 0, 0, nil, this.Root};
	var num_Nodes = TreeToVine(tempRoot);
	
	// VineToTree process -> Calls Compress
	var h = int(math.Log2(float64(num_Nodes + 1)));
	var m = int(math.Pow(2, float64(h)) - 1);

	Compress(tempRoot, num_Nodes - m);

	for m = m / 2; m > 0; m /= 2{
		Compress(tempRoot, m)
	}

	this.Root = tempRoot.Right_node;
}

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

// Test function: Show every value of every node in the tree
func (this *BT_Node) Print_Inorder() {
	if (this != nil) {
		var L_Value int = -1;
		var R_Value int = -1;

		if (this.Left_node != nil) {
			L_Value = this.Left_node.Key;
		}
		if (this.Right_node != nil) {
			R_Value = this.Right_node.Key;
		}

		this.Left_node.Print_Inorder();
		fmt.Printf("%v\t%v\t%v\t%v\n", this.Key, this.Counter, L_Value, R_Value);
		this.Right_node.Print_Inorder();
	}
}