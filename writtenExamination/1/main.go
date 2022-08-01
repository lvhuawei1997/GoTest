package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTravesal(root *TreeNode) (ans []int) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return
}

func buildTree(preOrder []int, inorder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preOrder[0]}
	var i int
	for index, value := range inorder {
		if value == preOrder[0] {
			i = index
			break
		}
	}
	root.Left = buildTree(preOrder[1:i+1], inorder[:i])
	root.Right = buildTree(preOrder[i+1:], inorder[i+1:])
	return root
}

func main() {
	var num int
	var preOrder, inOrder []int
	for {
		_, err := fmt.Scanf("%d", &num)
		if err == nil {
			preOrder = append(preOrder, num)
		} else {
			break
		}
	}
	for {
		_, err := fmt.Scanf("%d", &num)
		if err == nil {
			inOrder = append(inOrder, num)
		} else {
			break
		}
	}

	fmt.Println(preOrder)
	fmt.Println(inOrder)

	root := buildTree(preOrder, inOrder)
	ans := inorderTravesal(root)
	fmt.Println("ans:", ans)
	//for {
	//	var num []int
	//	reader := bufio.NewReader(os.Stdin)
	//	strBytes, _, _ := reader.ReadLine()
	//	str := strings.Fields(string(strBytes))
	//	for i := range str {
	//		n, _ := strconv.Atoi(str[i])
	//		num = append(num, n)
	//	}
	//	fmt.Println(num)
	//}
	//
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan() //读取一行内容
	//m, _ := strconv.Atoi(strings.Split(input.Text(), " ")[0])
	//n, _ := strconv.Atoi(strings.Split(input.Text(), " ")[1])
	//res := make([][]int, m)
	//for i := range res {
	//	res[i] = make([]int, n)
	//}
	//for i := 0; i < m; i++ {
	//	input.Scan() //读取一行内容
	//	for j := 0; j < n; j++ {
	//		res[i][j], _ = strconv.Atoi(strings.Split(input.Text(), " ")[j])
	//	}
	//}
	//fmt.Println(res)

}
