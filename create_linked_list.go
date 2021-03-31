package main

import "fmt"

// node creation
type node struct {
	data int
	next *node
}

// linkedList description
type linkedList struct {
	head   *node
	length int
}

// prepend / creating linkedList structure with method receiver
// method receiver means, we can call this func only on linkedList not at package level
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	// when list is empty
	if l.length == 0 {
		return
	}

	// when we delete the head value (node 1)
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		// when given value not present in the list, then last will be nil
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 38}
	node3 := &node{data: 28}
	node4 := &node{data: 18}
	node5 := &node{data: 8}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.printListData()
	myList.deleteWithValue(28)
	myList.deleteWithValue(8)
	// trying to delete the value which is not present
	myList.deleteWithValue(0)
	myList.printListData()

	// trying to delete value in empty list
	emptyList := linkedList{}
	emptyList.deleteWithValue(12)
}
