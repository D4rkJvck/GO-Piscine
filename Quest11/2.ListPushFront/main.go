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
func main() {

	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
