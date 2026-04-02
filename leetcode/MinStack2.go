package main

type MinStack2 struct {
	value    []int
	minvalue []int
}

func MinStackConstructor() MinStack2 {
	return MinStack2{
		value:    make([]int, 0),
		minvalue: make([]int, 0),
	}
}

func (this *MinStack2) Push(val int) {
	this.value = append(this.value, val)
	if len(this.minvalue) > 0 {
		this.minvalue = append(this.minvalue, min(this.minvalue[len(this.minvalue)-1], val))
	} else {
		this.minvalue = append(this.minvalue, val)
	}
}

func (this *MinStack2) Pop() {
	this.value = this.value[:len(this.value)-1]
	this.minvalue = this.minvalue[:len(this.minvalue)-1]
}

func (this *MinStack2) Top() int {
	return this.value[len(this.value)-1]
}

func (this *MinStack2) GetMin() int {
	return this.minvalue[len(this.minvalue)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
