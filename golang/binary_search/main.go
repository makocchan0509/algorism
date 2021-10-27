package main

import (
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	key   int
	value int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func maketree(parentNode *node, newKey int, newValue int) {

	var childNode node
	childNode.key = newKey
	childNode.value = newValue

	if newValue >= parentNode.value {
		if parentNode.right == nil {
			parentNode.right = &childNode
		} else {
			maketree(parentNode.right, newKey, newValue)
		}
	} else {
		if parentNode.left == nil {
			parentNode.left = &childNode
		} else {
			maketree(parentNode.left, newKey, newValue)
		}
	}
}

func findNumber(targetNode *node, target int) int {
	if target == targetNode.value {
		return target
	} else if target >= targetNode.value {
		if targetNode.right == nil {
			return -1
		} else {
			return findNumber(targetNode.right, target)
		}
	} else {
		if targetNode.left == nil {
			return -1
		} else {
			return findNumber(targetNode.left, target)
		}
	}
}

func main() {
	maxValue := 20
	var valueArray = []int{maxValue / 2}
	fmt.Println(valueArray)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < maxValue; i++ {
		valueArray = append(valueArray, rand.Intn(maxValue))
	}
	fmt.Println(valueArray)

	var myTree tree
	var firstNode node
	myTree.root = &firstNode

	for i := 0; i < len(valueArray); i++ {
		maketree(myTree.root, i, valueArray[i])
	}
	descTree(myTree.root, 0)

	fmt.Println(findNumber(myTree.root, 10))
}

func descTree(targetNode *node, depth int) {

	fmt.Println(depth, targetNode.key, targetNode.value)
	if targetNode.left != nil {
		descTree(targetNode.left, depth+1)
	}
	if targetNode.right != nil {
		descTree(targetNode.right, depth+1)
	}
}

func deleteNode(node *node, target int) *node {
	if node != nil {
		if node.value == target {
			if node.left == nil {
				return node.right
			} else if node.right == nil {
				return node.left
			} else {
				node.value = searchMin(node.right)
				node.right = deleteMin(node.right)
			}
		} else if node.value > target {
			node.left = deleteNode(node.left, target)
		} else {
			node.right = deleteNode(node.right, target)
		}
	}
	return node
}
func searchMin(node *node) int {
	if node.left == nil {
		return node.value
	}
	return searchMin(node.left)
}

func deleteMin(node *node) *node {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	return node
}
