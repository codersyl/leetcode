func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 返回有序数组 arr 中，第一个大于等于 target 的元素的索引
// 若所有元素均小于 target，返回 n = len(arr)
func lower_bound(arr []int, target int) int {
	n := len(arr)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}




// BigHeap
import (
	"container/heap"
)

type BigHeap []int
func (h BigHeap) Len() int { return len(h) }
// 名为 Less，实为 Greater
func (h BigHeap) Less(i, j int) bool { return h[i] > h[j] } // 大顶堆，返回值决定是否交换元素
func (h BigHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *BigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *BigHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *BigHeap) Top() int {
	arr := *h
	return arr[0]
}




// SmallHeap
import "container/heap"
type SmallHeap []int
func (h SmallHeap) Len() int { return len(h) }
func (h SmallHeap) Less(i, j int) bool { return h[i] < h[j] } // 小顶堆，返回值决定是否交换元素
func (h SmallHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *SmallHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *SmallHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *SmallHeap) Top() int {
	arr := *h
	return arr[0]
}