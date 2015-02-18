package aesf_test

import (
	. "github.com/vitas/artemis/aesf"
	"testing"
)

type MockComponent struct {
	ctype CTypeName
}

func (mc MockComponent) GetCType() CTypeName {
	return "MockCTypeName"
}

func Init() (*ComponentBag, Component, Component, Component) {
	a := NewComponentBag(10)
	te1 := MockComponent{"Mock1"}
	te2 := MockComponent{"Mock2"}
	te3 := MockComponent{"Mock3"}

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
	if a.GetCapacity() != 10 {
		t.Errorf("Add(): Capacity is wrong, expected 10 was %d", a.GetCapacity())
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
	if a.GetCapacity() != 10 {
		t.Errorf("Add(): Capacity is wrong, expected  10 was %d", a.GetCapacity())
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
	if a.GetCapacity() != 16 {
		t.Errorf("Add(): Capacity is wrong, expected  16 was %d", a.GetCapacity())
	}
}
