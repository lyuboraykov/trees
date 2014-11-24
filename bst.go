package main

import "fmt"
import "container/list"
import "math"

const (
	BLACK = false
	RED   = true
)

type treeNode struct {
	parent     *treeNode
	leftChild  *treeNode
	rightChild *treeNode
	key        string
	value      interface{}
	color      bool
}

type BinarySearchTree struct {
	rootNode *treeNode
}

func (bst *BinarySearchTree) Get(key string) (interface{}, bool) {
	if node, ok := bst.get(key, bst.rootNode); ok {
		return node.value, ok
	}
	return nil, false
}

func (bst *BinarySearchTree) get(key string, node *treeNode) (*treeNode, bool) {
	if node == nil {
		return nil, false
	}
	if node.key == key {
		return node, true
	}
	if node.key < key {
		return bst.get(key, node.rightChild)
	}
	// node.key > key
	return bst.get(key, node.leftChild)
}

func (bst *BinarySearchTree) Insert(key string, value interface{}) bool {
	return bst.insert(key, value, &bst.rootNode, nil)
}

func (bst *BinarySearchTree) insert(key string, value interface{},
	node **treeNode, parentNode *treeNode) bool {
	if *node == nil || (*node).key == "" {
		*node = newTreeNode(key, value, parentNode)
		bst.balance(*node)
		return true
	}
	if (*node).key == key {
		return false
	}
	if (*node).key < key {
		return bst.insert(key, value, &(*node).rightChild, *node)
	}
	// node.key > key
	return bst.insert(key, value, &(*node).leftChild, *node)
}

func (bst *BinarySearchTree) balance(node *treeNode) {
	parent := node.parent
	if parent == nil {
		node.color = BLACK
		return
	}
	if parent.color == BLACK {
		return
	}

	uncle := bst.uncle(node)
	grandparent := bst.grandparent(node)
	if parent.color == RED && uncle.color == RED {
		parent.color = BLACK
		uncle.color = BLACK
		grandparent.color = RED
		bst.balance(grandparent)
	}

	if parent.color == RED && uncle.color == BLACK {
		if node == parent.rightChild && parent == grandparent.leftChild {
			bst.rotateLeft(node)
		} else if node == parent.leftChild && parent == grandparent.rightChild {
			bst.rotateRight(node)
		} else if node == parent.leftChild && parent == grandparent.leftChild {
			parent.color = BLACK
			grandparent.color = RED
			bst.rotateRight(parent)
		} else if node == parent.rightChild && parent == grandparent.rightChild {
			parent.color = BLACK
			grandparent.color = RED
			bst.rotateLeft(parent)
		}
	}
	return
}

func (bst *BinarySearchTree) rotateLeft(node *treeNode) {
	savedParent := node.parent
	savedLeft := node.leftChild
	savedGrandparent := bst.grandparent(node)
	node.parent = savedGrandparent
	node.leftChild = savedParent
	savedParent.rightChild = savedLeft
	savedParent.parent = node
	savedLeft.parent = savedParent
	if savedGrandparent != nil {
		if savedParent == savedGrandparent.leftChild {
			savedGrandparent.leftChild = node
		} else {
			savedGrandparent.rightChild = node
		}
	}
}
func (bst *BinarySearchTree) rotateRight(node *treeNode) {
	savedParent := node.parent
	savedGrandparent := bst.grandparent(node)
	savedRight := node.rightChild
	node.rightChild = savedParent
	node.parent = savedGrandparent
	savedParent.parent = node
	savedParent.leftChild = savedRight
	savedRight.parent = savedParent

	if savedGrandparent != nil {
		if savedParent == savedGrandparent.leftChild {
			savedGrandparent.leftChild = node
		} else {
			savedGrandparent.rightChild = node
		}
	}
}

func (bst *BinarySearchTree) uncle(node *treeNode) *treeNode {
	if node.parent == nil {
		return nil
	}
	grandparent := bst.grandparent(node)
	if node.parent == grandparent.leftChild {
		return grandparent.rightChild
	}
	return grandparent.leftChild
}

func (bst *BinarySearchTree) grandparent(node *treeNode) *treeNode {
	if node.parent == nil {
		return nil
	}
	return node.parent.parent
}

func newTreeNode(key string, value interface{}, parent *treeNode) *treeNode {
	node := new(treeNode)
	node.key = key
	node.value = value
	node.parent = parent
	node.leftChild = newLeaf(node)
	node.rightChild = newLeaf(node)
	node.color = RED
	return node
}

func newLeaf(parent *treeNode) *treeNode {
	node := new(treeNode)
	node.parent = parent
	node.leftChild = nil
	node.rightChild = nil
	node.color = BLACK
	return node
}

func NewBinarySearchTree() *BinarySearchTree {
	bst := new(BinarySearchTree)
	bst.rootNode = nil
	return bst
}

func (bst *BinarySearchTree) Delete(key string) bool {
	if node, ok := bst.get(key, bst.rootNode); ok {
		bst.delete(&node)
		return true
	}
	return false
}

func (bst *BinarySearchTree) delete(node **treeNode) {
	if (*node).leftChild == nil && (*node).rightChild == nil {
		*node = nil // TODO: pointer remains in parent, so it is still not deleted
		return
	}
	if (*node).leftChild == nil {
		**node = *(*node).leftChild
		return
	}
	if (*node).rightChild == nil {
		**node = *(*node).rightChild
		return
	}
	*node = (*node).leftChild
	bst.delete(node)
}

func (bst *BinarySearchTree) Draw() {
	counter := 0
	powerOfTwo := 1.0
	queue := list.New()
	fmt.Println(bst.rootNode.key)
	queue.PushBack(bst.rootNode.leftChild)
	queue.PushBack(bst.rootNode.rightChild)
	currentNode := new(treeNode)
	for queue.Len() > 0 {
		currentNode = queue.Front().Value.(*treeNode)
		if currentNode.key != "" {
			fmt.Printf("%s ", currentNode.key)
		} else {
			fmt.Printf("L")
		}
		queue.Remove(queue.Front())
		counter++
		if currentNode.leftChild != nil {
			queue.PushBack(currentNode.leftChild)
			queue.PushBack(currentNode.rightChild)
		}
		if counter == int(math.Pow(2, float64(powerOfTwo))) {
			fmt.Println("")
			counter = 0
			powerOfTwo++
		}
	}
	return
}

func main() {
	bst := NewBinarySearchTree()
	bst.Insert("A", "A")
	bst.Insert("B", "B")
	bst.Insert("C", "C")
	bst.Insert("D", "D")
	bst.Insert("E", "E")
	bst.Insert("F", "F")

	bst.Draw()
}
