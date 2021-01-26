package main

import "fmt"

// Heap 堆类型
type Heap []int

func main() {
	var h = Heap{6, 20, 8, 27, 17, 11, 35, 32, 41}
	fmt.Println(h)
	h.Init()
	fmt.Printf("初始化后\n")
	fmt.Println(h) 

	h.Push(15)
	fmt.Println(h) 

	x, ok := h.Remove(11)
	fmt.Println(x, ok, h) 

	y, ok := h.Remove(8)
	fmt.Println(y, ok, h)

	z := h.Pop()
	fmt.Println(z, h)
}

// Init 初始化
func (h Heap) Init() {
	n := len(h)
	// 求出堆叶子结点的父结点，进行下沉操作
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

// swap 结点交换
func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// less 结点比较
func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

// up 上浮操作
func (h Heap) up(i int) {
	for {
		// f 为父结点
		f := (i - 1) / 2
		if i == f || h.less(f, i) {
			break
		}

		h.swap(f, i)
		i = f
	}
}

// Push 插入操作
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

// down 下沉操作
func (h Heap) down(i int) {
	for {
		l := 2*i + 1
		if l >= len(h) {
			// i 已经是叶子结点了
			break
		}

		j := l

		// 如果右子结点小于左子结点
		if r := l + 1; r < len(h) && h.less(r, l) {
			// 将右子结点与左子结点中最小值赋给j
			j = r
		}

		// 比较i处结点与i的的最小子树
		if h.less(i, j) {
			break
		}

		// 交换父结点和子结点
		h.swap(i, j)

		// 继续向下比较
		i = j
	}
}

// Remove 删除操作
func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}

	l := len(*h) - 1
	// 用最后的元素替换被删除元素
	h.swap(i, l)

	// 删除最后的元素
	x := (*h)[l]
	*h = (*h)[0:l]

	// 如果当前元素大于父结点
	if (*h)[i] > (*h)[(i-1)/2] {
		// 下沉
		h.down(i)
	} else {
		// 上浮
		h.up(i)
	}

	return x, true
}

// Pop 弹出堆顶
func (h *Heap) Pop() int {
	n := len(*h) - 1
	h.swap(0, n)
	x := (*h)[n]
	*h = (*h)[0:n]
	h.down(0)
	return x
}
