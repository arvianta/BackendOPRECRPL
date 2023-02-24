package main

import "fmt"

type slistnode struct {
	key  int
	next *slistnode
}

type slist struct {
	size int
	head *slistnode
}

func (list *slist) slist_pushBack(value int) {
	n := slistnode{}
	n.key = value
	if list.size == 0 {
		list.head = &n
		list.size++
		return
	}
	ptr := list.head
	for i := 0; i < list.size; i++ {
		if ptr.next == nil {
			ptr.next = &n
			list.size++
			return
		}
		ptr = ptr.next
	}
}

func (list *slist) slist_popBack() {
	if list.size > 0 {
		nextNode := list.head.next
		currNode := list.head
		if currNode.next == nil {
			list.head = nil
		}

		for nextNode.next != nil {
			currNode = nextNode
			nextNode = nextNode.next
		}
		currNode.next = nil
		list.size--
	}
}

func (list *slist) slist_printlist() {
	if list.size == 0 {
		fmt.Println("The list is empty")
	}
	ptr := list.head
	fmt.Print("List : ", ptr.key)
	ptr = ptr.next
	for i := 0; i < list.size-2; i++ {
		fmt.Print(" -> ", ptr.key)
		ptr = ptr.next
	}
	fmt.Print(" -> ", ptr.key, "\n")
}

func main() {
	var list slist

	for i := 1; i <= 10; i++ {
		list.slist_pushBack(i)
	}

	list.slist_printlist()
	list.slist_popBack()
	list.slist_printlist()
}
