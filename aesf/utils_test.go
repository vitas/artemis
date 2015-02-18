package aesf_test

import (
	. "github.com/vitas/artemis/aesf"
	"testing"
)

// test IntList
func TestIntList(t *testing.T) {
	tilist := NewIntList()
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
	tilist.Add([]int{4, 5, 6}...)
	if tilist.Size() != 5 {
		t.Errorf("Add slice: Size is wrong, expected  5 was %d", tilist.Size())
	}
}
