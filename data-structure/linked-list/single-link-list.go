package linkedlist

import "fmt"



type SingleNode struct {
	Value int
	Next *SingleNode
}


type LinkList struct {
	Head *SingleNode
}


func (l *LinkList) InsertAtBegninning(data int)  {
	newNode := &SingleNode{Value: data, Next: l.Head}
	l.Head = newNode	
}

func (l *LinkList) InsertAtEnd(data int) {
	newNode := &SingleNode{Value: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode

}


func(l *LinkList) InsertAfter(preNode *SingleNode, data int) {
	if preNode == nil {
		fmt.Println("previous Node can't be nil")
		return
	}
	newNode := &SingleNode{data,  preNode.Next}

	preNode.Next = newNode
}

func (l *LinkList) Delete(data int){

	if l.Head == nil {
		fmt.Println("Empty link list")
		return
	}

	if l.Head.Value == data {
		l.Head = l.Head.Next
	}
	current := l.Head

	for current.Next != nil {
		if current.Value == data {
			current.Next = current.Next.Next
		}
	}

}


func (l *LinkList) Search(data int)  *SingleNode{
	current := l.Head

	for current != nil {
		if current.Value ==  data {
			return current
		}
	}

	return nil
}

func (l *LinkList) Walk() {
	current := l.Head

	for current != nil {
		fmt.Println("Node Data", current.Value)
		current = current.Next
	}

	fmt.Println("Nil")
}

func (l *LinkList) Reverse(){
	var pre *SingleNode
	current := l.Head
	var next *SingleNode

	// 1 -> 2 -> 3 -> 4 -> nil 
	// next 2 -> 3 -> 4 -> nil
	// current.next = nil
	// pre = 1 -> nil
	// current  2 -> 3 -> 4 -> nil 
	// next = 3 -> 4 -> nil
	// 2 -> 1 -> nil 
	

	for current != nil {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}

	l.Head = pre
}
