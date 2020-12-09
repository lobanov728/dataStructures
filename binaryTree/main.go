package main

type BinaryTree struct {
	Left  *BinaryTree
	Right *BinaryTree
	Value int32
}

func main() {
}

func InOrderSortRecursiveStartPoint(tree *BinaryTree) []int32 {
	var totalResult []int32
	InOrderSortRecursive(tree, &totalResult)
	return totalResult
}

func InOrderSortRecursive(tree *BinaryTree, result *[]int32) {
	if tree.Left != nil {
		InOrderSortRecursive(tree.Left, result)
	}
	if tree.Left == nil && tree.Right == nil {
		*result = append(*result, tree.Value)
		return
	}
	*result = append(*result, tree.Value)
	if tree.Right != nil {
		InOrderSortRecursive(tree.Right, result)
	}
}

func InOrderSortIterativeStartPoint(tree *BinaryTree) []int32 {
	var totalResult []int32
	InOrderSortIterative(tree, &totalResult)
	return totalResult
}

func InOrderSortIterative(tree *BinaryTree, result *[]int32) {
	var stack []BinaryTree
	for ;len(stack) > 0 || tree != nil; {
		if tree != nil {
			stack = append(stack, *tree)
			tree = tree.Left
		} else {
			tree, stack = &stack[len(stack)-1], stack[:len(stack)-1]
			*result = append(*result, tree.Value)
			tree = tree.Right
		}
	}
}
