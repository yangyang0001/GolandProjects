package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Node struct {
	Name  string `json:"name"`
	Left  *Node  `json:"left"`
	Right *Node  `json:"right"`
}

func main() {

	// fmt.Println("-------------------------- 1、构造 二叉树  --------------------------")

	left1_left := Node{"left1_left", nil, nil}
	left1_righ := Node{"left1_righ", nil, nil}

	righ1_left := Node{"righ1_left", nil, nil}
	righ1_righ := Node{"righ1_righ", nil, nil}

	left1 := Node{"left1", &left1_left, &left1_righ}
	righ1 := Node{"righ1", &righ1_left, &righ1_righ}

	root := Node{"root", &left1, &righ1}

	fmt.Println("-------------------------- 2、二叉树 先根遍历 --------------------------")
	RootBegin(&root)

	fmt.Println("-------------------------- 3、二叉树 中根遍历 --------------------------")
	RootMidd(&root)

	fmt.Println("-------------------------- 4、二叉树 后根遍历 --------------------------")
	RootLast(&root)

	fmt.Println("-------------------------- 5、二叉树 自上而下 自左向右 遍历 --------------")
	TopBegin([]*Node{&root})
}

func RootBegin(root *Node) {

	if root == nil || root.Name == "" {
		return
	} else {
		fmt.Println(root.Name)
	}

	RootBegin(root.Left)
	RootBegin(root.Right)

}

func RootMidd(root *Node) {

	if root == nil || root.Name == "" {
		return
	}

	RootMidd(root.Left)

	fmt.Println(root.Name)

	RootMidd(root.Right)

}

func RootLast(root *Node) {

	if root == nil || root.Name == "" {
		return
	}

	RootLast(root.Left)

	RootLast(root.Right)

	fmt.Println(root.Name)

}

func TopBegin(roots []*Node) {

	if roots == nil || len(roots) == 0 {
		return
	}

	var children []*Node

	for _, root := range roots {
		if root == nil && root.Name == "" {
			return
		} else {
			fmt.Println(root.Name)
		}

		if root.Left != nil && root.Left.Name != "" {
			children = append(children, root.Left)
		}
		if root.Right != nil && root.Right.Name != "" {
			children = append(children, root.Right)
		}

	}

	TopBegin(children)

}

func jsonParse(obj any) string {

	marshal, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)
}
