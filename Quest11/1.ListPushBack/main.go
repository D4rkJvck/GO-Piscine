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

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

func main() {

	link := &List{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
