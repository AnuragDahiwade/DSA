package tree

import "fmt"

func intPtr(v int) *int {
	return &v
}

func main() {
	fmt.Println("Treee!!!!")
	// values := []*int{
	// 	intPtr(1),
	// 	intPtr(2), intPtr(3),
	// 	nil, intPtr(4), nil, nil,
	// }
	// root := BuildTreeFromLevelOrder(values)
	arr := []int{1, 2, -1, -1, 3, 4, -1, -1, 5, -1, -1}
	root := buildBinaryTree(arr)
	res := allTraversals(root)
	fmt.Println(res)
}
