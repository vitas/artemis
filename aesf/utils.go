package aesf

import (
	"errors"
	"fmt"
)

func ConvertTypeToString(any interface{}) (CTypeName, error) {
	if any == nil {
		return CTYPE_NAME_UNKNOWN, errors.New("cannot get typename, unknown set")
	}
	stname := (CTypeName)(fmt.Sprintf("%T", any))
	if len(stname) <= 1 {
		return CTYPE_NAME_UNKNOWN, errors.New("cannot get typename, unknown set")
	}
	if stname[0:1] == "*" {
		stname = stname[1:len(stname)]
	}
	return stname, nil
}

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
