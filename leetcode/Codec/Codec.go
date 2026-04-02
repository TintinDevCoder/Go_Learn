package main

// 297. 二叉树的序列化与反序列化
import (
	"fmt"
	tree "leetcode/TreeNode"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type tree.TreeNode struct {
 *     Val int
 *     Left *tree.TreeNode
 *     Right *tree.TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	c := Codec{}
	return c
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *tree.TreeNode) string {
	var builder strings.Builder
	var dfs func(*tree.TreeNode)
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			builder.WriteString("n,")
			return
		}
		builder.WriteString(strconv.Itoa(node.Val) + (","))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *tree.TreeNode {
	if data == "" {
		return nil
	}
	nodes := strings.Split(data, ",")
	numNodes := len(nodes)
	if len(nodes) > 0 && nodes[len(nodes)-1] == "" {
		nodes = nodes[:len(nodes)-1]
	}
	index := 0
	var dfs func() *tree.TreeNode
	dfs = func() *tree.TreeNode {
		if index >= numNodes || nodes[index] == "n" {
			index++
			return nil
		}
		v, _ := strconv.Atoi(nodes[index])
		index++

		return &tree.TreeNode{Val: v, Left: dfs(), Right: dfs()}
	}
	return dfs()
}
func main() {
	node1 := &tree.TreeNode{Val: 1}
	node2 := &tree.TreeNode{Val: 2}
	node3 := &tree.TreeNode{Val: 3}
	node4 := &tree.TreeNode{Val: 4}
	node5 := &tree.TreeNode{Val: 5}

	// 2. 建立层级关系 (Establish hierarchy)
	node1.Left = node2
	node1.Right = node3

	node3.Left = node4
	node3.Right = node5

	constructor := Constructor()
	serialize := constructor.serialize(node1)
	n := constructor.deserialize(serialize)
	fmt.Println(n)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
