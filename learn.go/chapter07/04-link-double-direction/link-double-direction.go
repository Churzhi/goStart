package main

import "fmt"

type LinkNode struct {
	data     int
	next     *LinkNode
	previous *LinkNode
}

func buildDLink() *LinkNode {
	n1 := &LinkNode{data: 1}
	n2 := &LinkNode{data: 1}
	n3 := &LinkNode{data: 1}

	n1.next = n2
	n2.previous = n1

	n2.next = n3
	n3.previous = n2

	return n1
}

func insertNode(root *LinkNode, newNode *LinkNode) *LinkNode {
	tmpNode := root

	// 给空链表添加节点
	if root == nil {
		return newNode
	}

	// 在链表头添加节点
	if root.data >= newNode.data {
		newNode.next = tmpNode
		tmpNode.previous = newNode
		return newNode
	}

	for {
		if tmpNode.next == nil {
			// 已经到头，追加节点即可
			tmpNode.next = newNode
			newNode.previous = tmpNode
			return root
		} else {
			if tmpNode.next.data >= newNode.data {
				newNode.previous = tmpNode
				newNode.next = tmpNode.next
				tmpNode.next = newNode
				newNode.next.previous = newNode
				return root

			}
		}
		tmpNode = tmpNode.next
	}
}

func deleteNode(root *LinkNode, v int) *LinkNode {
	if root == nil {
		return nil

	}
	// 如果需要删除的数据在头节点
	if root.data == v {
		leftHand := root
		root = root.next

		leftHand.next = nil
		root.previous = nil

		// TODO 需要解决只有一个节点的情况
		return root

	}

	tmpNode := root
	for {
		if tmpNode.next == nil {
			// 走到链表尾部，仍然没有找到要删除的数据，直接返回原 root
			return root
		} else {
			if tmpNode.next.data == v {
				// 找到节点 开始删除
				// TODO
				rightHand := tmpNode.next
				tmpNode.next = rightHand.next
				rightHand.next.previous = tmpNode

				// 清理掉右手上的
				return root
			}
		}
	}
}

func rangeLink(root *LinkNode) {
	fmt.Println("从头到尾：")
	tmpNode := root
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}

	fmt.Println("从尾到头：")
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.previous == nil {
			break
		}

	}
}

func main() {

	root := buildDLink()
	// 需要及时变更 root 的位置
	root = insertNode(root, &LinkNode{data: 3})
	root = insertNode(root, &LinkNode{data: 7})
	root = insertNode(root, &LinkNode{data: 11})
	root = insertNode(root, &LinkNode{data: 0})
	rangeLink(root)

	fmt.Println("删除节点")
	root = deleteNode(root, 7)
}
