type MinPair struct {
	val int
	min int
}

type MinStack struct {
	data []MinPair
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.data) != 0 {
		if min > this.data[len(this.data) - 1].min {
			min = this.data[len(this.data) - 1].min
		}
	}

	this.data = append(this.data, MinPair{
		val: val,
		min: min,
	})
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}

	this.data = this.data[:len(this.data) - 1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data) - 1].val
}

func (this *MinStack) GetMin() int {
	if len(this.data) == 0 {
		return 0
	}


	return this.data[len(this.data) - 1].min
}
