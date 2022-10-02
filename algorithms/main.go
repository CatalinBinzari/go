package main

import "fmt"

func main() {
	list := linkedList{}
	list1 := linkedList{}
	list_sum := linkedList{}
	// clone_list1 := linkedList{}
	Len := list.Len()
	fmt.Println("len", Len)
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 8}
	node4 := &node{data: 4}
	node5 := &node{data: 4}
	list.PushBack(node1)
	list.PushBack(node2)
	list.PushBack(node3)
	list.PushBack(node4)
	list.PushBack(node5)
	list1.PushBack(node1)
	list1.PushBack(node2)
	list1.PushBack(node3)
	// list1.PushBack(node4)
	// list1.PushBack(node5)
	list.Display()
	// clone_list1.cloneList(list)
	// list.Delete(40)
	// list.Display()
	// clone_list1.Display()
	// getNElem := list.getNElem(4)
	// getNElem := list.getNElem(2)
	// fmt.Println("getNElem", getNElem)
	// _ = list.pushAtNpos(2, 60)
	// getNElem = list.getNElem(2)
	// fmt.Println("getNElem", getNElem)
	list1.Display()
	list_sum.ListSum(list, list1)
	list_sum.Display()
}
