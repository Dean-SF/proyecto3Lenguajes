package avl

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BSTNoder interface {

}

type BSTNode[E constraints.Ordered] struct {
	left,right *BSTNode[E]
	value E

}

type AVLTreer[E constraints.Ordered] interface {
	insertAux(current *BSTNode[E],element E);
	insert(element E)
	rebalanceLeft(current *BSTNode[E])
	rebalanceRight(current *BSTNode[E])
	inorder()
}

type AVLTree[E constraints.Ordered] struct {
	root *BSTNode[E]
}

func (this *AVLTree[E]) rebalanceLeft(current *BSTNode[E])*BSTNode[E]{
	retorno := new(BSTNode[E])
	*retorno = BSTNode[E]{left: nil,right: nil}
	return retorno;
}

func (this *AVLTree[E]) rebalanceRight(current *BSTNode[E])*BSTNode[E]{
	retorno := new(BSTNode[E])
	*retorno = BSTNode[E]{left: nil,right: nil}
	return retorno;
}

func (this *AVLTree[E]) insertAux(current *BSTNode[E],element E)*BSTNode[E]{
	if current == nil {
		return &BSTNode[E]{left: nil,right: nil,value: element};
	}
	if element < current.value {
		current.left = this.insertAux(current.left,element)
		return this.rebalanceLeft(current)
	} else {
		current.right = this.insertAux(current.right,element)
		return this.rebalanceRight(current)
	}
}

func (this *AVLTree[E]) insert(element E) {
	this.root = (this.insertAux(this.root,element));
}

func Pruebas() {
	var arbol = AVLTree[int]{nil};
	arbol.insert(6);
	fmt.Println(arbol.root);
}