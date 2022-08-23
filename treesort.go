package sorts

/*
	type TypeNumber interface {
		int | int64 | float64
	}
*/

type treeNode [T TypeNumber] struct {
	Val T
	Left *treeNode[T]
	Right *treeNode[T]
}

func Tree[T TypeNumber](arr []T) []T {
	ans := make([]T, 0)
	var root *treeNode[T]
	for _, v := range arr {
		root = add(root, v)
	}
	readTree(root, &ans)
	return ans
}

func add[T TypeNumber](node *treeNode[T], value T) *treeNode[T] {
	if node == nil {
		node = &treeNode[T]{}
		node.Val = value
		return node
	}

	if value < node.Val {
		node.Left = add(node.Left, value)
	} else {
		node.Right = add(node.Right, value)
	}
	return node
}

func readTree[T TypeNumber](node *treeNode[T], arr *[]T)  {
	if node == nil {
		return
	}

	if node.Left != nil {
		readTree(node.Left, arr)
	}
	*arr = append(*arr, node.Val)

	if node.Right != nil {
		readTree(node.Right, arr)
	}
}
