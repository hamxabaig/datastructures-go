package linked

import "fmt"

type node struct {
	next  *node
	value interface{}
}

// List exposes the fns on a linked list e.g Append, Prepend, ForEach, Print...
type List struct {
	head *node
}

// New creates new list exposing methods such as Find, ForEach, Append...
func New() List {
	return List{
		head: nil,
	}
}

// Append the value to the end of the linked list
// if there are no values, that val becomes the first item
func (l *List) Append(val interface{}) {
	fmt.Println("Appending: ", val, " to list!")
	var node = &node{
		next:  nil,
		value: val,
	}

	if l.head == nil {
		l.head = node
	} else {
		var tail = getTail(l.head)
		tail.next = node
	}
}

// Prepend the val to the list by changing the head to the new node
// and assigning head to the node's next
func (l *List) Prepend(val interface{}) {
	fmt.Println("Prepending: ", val, " to list!")
	var node = &node{
		next:  nil,
		value: val,
	}

	if l.head == nil {
		l.head = node
	} else {
		// Change head to the current node
		node.next = l.head
		l.head = node
	}
}

// Print all the values in a linked list
func (l List) Print() {
	l.ForEach(func(val interface{}) {
		fmt.Println(val)
	})
}

// Length returns the count of items in the linked list
func (l List) Length() int {
	length := 0
	l.ForEach(func(val interface{}) {
		length++
	})
	return length
}

// ForEach loops over all the linked list items and calls the
// fn passed to it with current val
func (l List) ForEach(fn func(val interface{})) {
	var head = l.head

	if head != nil {
		for head != nil {
			fn(head.value)
			head = head.next
		}
	}
}

func getTail(node *node) *node {
	var head = node
	if head != nil {
		for head.next != nil {
			head = head.next
		}
	}

	return head
}

// TODO, implement Exists, Get, InsertAt, Clear fns
