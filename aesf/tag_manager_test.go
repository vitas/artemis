package aesf_test

import (
	. "github.com/vitas/artemis/aesf"
	"testing"
)

type MockWorld struct{}

func (mw MockWorld) GetName() string                  { return "MockWorld" }
func (mw *MockWorld) Initialize()                     {}
func (mw MockWorld) GetEntityManager() *EntityManager { return nil }
func (mw MockWorld) GetSystemManager() *SystemManager { return nil }
func (mw MockWorld) RefreshEntity(e *Entity)          {}
func (mw MockWorld) DeleteEntity(e *Entity)           {}
func (mw MockWorld) CreateEntity() *Entity            { return nil }
func (mw MockWorld) GetEntity(id int) *Entity         { return nil }
func (mw MockWorld) GetManagers() []Manager           { return []Manager{} }
func (mw MockWorld) GetDelta() int                    { return 0 }
func (mw MockWorld) SetDelta(delta int)               {}
func (mw MockWorld) LoopStart()                       {}

func init() {

}

func TestTagManager(t *testing.T) {
	var (
		w  World
		tm *TagManager
	)
	w = &MockWorld{}
	tm = NewTagManager(w)
	e1, e2, e3, e4 := NewEntity(w, 1), NewEntity(w, 2), NewEntity(w, 3), NewEntity(w, 4)

	tm.Register("tag1", e1)
	tm.Register("tag2", e2)
	tm.Register("tag3", e3)
	tm.Register("tag4", e4)

	if len(tm.GetTags()) != 4 {
		t.Errorf("Register failed, expected 4 was %d", len(tm.GetTags()))
	}

	tm.Unregister("tag2")
	if len(tm.GetTags()) != 3 {
		t.Errorf("Unregister failed, expected 3 was %d", len(tm.GetTags()))
	}

	tm.Remove(e3)
	if len(tm.GetTags()) != 2 {
		t.Errorf("Remove failed, expected 2 was %d", len(tm.GetTags()))
	}

	//t.Errorf("%#v", tm.GetTags())
}
