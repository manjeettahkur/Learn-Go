package linkedlist

import "fmt"

type doubleNode struct {
	Value int
	Prev  *doubleNode
	Next  *doubleNode
}

type DoublyLinkList struct {
	Head *doubleNode
	Tail *doubleNode
}

// 1. Insertion at beginning
func (dll *DoublyLinkList) InsertAtBeginning(data int){
	newNode := &doubleNode{Value: data, Prev: nil, Next: dll.Head }
 
	if dll.Head == nil{ // List is empty
	  dll.Head = newNode
	  dll.Tail = newNode
	  return
	}
 
	dll.Head.Prev = newNode
	dll.Head = newNode
 }
 
 // 2. Insertion at end
 func (dll *DoublyLinkList) InsertAtEnd(data int) {
   newNode := &doubleNode{Value: data, Prev: dll.Tail, Next: nil}
 
   if dll.Tail == nil{ //list is empty
	 dll.Head = newNode
	 dll.Tail = newNode
	 return
   }
 
   dll.Tail.Next = newNode
   dll.Tail = newNode
 
 }

func (dll *DoublyLinkList) DisplayForward() {
	fmt.Print("Forward: ")
	current := dll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}


func (dll *DoublyLinkList) InsertAfter(node *doubleNode, data int) {
	if node == nil {
		fmt.Println("Node can't be nil")
		return
	}

	n := &doubleNode{Value: data, Prev: node, Next: node.Next}

	if node.Next != nil {

		node.Next.Prev = n

	}

	node.Next = n

	if n.Next == nil {
		dll.Tail = n
	}
}