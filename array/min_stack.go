package main

import (
	"container/list"
	"fmt"
)

type VarWithPos struct {
	v   int
	pos *list.Element
}

type MinStack struct {
	var_pos     []VarWithPos
	sorted_nums *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var ms *MinStack = new(MinStack)

	ms.var_pos = []VarWithPos{}
	ms.sorted_nums = list.New()

	return *ms
}

func (this *MinStack) Push(x int) {
	var mark *list.Element
	for e := this.sorted_nums.Front(); e != nil; e = e.Next() {
		if x <= e.Value.(int) {
			mark = e
			break
		}
	}

	var pos *list.Element
	if mark == nil {
		pos = this.sorted_nums.PushBack(x)
	} else {
		pos = this.sorted_nums.InsertBefore(x, mark)
	}

	var_with_pos := new(VarWithPos)
	var_with_pos.v = x
	var_with_pos.pos = pos

	this.var_pos = append(this.var_pos, *var_with_pos)
}

func (this *MinStack) Pop() {
	l := len(this.var_pos)
	if l == 0 {
		return
	}

	top_pos := this.var_pos[l-1].pos

	this.var_pos = this.var_pos[0 : l-1]
	this.sorted_nums.Remove(top_pos)
}

func (this *MinStack) Top() int {
	l := len(this.var_pos)
	if l == 0 {
		return 0
	}

	return this.var_pos[l-1].v
}

func (this *MinStack) GetMin() int {
	if len(this.var_pos) == 0 {
		return 0
	}

	return this.sorted_nums.Front().Value.(int)
}

func main() {
	obj := Constructor()
	obj.Push(10)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
}
