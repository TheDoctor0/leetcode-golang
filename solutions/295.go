package solutions

import "container/heap"

type MinHeap []int

func (heap MinHeap) Len() int {
    return len(heap)
}

func (heap MinHeap) Less(i, j int) bool {
    return heap[i] < heap[j]
}

func (heap MinHeap) Swap(i, j int) {
    heap[i], heap[j] = heap[j], heap[i]
}

func (heap *MinHeap) Push(x interface{}) {
    *heap = append(*heap, x.(int))
}

func (heap *MinHeap) Pop() interface{} {
    min := (*heap)[len(*heap) - 1]
    *heap = (*heap)[:len(*heap) - 1]

    return min
}

type MaxHeap []int

func (heap MaxHeap) Len() int {
    return len(heap)
}

func (heap MaxHeap) Less(i, j int) bool {
    return heap[i] > heap[j]
}

func (heap MaxHeap) Swap(i, j int) {
    heap[i], heap[j] = heap[j], heap[i]
}

func (heap *MaxHeap) Push(x interface{}) {
    *heap = append(*heap, x.(int))
}

func (heap *MaxHeap) Pop() interface{} {
    min := (*heap)[len(*heap) - 1]
    *heap = (*heap)[:len(*heap) - 1]

    return min
}

type MedianFinder struct {
    minHeap MinHeap
    maxHeap MaxHeap
}

func Constructor() MedianFinder {
    return MedianFinder{
        minHeap: MinHeap{},
        maxHeap: MaxHeap{},
    }
}

func (this *MedianFinder) AddNum(num int) {
    if this.maxHeap.Len() == 0 || num <= this.maxHeap[0] {
        heap.Push(&this.maxHeap, num)
    } else {
        heap.Push(&this.minHeap, num)
    }

    if this.maxHeap.Len() > this.minHeap.Len() + 1 {
        maxLeft := heap.Pop(&this.maxHeap)
        heap.Push(&this.minHeap, maxLeft)
    }

    if this.minHeap.Len() > this.maxHeap.Len() {
        minRight := heap.Pop(&this.minHeap)
        heap.Push(&this.maxHeap, minRight)
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.minHeap.Len() == this.maxHeap.Len() {
        return float64(this.minHeap[0] + this.maxHeap[0]) / 2
    }

    return float64(this.maxHeap[0])
}
