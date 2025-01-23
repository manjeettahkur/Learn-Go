package linkedlist

import "fmt"

type CircularNode struct {
	Value int
	Next  *CircularNode
}

type CircularLinkedList struct {
	Head *CircularNode
}

func (cll *CircularLinkedList) InsertAtBeginning(val int) {
	n := &CircularNode{Value: val, Next: nil}

	if cll.Head == nil {
		cll.Head = n
		cll.Head.Next = n
		return
	}

	current := cll.Head

	for current.Next != cll.Head {
		current = current.Next
	}

	n.Next = cll.Head
	cll.Head = n
	current.Next = n



	fmt.Println("current:", current)

}

func (cll *CircularLinkedList) InsertAtEnd(val int){
	n := &CircularNode{Value: val}
	if cll.Head == nil {
		fmt.Println("List is empty")
	}

	current := cll.Head


	for current.Next != cll.Head {
		current = current.Next
	}

	n.Next = current.Next
	current.Next = n
}

func (cll *CircularLinkedList) DisplayForward() {

	if cll.Head == nil {
		fmt.Println("List is empty")
	}

	current := cll.Head

	fmt.Printf("%d  ->  ", current.Value)

	current = current.Next

	for current != cll.Head {
		fmt.Printf("%d  -> ", current.Value)
		current = current.Next
	}

	fmt.Println("Nil")
}
