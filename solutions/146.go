package solutions

type DLinkedList struct {
    key  int
    val  int
    prev *DLinkedList
    next *DLinkedList
}

type LRUCache struct {
    capacity int
    head     *DLinkedList
    tail     *DLinkedList
    cache    map[int]*DLinkedList
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*DLinkedList),
    }
}

func (lru *LRUCache) Get(key int) int {
    element, ok := lru.cache[key]

    if !ok {
        return -1
    }

    lru.removeFromChain(element)
    lru.addToChain(element)

    return element.val
}

func (lru *LRUCache) Put(key int, value int) {
    list, ok := lru.cache[key]

    if !ok {
        list = &DLinkedList{key: key, val: value}
        lru.cache[key] = list
    } else {
        list.val = value

        lru.removeFromChain(list)
    }

    lru.addToChain(list)

    if len(lru.cache) > lru.capacity {
        last := lru.tail

        lru.removeFromChain(last)

        delete(lru.cache, last.key)
    }
}

func (lru *LRUCache) addToChain(list *DLinkedList) {
    list.next = nil

    if lru.head != nil {
        lru.head.next = list
        list.prev = lru.head
    }

    lru.head = list

    if lru.tail == nil {
        lru.tail = list
    }
}

func (lru *LRUCache) removeFromChain(list *DLinkedList) {
    if list == lru.head {
        lru.head = list.prev
    }

    if list == lru.tail {
        lru.tail = list.next
    }

    if list.next != nil {
        list.next.prev = list.prev
    }

    if list.prev != nil {
        list.prev.next = list.next
    }
}
