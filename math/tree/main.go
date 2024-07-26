package main
import (
	"fmt"
)
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func NewTreeNode(v int) *TreeNode{
	return  &TreeNode{
        Left:  nil, // 左子节点指针
        Right: nil, // 右子节点指针
        Val:   v,   // 节点值
    }
}
//将数组转变为二叉树
func list_to_tree(arr []int) *TreeNode{
		return list_to_tree_dfs(arr,0)
}
func list_to_tree_dfs(arr []int,i int) *TreeNode{
	if i>=len(arr) {
		return nil
	}
	root := &TreeNode{Val: arr[i]}
	root.Left = list_to_tree_dfs(arr,2*i+1)
	root.Right = list_to_tree_dfs(arr,2*i+2)
	return root
}
/* 层序遍历 */
func levelOrder(root *TreeNode) {
	queue :=[]*TreeNode{root}
    // 初始化队列，加入根节点
	for len(queue)>0 {
		node :=queue[0]
		queue = queue[1:]
		fmt.Print(node.Val," ")
		if node.Left !=nil{
			queue = append(queue, node.Left)
		}
		if node.Right !=nil{
			queue = append(queue, node.Right)
		}
	}
    
}
//前序遍历
func pre_order(root *TreeNode){
	if root == nil{
		return
	}
	fmt.Print(root.Val," ")
	pre_order(root.Left)
	pre_order(root.Right)
}
//中序遍历
func in_order(root *TreeNode){
	if root == nil{
		return
	}
	in_order(root.Left)
	fmt.Print(root.Val," ")
	in_order(root.Right)
}
//后序遍历
func post_order(root *TreeNode){
	if root == nil{
		return
	}
	post_order(root.Left)
	post_order(root.Right)
	fmt.Print(root.Val," ")
}
func main(){
	t1 := NewTreeNode(1)
	t2 := NewTreeNode(2)
	t3 := NewTreeNode(3)
	t4 := NewTreeNode(4)
	t5 := NewTreeNode(5)
	//构建节点之间的引用
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	fmt.Println(t1.Val)
	//插入
	p := NewTreeNode(0)
	p.Left = t2
	t1.Left = p
	//删除
	t1.Left = t2
	//创建数组转变成二叉树
	arr := []int{1,3,4,2,6}
	root := list_to_tree(arr)
	//遍历二叉树
	//层序遍历
	levelOrder(root)
	fmt.Println()
	levelOrder(t1)
	fmt.Println()
	pre_order(t1)
	fmt.Println()
	in_order(t1)
	fmt.Println()
	post_order(t1)

}