package aesf_test

import (
	. "github.com/vitas/artemis/aesf"
	"testing"
)

type TestEntity struct {
	id       uint
	uniqueId uint
}

func (te TestEntity) GetID() uint {
	return te.id
}

func (te TestEntity) GetUniqueID() uint {
	return te.uniqueId
}

func (te TestEntity) Reset() {

}

func Init() (*EntityBag, Entity, Entity, Entity) {
	a := NewBag()
	te1 := TestEntity{1, 100}
	te2 := TestEntity{2, 200}
	te3 := TestEntity{3, 300}
	//e1 = te1
	a.Add(te1, te2, te3)
	return a, te1, te2, te3
}

func TestGet(t *testing.T) {
	a, e1, e2, e3 := Init()
	ae1 := a.Get(0)
	ae2 := a.Get(1)
	ae3 := a.Get(2)
	if e1 != ae1 {
		t.Errorf("Get(): Object is wrong, expected %v was %v", e1, ae1)
	}
	if e2 != ae2 {
		t.Errorf("Get(): Object is wrong, expected %v was %v", e2, ae2)
	}
	if e3 != ae3 {
		t.Errorf("Get(): Object is wrong, expected %v was %v", e3, ae3)
	}
}

func TestContains(t *testing.T) {
	a, e1, _, _ := Init()

	if a.Size() != 3 {
		t.Errorf("Add(): Size is wrong, expected  3 was %d", a.Size())
	}
	if a.GetCapacity() != 64 {
		t.Errorf("Add(): Capacity is wrong, expected  64 was %d", a.GetCapacity())
	}
	if !a.Contains(e1) {
		t.Errorf("Contains false %v", e1)
	}
}

func TestRemove(t *testing.T) {
	a, _, _, e3 := Init()
	ar := a.Remove(1)
	if a.Size() != 2 {
		t.Errorf("Add(): Size is wrong, expected  2 was %d", a.Size())
	}
	if a.GetCapacity() != 64 {
		t.Errorf("Add(): Capacity is wrong, expected  64 was %d", a.GetCapacity())
	}
	if a.Contains(ar) {
		t.Errorf("Contains true, not removed %v", ar)
	}
	// last must be on place of removed
	ae3 := a.Get(1)
	if e3 != ae3 {
		t.Errorf("Remove(): Object is wrong, expected %v was %v", e3, ae3)
	}
	arl := a.RemoveLast()
	if arl != ae3 {
		t.Errorf("RemoveLast(): Object is wrong, expected %v was %v", arl, ae3)
	}

}
func TestGrow(t *testing.T) {
	a, _, _, _ := Init()
	a.GrowSize(6)
	if a.Size() != 3 {
		t.Errorf("Add(): Size is wrong, expected  3 was %d", a.Size())
	}
	if a.GetCapacity() != 70 {
		t.Errorf("Add(): Capacity is wrong, expected  70 was %d", a.GetCapacity())
	}
}

// test UIntList
func TestUIntList(t *testing.T) {
	tilist := NewUIntList()
	tilist.Add(1)
	tilist.Add(3)
	tilist.Add(2)
	if tilist.Size() != 3 {
		t.Errorf("Add(): Size is wrong, expected  3 was %d", tilist.Size())
	}
	tilist.Pop()
	if tilist.Size() != 2 {
		t.Errorf("Pop(): Size is wrong, expected  2 was %d", tilist.Size())
	}
	tilist.Add([]uint{4, 5, 6}...)
	if tilist.Size() != 5 {
		t.Errorf("Add slice: Size is wrong, expected  5 was %d", tilist.Size())
	}
}
