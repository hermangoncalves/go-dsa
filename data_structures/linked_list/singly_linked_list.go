package main

import "fmt"

// Node represents a single node in a linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a singly linked list
// Characteristics:
// - Dynamic size
// - Efficient insertion/deletion
type LinkedList struct {
	head *Node
	tail *Node
}

// Add adds a new node with the given value to the end of the list.
// If the list is empty, the new node becomes the head (and tail).
func (l *LinkedList) Add(value int) {
	newNode := &Node{value: value} // Create a new node
	if l.head == nil {            // If the list is empty, initialize head and tail
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode // Link the new node to the end
	l.tail = newNode      // Update the tail
}

// Remove deletes the first occurrence of the given value in the list.
// Returns true if the value was removed, false otherwise.
func (l *LinkedList) Remove(value int) bool {
	if l.head == nil { // Empty list case
		fmt.Println("List is empty, nothing to remove.")
		return false
	}

	if l.head.value == value { // Removing the head
		l.head = l.head.next
		if l.head == nil { // If the list becomes empty, reset tail
			l.tail = nil
		}
		return true
	}

	current := l.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next != nil { // Value found
		current.next = current.next.next
		if current.next == nil { // If the removed node was the tail, update tail
			l.tail = current
		}
		return true
	}

	fmt.Printf("Value %d not found in the list.\n", value)
	return false
}

// Print displays all elements in the list.
// If the list is empty, prints an appropriate message.
func (l *LinkedList) Print() {
	if l.head == nil {
		fmt.Println("The list is empty.")
		return
	}

	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

// Example Test Cases
func main() {
	fmt.Println("### Example Test Cases ###")
	list := LinkedList{}

	fmt.Println("Add elements: 10, 20, 30")
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Print()

	fmt.Println("Remove element: 20")
	list.Remove(20)
	list.Print()

	fmt.Println("Remove element: 10 (head)")
	list.Remove(10)
	list.Print()

	fmt.Println("Remove element: 30 (tail)")
	list.Remove(30)
	list.Print()

	fmt.Println("Try removing from an empty list")
	list.Remove(40)
	list.Print()
}
