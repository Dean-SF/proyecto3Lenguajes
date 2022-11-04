package avl

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type BSTNoder interface {
	String()
}

type BSTNode[E constraints.Ordered] struct {
	left,right *BSTNode[E]
	key E
	value int
	compares int

}

func (this *BSTNode[E]) String() string {
	return fmt.Sprintf("(%v, %d)", this.key, this.value)
}

type AVLTreer[E constraints.Ordered] interface {
	insertAux(current *BSTNode[E],key E,compares int)
	rebalanceRight(current *BSTNode[E],compares int)
	rebalanceLeft(current *BSTNode[E],compares int)
	findAux(current *BSTNode[E],key E,compares int)
	rotateRight(current *BSTNode[E],compares int)
	rotateLeft(current *BSTNode[E],compares int)	
	height(current *BSTNode[E],compares int)
	printInorderAux(current *BSTNode[E])
	printInorder()
	weightedComparison(current *BSTNode[E], level int)
	AvgWeightedComparison()
	insert(key E)
	find(key E)
	
	inorder()
}

type AVLTree[E constraints.Ordered] struct {
	Root *BSTNode[E]
}

func (this *AVLTree[E]) weightedComparison(current *BSTNode[E],level int) int {
	if current == nil {
		return 0
	}
	return this.weightedComparison(current.left, level+1) + this.weightedComparison(current.right, level+1) + (current.compares * level)
}

func (this *AVLTree[E]) AvgWeightedComparison() float32 {
	return float32(this.weightedComparison(this.Root,1)) / float32(this.numberNodes(this.Root))
}

func (this *AVLTree[E]) rotateRight(current *BSTNode[E],compares int)(*BSTNode[E],int) {
	compares += 1
	if(current == nil) {
		panic("Can't rotate right on null.")
	} 
	compares += 1
	current.compares += 1
	if(current.left == nil) {
		panic("Can't rotate right with null left child.")
	}
	current.compares += 1
	var temp *BSTNode[E] = current.left
	current.left = temp.right
	temp.right = current
	return temp,compares
}

func (this *AVLTree[E]) rotateLeft(current *BSTNode[E],compares int)(*BSTNode[E],int){
	compares += 1
	if(current == nil) {
		panic("Can't rotate left on null.")
	} 
	compares += 1
	current.compares += 1
	if(current.right == nil) {
		panic("Can't rotate left with null right child.")
	}
	current.compares += 1
	var temp *BSTNode[E] = current.right
	current.right = temp.left
	temp.left = current
	return temp,compares
}

func (this *AVLTree[E]) GetHeight() (int, int) {
	return this.height(this.Root, 0);
}

func (this *AVLTree[E]) GetNumberNodes() int {
	return this.numberNodes(this.Root);
}

func (this *AVLTree[E]) numberNodes(current *BSTNode[E]) int {
	if (current == nil) {
		return 0;
	}
	return this.numberNodes(current.left) + this.numberNodes(current.right) + 1;
}

func (this *AVLTree[E]) height(current *BSTNode[E],compares int) (int,int) {
	compares += 1
	if(current == nil) {
		return 0,compares;
	}
	var leftHeight,comparesLeft int = this.height(current.left,0)
	var rightHeight,comparesRight int = this.height(current.right,0)
	compares += (comparesLeft + comparesRight + 1)
	if(leftHeight > rightHeight) {
		return 1 + leftHeight, compares
	} else {
		return 1 + rightHeight, compares
	}
}

func (this *AVLTree[E]) rebalanceLeft(current *BSTNode[E],compares int)(*BSTNode[E],int){
	var leftHeight,comparesLeft int = this.height(current.left,0)
	var rightHeight,comparesRight int = this.height(current.right,0)
	
	current.compares += 1
	compares += (1 + comparesLeft + comparesRight)
	if(leftHeight - rightHeight > 1) {
		var leftLeftHeight,comparesLeft int = this.height(current.left.left,0)
		var leftRightHeight,comparesRight int = this.height(current.left.right,0)
		compares += (1 + comparesLeft + comparesRight)
		current.compares += 1
		if(leftLeftHeight >= leftRightHeight) {
			return this.rotateRight(current,compares);
		} else {
			current.left,comparesLeft = this.rotateLeft(current.left, 0)
			compares += (comparesLeft)
			return this.rotateRight(current,compares)
		}
	}
	return current,compares
}

func (this *AVLTree[E]) rebalanceRight(current *BSTNode[E],compares int)(*BSTNode[E],int){
	var leftHeight,comparesLeft int = this.height(current.left,0)
	var rightHeight,comparesRight int = this.height(current.right,0)
	compares += (1 + comparesLeft + comparesRight)
	current.compares += 1
	if(rightHeight - leftHeight > 1) {
		var rightRightHeight,comparesLeft int = this.height(current.right.right,0)
		var rightLeftHeight,comparesRight int = this.height(current.right.left,0)
		compares += (1 + comparesLeft + comparesRight)
		current.compares += 1
		if(rightRightHeight >= rightLeftHeight) {
			return this.rotateLeft(current,compares)
		} else {
			current.right,comparesRight = this.rotateRight(current.right,0)
			compares += (comparesRight)
			return this.rotateLeft(current,compares)
		}
	}
	return current,compares
}

func (this *AVLTree[E]) insertAux(current *BSTNode[E],key E,compares int)(*BSTNode[E],int){
	compares += 1
	if current == nil {
		return &BSTNode[E]{left: nil,right: nil,key: key, value: 0,compares: 1},compares
	}
	compares += 1
	current.compares += 2
	if key == current.key {
		current.value += 1
		return current,compares
	}
	compares += 1
	current.compares += 1
	if key < current.key {
		var comparesLeft int
		current.left,comparesLeft = this.insertAux(current.left,key,0)
		compares += comparesLeft
		return this.rebalanceLeft(current,compares)
	} else {
		var comparesRight int
		current.right,comparesRight = this.insertAux(current.right,key,0)
		compares += comparesRight
		return this.rebalanceRight(current,compares)
	}
}

func (this *AVLTree[E]) findAux(current *BSTNode[E],key E, compares int)(bool,int) {
	compares +=1
	if current == nil {
		return false,compares
	}
	compares +=1
	if key == current.key {
		return true,compares
	}
	compares +=1
	if key < current.key {
		return this.findAux(current.left,key,compares)
	} else {
		return this.findAux(current.right,key,compares)
	}
}

func (this *AVLTree[E]) Insert(key E)int {
	var compares int
	this.Root,compares = (this.insertAux(this.Root,key,0))
	return compares
}

func (this *AVLTree[E]) Find(key E) (bool,int) {
	return this.findAux(this.Root,key,0)
}

func (this *AVLTree[E]) printInorderAux(current *BSTNode[E]) {
	if current.left != nil {
		this.printInorderAux(current.left)
	}
	fmt.Println(current)
	if current.right != nil {
		this.printInorderAux(current.right)
	}
}

func (this *AVLTree[E]) PrintInorder() {
	this.printInorderAux(this.Root)
}