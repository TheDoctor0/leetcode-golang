package solutions

import (
    "container/heap"
)

type Pair struct {
    FirstIndex  int
    SecondIndex int
    Value       int
}

type PairHeap []Pair

func (pairHeap PairHeap) Len() int {
    return len(pairHeap)
}

func (pairHeap PairHeap) Less(i, j int) bool {
    return pairHeap[i].Value < pairHeap[j].Value
}

func (pairHeap PairHeap) Swap(i, j int) {
    pairHeap[i], pairHeap[j] = pairHeap[j], pairHeap[i]
}

func (pairHeap *PairHeap) Push(x interface{}) {
    *pairHeap = append(*pairHeap, x.(Pair))
}

func (pairHeap *PairHeap) Pop() interface{} {
    root := (*pairHeap)[len(*pairHeap) - 1]
    *pairHeap = (*pairHeap)[:len(*pairHeap) - 1]

    return root
}


func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return nil
    }

    pairHeap := new(PairHeap)
    heap.Init(pairHeap)

    for i := 0; i < len(nums2); i++ {
        heap.Push(pairHeap, Pair{0, i, nums1[0] + nums2[i]})
    }

    result := make([][]int, 0)

    if k > len(nums1) * len(nums2) {
        k = len(nums1) * len(nums2)
    }

    for i := 0; i < k; i++ {
        pair := heap.Pop(pairHeap).(Pair)
        result = append(result, []int{nums1[pair.FirstIndex], nums2[pair.SecondIndex]})

        if pair.FirstIndex < len(nums1) - 1 {
            heap.Push(pairHeap, Pair{pair.FirstIndex + 1, pair.SecondIndex, nums1[pair.FirstIndex + 1] + nums2[pair.SecondIndex]})
        }
    }

    return result
}