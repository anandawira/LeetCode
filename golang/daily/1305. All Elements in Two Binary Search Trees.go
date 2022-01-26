package daily

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorderTraversal(result []int, root *TreeNode) []int {
	if root == nil {
		return result
	}

	result = inorderTraversal(result, root.Left)
	result = append(result, root.Val)
	result = inorderTraversal(result, root.Right)
	return result
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	list1 := inorderTraversal([]int{}, root1)
	list2 := inorderTraversal([]int{}, root2)

	var index1, index2 int

	var result []int
	for index1 != len(list1) && index2 != len(list2) {
		if list1[index1] < list2[index2] {
			result = append(result, list1[index1])
			index1++
		} else {
			result = append(result, list2[index2])
			index2++
		}
	}

	result = append(result, list1[index1:]...)
	result = append(result, list2[index2:]...)

	return result
}
