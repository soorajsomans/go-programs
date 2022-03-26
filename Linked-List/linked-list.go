package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) InsertAtTail(val int) {
	fmt.Printf("Adding %v at tail\n", val)
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	headPtr := l.head
	for i := 0; i < l.len; i++ {
		if headPtr.next == nil {
			headPtr.next = &n
			l.len++
			return
		}
		headPtr = headPtr.next
	}
}

func (l *LinkedList) InsertAtHead(val int) {
	fmt.Printf("Adding %v at head\n", val)
	n := Node{}
	n.value = val

	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	headPtr := l.head

	n.next = headPtr
	l.head = &n
	l.len++

}

func (l *LinkedList) InsertAtPosition(pos int, val int) {
	fmt.Printf("Adding %v at position %v\n", val, pos)

	n := Node{}
	n.value = val

	if pos < 0 || pos > l.len {
		fmt.Println("Invalid position")
		return
	}
	if pos == 0 {
		l.head = &n
		l.len++
		return
	}

	headPtr := l.head

	for i := 0; i < pos-1; i++ {
		headPtr = headPtr.next
	}
	n.next = headPtr.next
	headPtr.next = &n
	l.len++

}

func (l *LinkedList) Search(val int) {
	if l.len == 0 {
		fmt.Println("Empty Linked List")
		return
	}
	headPtr := l.head
	for i := 0; i < l.len; i++ {
		if headPtr.value == val {
			fmt.Printf("Element %v found at position %v \n", headPtr.value, i)
			return
		}
		headPtr = headPtr.next
	}
	fmt.Printf("Element %v not found in Linked List \n", val)

}

func (l *LinkedList) DeleteFromHead() {

	if l.len == 0 {
		fmt.Println("Empty Linked List")
		return
	}
	fmt.Printf("Deleting element %v from head", l.head.value)
	l.head = l.head.next
	l.len--

}

func (l *LinkedList) DeleteFromTail() {
	if l.len == 0 {
		fmt.Println("Empty Linked List")
		return
	}
	headPtr := l.head
	for i := 0; i < l.len-1; i++ {
		headPtr = headPtr.next
	}
	fmt.Printf("Deleting element %v from tail", headPtr.value)
	headPtr.next = nil
	l.len--
}

func (l *LinkedList) DeleteFromPostion(pos int) error {
	if pos < 0 || pos > l.len {
		fmt.Println("Invalid Position")
		return errors.New("Invalid position")
	}
	if l.len == 0 {
		fmt.Println("Empty Linked List")
		return errors.New("Empty Linked List")
	}
	if pos == 0 {
		l.DeleteFromHead()
		return nil
	} else if pos == l.len {
		l.DeleteFromTail()
		return nil
	} else {
		headPtr := l.head
		for i := 0; i < pos-1; i++ {
			headPtr = headPtr.next
		}
		fmt.Printf("Deleting element %v from position %v \n ", headPtr.next.value, pos)
		headPtr.next = headPtr.next.next
		l.len--
		return nil
	}
}

func (l *LinkedList) Print() {
	fmt.Println("\n=========== Linked List ============\n")
	if l.len == 0 {
		fmt.Println("Linked List is empty")
	}
	headPtr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf(" %v ->", headPtr.value)
		headPtr = headPtr.next
	}
	fmt.Println("\n====================================\n")
}

func main() {
	element := LinkedList{}
	options := `Select an operation 

	1. Insert At Head
	2. Insert At Tail
	3. Insert At Position
	4. Search an element
	5. Delete From Head
	6. Delete From Tail
	7. Delete From Position
	8. Print Linked List 
	9. Exit
`
loop:
	for {
		fmt.Println("\n============Linked List=============\n")
		fmt.Print(options)
		fmt.Println("\n====================================\n")
		var selectedOption int
		fmt.Scan(&selectedOption)
		switch selectedOption {
		case 1:
			fmt.Println("Insert At Head")
			var val int
			fmt.Println("Enter value to be inserted")
			fmt.Scan(&val)
			element.InsertAtHead(val)
		case 2:
			fmt.Println("Insert At Tail")
			var val int
			fmt.Println("Enter value to be inserted")
			fmt.Scan(&val)
			element.InsertAtTail(val)
		case 3:
			fmt.Println("Insert At Position")
			var val, pos int
			fmt.Println("Enter value to be inserted")
			fmt.Scan(&val)
			fmt.Println("Enter position")
			fmt.Scan(&pos)
			element.InsertAtPosition(pos, val)
		case 4:
			fmt.Println("Search an element")
			var val int
			fmt.Println("Enter element to be searched")
			fmt.Scan(&val)
			element.Search(val)
		case 5:
			fmt.Println("Delete From Head")
			element.DeleteFromHead()
		case 6:
			fmt.Println("Delete From Tail")
			element.DeleteFromTail()
		case 7:
			fmt.Println("Delete From Position")
			var val int
			fmt.Println("Enter position")
			fmt.Scan(&val)
			element.DeleteFromPostion(val)
		case 8:
			element.Print()
		case 9:
			break loop
		default:
			fmt.Println("Invalid Option")
		}

	}
	// element.InsertAtTail(5)

	// element.Print()

	// element.InsertAtHead(1)
	// element.Print()

	// element.InsertAtPosition(2, 4)
	// element.Print()
	// element.InsertAtPosition(2, 3)
	// element.Print()
	// element.Search(13)

	//element.DeleteFromHead()
	//element.Print()
	//element.DeleteFromTail()
	//element.Print()
	// element.DeleteFromPostion(2)
	// element.Print()
}
