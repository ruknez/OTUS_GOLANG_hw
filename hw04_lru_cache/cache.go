package hw04lrucache

import (
	"log"
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Len() int
	Remove(key Key) bool
	Clear()
}

type lruCache struct {
	mtx      sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lru *lruCache) Set(key Key, value interface{}) bool {
	lru.mtx.Lock()
	defer lru.mtx.Unlock()
	if val, exists := lru.items[key]; exists {
		lru.queue.MoveToFront(val)
		val.Value.(*cacheItem).value = value
		return true
	} else if lru.queue.Len() >= lru.capacity {
		backItem := lru.queue.Back()
		delete(lru.items, backItem.Value.(*cacheItem).key)
		lru.queue.Remove(backItem)
	}
	lru.items[key] = lru.queue.PushFront(&cacheItem{key, value})
	return false
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	lru.mtx.Lock()
	defer lru.mtx.Unlock()
	if val, exists := lru.items[key]; exists {
		lru.queue.MoveToFront(val)
		return val.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (lru *lruCache) Clear() {
	lru.mtx.Lock()
	defer lru.mtx.Unlock()
	lru.queue = NewList()
	lru.items = make(map[Key]*ListItem, lru.capacity)
}

func (lru *lruCache) Len() int {
	lru.mtx.Lock()
	defer lru.mtx.Unlock()
	if len(lru.items) != lru.queue.Len() {
		log.Panic("len(lru.items) ", len(lru.items), " != lru.queue.Len() ", lru.queue.Len())
	}
	return len(lru.items)
}

func (lru *lruCache) Remove(key Key) bool {
	lru.mtx.Lock()
	defer lru.mtx.Unlock()
	if val, ok := lru.items[key]; ok {
		lru.queue.Remove(val)
		delete(lru.items, key)
		return true
	}
	return false
}
