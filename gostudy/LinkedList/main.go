package main

import (
	"fmt"
)

type Node struct {
	next *Node // 다음 노드의 주소값
	val  int   // 현재 노드의 값
}

func main() {
	var root *Node
	var tail *Node

	// Node struct 의 주소값을 저장
	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		// tail 값을 update
		tail = AddNode(tail, i)
	}

	PrintNodes(root)

	root, tail = RemoveNode(root.next, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(root.next, root, tail)

	PrintNodes(root)

}

func AddNode(tail *Node, val int) *Node {

	// tail 에 새 값을 Add 후 tail 반환
	node := &Node{val: val}
	tail.next = node
	return node
}

// 지우고자 하는 노드, Root, Tail
func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {

	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	// 지우고자 하는 노드의 전노드 찾기
	prev := root
	for prev.next != node {
		prev = prev.next
	}

	// 지우고자 하는 노드가 tail인 경우
	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		// 아닌 경우
		// 바로 다음 노드 건너뛰고 참조시키기
		prev.next = prev.next.next
	}

	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d \n", node.val)
}
