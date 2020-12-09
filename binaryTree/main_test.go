package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInOrderSortRecursive(t *testing.T) {
	left121 := &BinaryTree{Value: 5}
	left122 := &BinaryTree{Value: 11}

	left12 := &BinaryTree{Value: 6, Left: left121, Right: left122}
	left11 := &BinaryTree{Value: 2}
	left1 := &BinaryTree{Value: 7, Left: left11, Right: left12}

	right121 := &BinaryTree{Value: 4}
	right12 := &BinaryTree{Value: 9, Left: right121}
	right1 := &BinaryTree{Value: 5, Right: right12}

	root := &BinaryTree{Value: 2, Left: left1, Right: right1}

	assert := []int32{2, 7, 5, 6, 11, 2, 5, 4, 9}
	result := InOrderSortRecursiveStartPoint(root)
	fmt.Println(result)
	if !reflect.DeepEqual(assert, result) {
		t.Error(assert, " not equal ", result)
	}
}

func TestInOrderSortIterative(t *testing.T) {
	left121 := &BinaryTree{Value: 5}
	left122 := &BinaryTree{Value: 11}

	left12 := &BinaryTree{Value: 6, Left: left121, Right: left122}
	left11 := &BinaryTree{Value: 2}
	left1 := &BinaryTree{Value: 7, Left: left11, Right: left12}

	right121 := &BinaryTree{Value: 4}
	right12 := &BinaryTree{Value: 9, Left: right121}
	right1 := &BinaryTree{Value: 5, Right: right12}

	root := &BinaryTree{Value: 2, Left: left1, Right: right1}

	assert := []int32{2, 7, 5, 6, 11, 2, 5, 4, 9}
	result := InOrderSortIterativeStartPoint(root)
	fmt.Println(result)
	if !reflect.DeepEqual(assert, result) {
		t.Error(assert, " not equal ", result)
	}
}
