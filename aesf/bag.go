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
	data []*Entity
	size int
}

func NewEntityBag(c int) *EntityBag {
	return &EntityBag{make([]*Entity, c), 0}
}

func (eb *EntityBag) String() string {
	return fmt.Sprintf("%#v", eb.data)
}

func (eb *EntityBag) Get(idx int) *Entity {
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

func (eb *EntityBag) Add(entities ...*Entity) {
	for _, e := range entities {
		if eb.size == len(eb.data) {
			eb.Grow()
		}
		eb.data[eb.size] = e
		eb.size++
	}
}

func (eb *EntityBag) Contains(e *Entity) bool {
	for _, ie := range eb.data {
		if ie == e {
			return true
		}
	}
	return false
}

func (eb *EntityBag) Remove(idx int) *Entity {
	ce := eb.data[idx] // make copy of element to remove so it can be returned
	eb.size--
	eb.data[idx] = eb.data[eb.size] // overwrite item to remove with last element
	eb.data[eb.size] = nil          // null last element, so gc can do its work
	return ce
}

func (eb *EntityBag) RemoveEntity(e *Entity) bool {
	for idx, ie := range eb.data {
		if ie == e {
			eb.Remove(idx)
			return true
		}
	}
	return false
}

func (eb *EntityBag) Set(idx int, e *Entity) {
	if idx >= len(eb.data) {
		eb.GrowSize(idx * 2)
	}
	eb.size = idx + 1
	eb.data[idx] = e
}

func (eb *EntityBag) Grow() {
	newCapacity := (len(eb.data)*3)/2 + 1
	eb.GrowSize(newCapacity)
}

func (eb *EntityBag) GrowSize(gsize int) {
	ndata := make([]*Entity, gsize)
	eb.data = append(eb.data, ndata...)
	// replace with copy? need test performance first
}

func (eb *EntityBag) RemoveLast() *Entity {
	if eb.IsEmpty() {
		return nil
	}
	return eb.Remove(eb.size - 1)
}

type ComponentBag struct {
	data []Component
	size int
}

func NewComponentBag(c int) *ComponentBag {
	return &ComponentBag{make([]Component, c), 0}
}

func (cb *ComponentBag) String() string {
	return fmt.Sprintf("%#v", cb.data)
}

func (cb *ComponentBag) Get(idx int) Component {
	return cb.data[idx]
}

func (cb *ComponentBag) Size() int {
	return cb.size
}

func (cb *ComponentBag) GetCapacity() int {
	return len(cb.data)
}

func (cb *ComponentBag) IsEmpty() bool {
	return cb.size == 0
}

func (cb *ComponentBag) Clear() {
	for i := 0; i < cb.size-1; i++ {
		cb.data[i] = nil
	}
}

func (cb *ComponentBag) Add(components ...Component) {
	for _, e := range components {
		if cb.size == len(cb.data) {
			cb.Grow()
		}
		cb.data[cb.size] = e
		cb.size++
	}
}

func (cb *ComponentBag) Contains(c Component) bool {
	for _, ic := range cb.data {
		if ic == c {
			return true
		}
	}
	return false
}

func (cb *ComponentBag) Remove(idx int) Component {
	cc := cb.data[idx] // make copy of element to remove so it can be returned
	cb.size--
	cb.data[idx] = cb.data[cb.size] // overwrite item to remove with last element
	cb.data[cb.size] = nil          // null last element, so gc can do its work
	return cc
}

func (cb *ComponentBag) RemoveComponent(c Component) bool {
	for idx, ic := range cb.data {
		if &ic == &c {
			cb.Remove(idx)
			return true
		}
	}
	return false
}

func (cb *ComponentBag) Set(idx int, c Component) {
	if idx >= len(cb.data) {
		cb.GrowSize(idx * 2)
	}
	cb.size = idx + 1
	cb.data[idx] = c
}

func (cb *ComponentBag) Grow() {
	newCapacity := (len(cb.data)*3)/2 + 1
	cb.GrowSize(newCapacity)
}

func (cb *ComponentBag) GrowSize(gsize int) {
	ndata := make([]Component, gsize)
	cb.data = append(cb.data, ndata...)
	// replace with copy? need test performance first
}

func (cb *ComponentBag) RemoveLast() Component {
	if cb.IsEmpty() {
		return nil
	}
	return cb.Remove(cb.size - 1)
}
