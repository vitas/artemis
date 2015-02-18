package aesf

import (
	"fmt"
)

type Bag interface {
	Get(idx int) Entity
	Size() int
	IsEmpty() bool
	Contains(e Entity) bool
	Add(e ...Entity)
	Set(idx int, e Entity)
	Remove(idx int) Entity
	RemoveEntity(e Entity) bool
	Grow()
	GrowSize(sz int)
}

type EntityBag struct {
	data []Entity
	size int
}

func NewBag() *EntityBag {
	return &EntityBag{make([]Entity, 64), 0}
}

func NewBagCap(c int) *EntityBag {
	return &EntityBag{make([]Entity, c), 0}
}

func (eb *EntityBag) Get(idx int) Entity {
	return eb.data[idx]
}

func (eb *EntityBag) Size() int {
	return eb.size
}

func (eb *EntityBag) GetCapacity() int {
	return len(eb.data)
}

func (eb *EntityBag) IsEmpty() bool {
	return eb.size == 0
}

func (eb *EntityBag) Clear() {
	for i := 0; i < eb.size-1; i++ {
		eb.data[i] = nil
	}
}

func (eb *EntityBag) Add(entities ...Entity) {
	for _, e := range entities {
		if eb.size == len(eb.data) {
			eb.Grow()
		}
		eb.data[eb.size] = e
		eb.size++
	}
}

func (eb *EntityBag) Contains(e Entity) bool {
	for _, ie := range eb.data {
		if ie == e {
			return true
		}
	}
	return false
}

func (eb *EntityBag) Remove(idx int) Entity {
	ce := eb.data[idx] // make copy of element to remove so it can be returned
	eb.size--
	eb.data[idx] = eb.data[eb.size] // overwrite item to remove with last element
	eb.data[eb.size] = nil          // null last element, so gc can do its work
	return ce
}

func (eb *EntityBag) RemoveEntity(e Entity) bool {
	for idx, ie := range eb.data {
		if ie == e {
			eb.Remove(idx)
			return true
		}
	}
	return false
}

func (eb *EntityBag) Set(idx uint, e Entity) {
	var lidx int
	lidx = (int)(idx)
	if lidx >= len(eb.data) {
		eb.GrowSize(lidx * 2)
	}
	eb.size = lidx + 1
	eb.data[lidx] = e
}

func (eb *EntityBag) Grow() {
	newCapacity := (len(eb.data)*3)/2 + 1
	eb.GrowSize(newCapacity)
}

func (eb *EntityBag) GrowSize(gsize int) {
	ndata := make([]Entity, gsize)
	eb.data = append(eb.data, ndata...)
	// replace with copy? need test performance first
}

func (eb *EntityBag) RemoveLast() Entity {
	if eb.IsEmpty() {
		return nil
	}
	return eb.Remove(eb.size - 1)
}

type IntList struct {
	uintlist []uint
}

func NewIntList() *IntList {
	return &IntList{[]uint{}}
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

func (il *IntList) Add(i ...uint) {
	il.uintlist = append(il.uintlist, i...)
}

func (il *IntList) String() string {
	return fmt.Sprintf("%#v", il.uintlist)
}
