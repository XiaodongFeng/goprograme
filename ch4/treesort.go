package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	values := []int{34, 22, 14, 78, 1, 4, 3, 22, 17, 91, 45}
	fmt.Println(values)
	Sort(values)
	fmt.Println(values)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, val int) *tree {
	if t == nil {
		t := new(tree)
		t.value = val
		return t
	}

	if val < t.value {
		t.left = add(t.left, val)
	} else {
		t.right = add(t.right, val)
	}
	return t
}
