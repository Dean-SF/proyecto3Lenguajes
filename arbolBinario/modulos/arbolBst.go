package bst

import "fmt"

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