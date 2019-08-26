package main

import (
	"fmt"
)

type tree struct {
	value int

	left, right *tree
}

func (t *tree) setValue(v int) {
	t.value = v
}

func (t tree) print() {
	fmt.Printf("%d ", t.value)
}

// 先序
func (t *tree) before_travel() {
	if t == nil {
		return
	}
	t.print()
	t.left.before_travel()
	t.right.before_travel()
}

// 中序
func (t *tree) in_travel() {
	if t == nil {
		return
	}

	t.left.in_travel()
	t.print()
	t.right.in_travel()
}

// 后续
func (t *tree) after_travel() {
	if t == nil {
		return
	}
	t.left.after_travel()
	t.right.after_travel()
	t.print()
}

func main() {

	demo()
}
func demo() {
	node := tree{value: 10}

	node.left = &tree{value: 8}

	node.left.left = &tree{value: 3}

	node.right = &tree{value: 16}

	node.right.left = &tree{value: 13}

	node.right.right = &tree{value: 18}

	node.before_travel()
	fmt.Println()
	node.in_travel()
	fmt.Println()
	node.after_travel()
}
