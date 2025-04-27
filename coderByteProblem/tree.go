package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeConstruct struct {
	val   int
	left  *TreeConstruct
	right *TreeConstruct
}

func BuildBinaryTree(str []string) (*TreeConstruct, bool) {
	node := make(map[int]*TreeConstruct)
	childtoparent := make(map[int]int)
	parenttochild := make(map[int][]int)
	for _, s := range str {
		remove := strings.Trim(s, "()")
		replace := strings.Split(remove, ",")
		fmt.Println(len(replace))
		if len(replace) != 2 {
			return nil, false
		}
		childval, err1 := strconv.Atoi(replace[0])
		parentval, err2 := strconv.Atoi(replace[1])
		if err1 != nil || err2 != nil {
			return nil, false
		}
		if _, exist := node[childval]; !exist {
			node[childval] = &TreeConstruct{val: childval}
		}
		if _, exist := node[parentval]; !exist {
			node[parentval] = &TreeConstruct{val: parentval}
		}
		if _, exist := childtoparent[childval]; exist {
			return nil, false
		}
		childtoparent[childval] = parentval
		parenttochild[parentval] = append(parenttochild[parentval], childval)
		if len(parenttochild[parentval]) > 2 {
			return nil, false
		}
		parent := node[parentval]
		child := node[childval]
		if parent.left == nil {
			parent.left = child
		} else if parent.right == nil {
			parent.right = child
		} else {
			return nil, false
		}
	}
	var root *TreeConstruct
	var rootCount int = 0
	for val, node := range node {
		if _, exit := childtoparent[val]; !exit {
			root = node
			rootCount++
		}
	}
	if rootCount != 1 {
		return nil, false
	}

	return root, true
}
func main() {
	input := []string{"(1,2)", "(2,4)", "(5,7)", "(7,2)", "(9,5)"}
	root, isValid := BuildBinaryTree(input)
	fmt.Println(root, isValid)
}
