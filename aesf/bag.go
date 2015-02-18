package aesf

import (
	"fmt"
)

type BaseEntityBag struct {
	data []*BaseEntity
	size int
}

func New() *BaseEntityBag {
	return &BaseEntityBag{make([]*BaseEntity, 64), 0}
}

func NewCap(c int) *BaseEntityBag {
	return &BaseEntityBag{make([]*BaseEntity, c), 0}
}

func (eb *BaseEntityBag) Get(idx int) *BaseEntity {
	return eb.data[idx]
}

func (eb *BaseEntityBag) Size() int {
	return eb.size
}

func (eb *BaseEntityBag) GetCapacity() int {
	return len(eb.data)
}

func (eb *BaseEntityBag) IsEmpty() bool {
	return eb.size == 0
}

func (eb *BaseEntityBag) Clear() {
	for i := 0; i < eb.size-1; i++ {
		eb.data[i] = nil
	}
}

func (eb *BaseEntityBag) Add(entities ...*BaseEntity) {
	for _, e := range entities {
		if eb.size == len(eb.data) {
			eb.Grow()
		}
		eb.data[eb.size] = e
		eb.size++
	}
}

func (eb *BaseEntityBag) Contains(e *BaseEntity) bool {
	for _, ie := range eb.data {
		if ie == e {
			return true
		}
	}
	return false
}

func (eb *BaseEntityBag) Remove(idx int) *BaseEntity {
	ce := eb.data[idx] // make copy of element to remove so it can be returned
	eb.size--
	eb.data[idx] = eb.data[eb.size] // overwrite item to remove with last element
	eb.data[eb.size] = nil          // null last element, so gc can do its work
	return ce
}

func (eb *BaseEntityBag) RemoveEntity(e *BaseEntity) bool {
	for idx, ie := range eb.data {
		if ie == e {
			eb.Remove(idx)
			return true
		}
	}
	return false
}

func (eb *BaseEntityBag) Set(idx uint, e *BaseEntity) {
	var lidx int
	lidx = (int)(idx)
	if lidx >= len(eb.data) {
		eb.GrowSize(lidx * 2)
	}
	eb.size = lidx + 1
	eb.data[idx] = e
}

func (eb *BaseEntityBag) Grow() {
	newCapacity := (len(eb.data)*3)/2 + 1
	eb.GrowSize(newCapacity)
}

func (eb *BaseEntityBag) GrowSize(gsize int) {
	ndata := make([]*BaseEntity, gsize)
	eb.data = append(eb.data, ndata...)
	// replace with copy? need test performance first
}

func (eb *BaseEntityBag) RemoveLast() *BaseEntity {
	if eb.IsEmpty() {
		return nil
	}
	return eb.Remove(eb.size - 1)
}

type UIntList struct {
	uintlist []uint
}

func NewUIntList() *UIntList {
	return &UIntList{[]uint{}}
}

func (il *UIntList) Size() int {
	return len(il.uintlist)
}

func (il *UIntList) Pop() {
	if len(il.uintlist) == 0 {
		return
	}
	il.uintlist = il.uintlist[:len(il.uintlist)-1]
}

func (il *UIntList) Add(i ...uint) {
	il.uintlist = append(il.uintlist, i...)
}

func (il *UIntList) String() string {
	return fmt.Sprintf("%#v", il.uintlist)
}
