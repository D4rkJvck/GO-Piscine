package main

import "fmt"

type List struct {
	Head *NodeL
	Tail *NodeL
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListPushFront(l *List, data interface{}) {

	node := &NodeL{Data: data}

	if l.Head == nil {
		l.Tail = node
		l.Head = l.Tail
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

func ListSize(l *List) int {
	var count int
	for l.Head != nil {
		l.Head = l.Head.Next
		count++
	}
	return count
}

func main() {

	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}
