package aesf

import (
	"fmt"
)

type IntList struct {
	uintlist []int
}

func NewIntList() *IntList {
	return &IntList{[]int{}}
}

func (il *IntList) Size() int {
	return len(il.uintlist)
}

func (il *IntList) Pop() {
	if len(il.uintlist) == 0 {
		return
	}
	il.uintlist = il.uintlist[:len(il.uintlist)-1]
}

func (il *IntList) Add(i ...int) {
	il.uintlist = append(il.uintlist, i...)
}

func (il *IntList) String() string {
	return fmt.Sprintf("%#v", il.uintlist)
}
