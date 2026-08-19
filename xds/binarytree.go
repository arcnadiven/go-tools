package xds

import (
	fmt "fmt"
	"github.com/arcnadiven/go-tools/tools"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = "null"

func (root *TreeNode) AddLeftNode(val int) {
	if root.Left == nil {
		root.Left = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	} else {
		root.Left.Val = val
	}
}

func (root *TreeNode) AddRightNode(val int) {
	if root.Right == nil {
		root.Right = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	} else {
		root.Right.Val = val
	}
}

func GetMaxNodeNum(col int) int {
	return tools.Pow(2, col) - 1
}

func GetBinaryTreeDepth(nodeNum int) int {
	count := 0
	for {
		if nodeNum > GetMaxNodeNum(count) && nodeNum <= GetMaxNodeNum(count+1) {
			return count + 1
		}
		count++
	}
}

// LeetCode用，levelOrder序列化，节点缺失时使用null代替
func MarshalBinaryTree(src []interface{}) *TreeNode {
	if len(src) == 0 {
		return nil
	}
	if src[0] == nil {
		return nil
	}
	root := &TreeNode{}
	if val, ok := src[0].(float64); ok {
		root.Val = int(val)
	} else {
		panic("root node value is not a number")
	}
	currentLevel := []*TreeNode{root} //level Order需要维护一个临时队列
	srcIdx := 1
	for srcIdx < len(src) {
		maxIdx := 0
		for curIdx, node := range currentLevel {
			//尝试分配左子节点
			if srcIdx == len(src) { //判断是否越界
				break
			}
			{
				if val, ok := src[srcIdx].(float64); ok { //判断是否为int,有int则补一位
					lChild := &TreeNode{Val: int(val)} //那我怎么关联上源地址？？？
					currentLevel = append(currentLevel, lChild)
					node.Left = lChild
				} else {
					if src[srcIdx] == nil {
					} else {
						panic(fmt.Sprintf("node: %d left child value is not number", node.Val))
					}
				}
			}
			srcIdx++

			//尝试分配右子节点
			if srcIdx == len(src) { //判断是否越界
				break
			}
			{
				if val, ok := src[srcIdx].(float64); ok { //判断是否为int,有int则补一位
					rChild := &TreeNode{Val: int(val)} //那我怎么关联上源地址？？？
					currentLevel = append(currentLevel, rChild)
					node.Right = rChild
				} else {
					if src[srcIdx] == nil {
					} else {
						panic(fmt.Sprintf("node: %d right child value is not number", node.Val))
					}
				}
			}
			srcIdx++

			//更新currentLevel临时队列的下标
			maxIdx = curIdx
		}
		currentLevel = currentLevel[maxIdx+1:]
	}
	return root
}

// 自动填充null
func LevelOrder(root *TreeNode) []interface{} {
	result := []interface{}{}
	if root == nil {
		return nil
	}
	lastLevel := []*TreeNode{root}
	for len(lastLevel) != 0 {
		lastIndex := 0
		for idx, node := range lastLevel {
			if node == nil {
				result = append(result, NULL)
				continue
			}
			result = append(result, node.Val)
			lastLevel = append(lastLevel, node.Left)
			lastLevel = append(lastLevel, node.Right)
			lastIndex = idx
		}
		lastLevel = lastLevel[lastIndex+1:]
	}
	idx := 0
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != NULL {
			idx = i
			break
		}
	}
	return result[:idx+1]
}
