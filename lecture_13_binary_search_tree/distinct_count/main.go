package main

/*
Create a binary search tree
add node to BST
count the BST
*/

type node struct {
	key   int
	left  *node
	right *node
}

func NewBinarySearchTree(key int) *node {
	return &node{key: key, left: nil, right: nil}
}

// O(h) h is the height of the tree, worst case O(n)
func Insert(root *node, key int) *node {
	if root == nil {
		return NewBinarySearchTree(key)
	}
	if key < root.key {
		root.left = Insert(root.left, key)
	} else if key > root.key {
		root.right = Insert(root.right, key)
	}
	return root
}

func main() {
	// time limit exceed on hashmap

}
