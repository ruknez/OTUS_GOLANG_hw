package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	value interface{}
	next  *ListItem
	prev  *ListItem
}

func (ls *ListItem) Getvalue() interface{} {
	return ls.value
}

type list struct {
	header *ListItem
	tail   *ListItem
	len    int
}

func NewList() List {
	return new(list)
}

func (ls *list) Len() int {
	return ls.len
}

func (ls *list) Front() *ListItem {
	return ls.header
}

func (ls *list) Back() *ListItem {
	return ls.tail
}

func (ls *list) PushFront(v interface{}) *ListItem {
	newListItem := new(ListItem)
	newListItem.value = v
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
	newListItem.value = v
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
		if ls.header != nil && ls.header.prev != nil {
			ls.header.prev = nil
		}
	case ls.tail:
		ls.tail = current.prev
		if ls.tail != nil && ls.tail.next != nil {
			ls.header.prev = nil
		}
	default:
		current.next.prev = current.prev
		current.prev.next = current.next
	}
	ls.len--
}

func (ls *list) MoveToFront(current *ListItem) {
	if current == nil || ls.header == current || ls.len == 1 {
		return
	}
	ls.header.prev = current

	if ls.tail == current {
		ls.tail = current.prev
	} else {
		current.next.prev = current.prev
	}
	current.prev.next = current.next
	current.next = ls.header
	current.prev = nil
	ls.header = current
}
