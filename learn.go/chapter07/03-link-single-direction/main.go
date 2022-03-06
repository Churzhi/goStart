package main

import "fmt"

// LinkNode 单链表
type LinkNode struct {
	data int
	next *LinkNode
}

func main() {

	n1 := &LinkNode{
		data: 1,
		next: nil,
	}
	n2 := &LinkNode{
		data: 2,
		next: nil,
	}
	n3 := &LinkNode{
		data: 3,
		next: nil,
	}
	n4 := &LinkNode{
		data: 4,
		next: nil,
	}
	n6 := &LinkNode{
		data: 6,
		next: nil,
	}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6

	{
		rangeLink(n1)
	}

	{
		fmt.Println("插入数据5")

		// 遍历节点: 需要从头开始遍历
		n5 := &LinkNode{
			data: 5,
			next: nil,
		}
		insertNode(n1, n5)
		insertNode(n1, &LinkNode{
			data: 7,
			next: nil,
		})
		rangeLink(n1)
		{
			fmt.Println("删除节点")
			n1 = deketeNode(n1, 3)
			n1 = deketeNode(n1, 1)
			n1 = deketeNode(n1, 7)
			rangeLink(n1)
		}
	}
}

func deketeNode(root *LinkNode, data int) *LinkNode {
	tmpNode := root
	// 如果目标节点为第一个节点
	if root != nil && root.data == data {
		if root.next == nil {
			return nil
		}
		right := root.next
		tmpNode.next = nil
		return right
	}
	for {
		if tmpNode.next == nil {
			break
		}
		right := tmpNode.next
		if right.data == data {
			//	找到要删除的节点
			tmpNode.next = right.next // tmpNode.next.next
			right.next = nil
			return root
		}
		// 没找到目标节点则往下走
		tmpNode = tmpNode.next
	}
	return root

}

func insertNode(root *LinkNode, newNode *LinkNode) {
	tmpNode := root // 头节点
	for {
		if tmpNode != nil {
			//fmt.Println(tmpNode.data)
			if newNode.data > tmpNode.data {
				if tmpNode.next == nil {
					// 已经到尾节点，直接添加新节点
					tmpNode.next = newNode
				} else {
					if tmpNode.next.data > newNode.data {
						//	找到合适的位置，插入新节点
						newNode.next = tmpNode.next
						tmpNode.next = newNode
					}
				}
			}

		} else {
			break
		}
		tmpNode = tmpNode.next
	}
}

/* func rangeLink(root *LinkNode, f func(data interface{})) {}
 */
func rangeLink(root *LinkNode) {
	tmpNode := root // 头节点
	// 遍历节点: 需要从头开始遍历
	for {
		if tmpNode != nil {
			//f(tmpNode.data)
			fmt.Println(tmpNode.data)
		} else {
			break
		}
		tmpNode = tmpNode.next
	}
}
