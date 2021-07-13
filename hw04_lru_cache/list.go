package hw04lrucache

import (
	"fmt"
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	PrintList(s int)
}

type ListItem struct {
	Value interface{}
	next  *ListItem // зачем их делать public ???
	prev  *ListItem // зачем их делать public ???
}

type list struct {
	header *ListItem
	tail   *ListItem
	len    uint64
}

func NewList() List {
	return new(list)
}

func (ls *list) Len() int {
	return int(ls.len)
}

func (ls *list) Front() *ListItem {
	return ls.header
}

func (ls *list) Back() *ListItem {
	return ls.tail
}

func (ls *list) PushFront(v interface{}) *ListItem {
	newListItem := new(ListItem)
	newListItem.Value = v
	newListItem.prev = nil
	newListItem.next = ls.header

	if ls.header != nil {
		ls.header.prev = newListItem
	}
	ls.header = newListItem

	if ls.tail == nil {
		ls.tail = newListItem
	}
	ls.len++
	return newListItem
}

func (ls *list) PushBack(v interface{}) *ListItem {
	newListItem := new(ListItem)
	newListItem.Value = v
	newListItem.next = nil

	newListItem.prev = ls.tail
	if ls.tail != nil {
		ls.tail.next = newListItem
	}
	ls.tail = newListItem
	if ls.header == nil {
		ls.header = newListItem
	}
	ls.len++
	return newListItem
}

func (ls *list) Remove(current *ListItem) {
	switch current {
	case nil:
		return
	case ls.header:
		ls.header = current.next
		current = nil
	case ls.tail:
		ls.tail = current.prev
		current = nil
	default:
		current.next.prev = current.prev
		current.prev.next = current.next
		current = nil
	}
	ls.len--
}

func (ls *list) MoveToFront(current *ListItem) {
	if current == nil || ls.header == current || ls.len == 1 {
		return
	}
	ls.header.prev = current

	if ls.tail == current {
		current.next = ls.header
		ls.tail = current.prev
		current.prev.next = nil
		current.prev = nil
		ls.header = current
	} else {
		current.prev.next = current.next
		current.next.prev = current.prev
		current.next = ls.header
		current.prev = nil
		ls.header = current
	}
}

func (ls *list) PrintList(s int) {
	fmt.Println("Print List ", s)
	currentItem := ls.header
	for i := 0; i < int(ls.len); i++ {
		fmt.Printf("i = %d %v ", i, currentItem.Value)
		currentItem = currentItem.next
	}
	fmt.Println("\n ")
}
