package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

func (l linkedList) Len() int {
	return l.length
}

func (l linkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l *linkedList) PushBack(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func (l *linkedList) Delete(key int) {

	if l.head.data == key {
		l.head = l.head.next
		l.length--
		return
	}
	var prev *node = nil
	curr := l.head
	for curr != nil && curr.data != key {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		fmt.Println("Key Not found")
		return
	}
	prev.next = curr.next
	l.length--
	fmt.Println("Node Deleted")
}

func (l_copy *linkedList) cloneList(l_orig linkedList) {
	for l_orig.head != nil {
		l_copy.PushBack(&node{data: l_orig.head.data})
		l_orig.head = l_orig.head.next
	}
}

func (l *linkedList) getNElem(idx int) int {
	if idx > l.length {
		return -1 //todo return error
	}
	count := 0
	current := l.head
	for current != nil {
		if idx == count {
			return current.data
		}
		current = current.next
		count++
	}

	return -1 //todo return error
}

func (l *linkedList) pushAtNpos(idx int, val int) int {
	if idx > l.length {
		return -1 //todo return error
	}
	count := 0
	current := l.head
	for current != nil {
		if idx == count {
			current.data = val
			return 1
		}
		current = current.next
		count++
	}

	return -1 //todo return error
}

func (l *linkedList) ListSum(l1 linkedList, l2 linkedList) {
	// max := (l1.length >= l2.length) ? l1.length : l2.length
	deltal1l2, addNext, addNow := 0, 0, 0
	if l2.length > l1.length {
		l.cloneList(l2)
		deltal1l2 = l2.length - l1.length
	} else {
		l.cloneList(l1)
		deltal1l2 = l1.length - l2.length
	}
	l.Display()
	fmt.Println("delta", deltal1l2)
	for l.length-deltal1l2 > 0 {
		addNow = addNext
		addNext = 0
		sum := l1.getNElem(l.length-deltal1l2-1) + l2.getNElem(l.length-deltal1l2-1)
		fmt.Println("sum", sum)
		if sum > 9 {
			sum = sum % 10
			addNext = 1
		}
		fmt.Println("sum", sum)
		if addNow == 1 {
			sum++
		}
		l.pushAtNpos(l.length-deltal1l2-1, sum)

		deltal1l2++
	}
}
