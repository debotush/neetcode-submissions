type MinStack struct {
	top     int
	minArr  []int
	stack   []int
}

func Constructor() MinStack {
	return MinStack {
		top: 0,
		minArr: make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.top == 0 {
		this.minArr = append(this.minArr, val) 
	} else {
		this.minArr = append(this.minArr, min(this.minArr[this.top - 1], val))
	}
	this.top += 1
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minArr = this.minArr[:len(this.minArr)-1]
	this.top -= 1
}

func (this *MinStack) Top() int {
	return this.stack[this.top - 1]
}

func (this *MinStack) GetMin() int {
	return this.minArr[this.top - 1]
}

func min (a, b int) int {
	if a > b {
		return b
	}

	return a
}
