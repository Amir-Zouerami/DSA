package leetcode

type MinStack struct {
	min   []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{min: make([]int, 0), stack: make([]int, 0)}
}

func (ms *MinStack) Push(val int) {
	if len(ms.min) == 0 || ms.min[len(ms.min)-1] > val {
		ms.min = append(ms.min, val)
	} else {
		ms.min = append(ms.min, ms.min[len(ms.min)-1])
	}

	ms.stack = append(ms.stack, val)
}

func (ms *MinStack) Pop() {
	ms.min = ms.min[:len(ms.min)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.min[len(ms.min)-1]
}
