package leetcode

import "math"

type MinStackOptimized struct {
	min   int
	stack []int
}

func ConstructorOptimized() MinStackOptimized {
	return MinStackOptimized{
		min:   math.MaxInt64,
		stack: []int{},
	}
}

func (ms *MinStackOptimized) Push(x int) {
	if len(ms.stack) == 0 {
		ms.stack = append(ms.stack, 0)
		ms.min = x
	} else {
		ms.stack = append(ms.stack, x-ms.min)

		if x < ms.min {
			ms.min = x
		}
	}
}

func (ms *MinStackOptimized) Pop() {
	if len(ms.stack) == 0 {
		return
	}

	pop := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]

	if pop < 0 {
		ms.min = ms.min - pop
	}
}

func (ms *MinStackOptimized) Top() int {
	top := ms.stack[len(ms.stack)-1]

	if top > 0 {
		return top + ms.min
	}

	return ms.min
}

func (ms *MinStackOptimized) GetMin() int {
	return ms.min
}
