package main

import "math"

type MinStack1 struct {
	minValue int
	value    int
	pre      *MinStack1
	next     *MinStack1
}

func MinStackConstructor1() MinStack1 {
	return MinStack1{minValue: math.MaxInt, value: math.MaxInt, pre: nil, next: nil}
}

func (this *MinStack1) Push1(val int) {
	oldnext := this.next
	newNext := &MinStack1{}
	newNext.value = val
	newNext.next = oldnext
	newNext.pre = this
	if oldnext == nil {
		newNext.minValue = val
	} else {
		newNext.minValue = min(val, oldnext.minValue)
		oldnext.pre = newNext
	}
	this.next = newNext
}

func (this *MinStack1) Pop1() {
	oldnext := this.next
	if oldnext.next == nil {
		this.next = nil
	} else {
		this.next = oldnext.next
		oldnext.next.pre = this
	}
}

func (this *MinStack1) Top1() int {
	return this.next.value
}

func (this *MinStack1) GetMin1() int {
	return this.next.minValue
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
