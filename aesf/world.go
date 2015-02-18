// Package main provides ...
package aesf

import (
	"fmt"
)

type World interface {
	GetName()
	Initialize()
	GetEntityManager() *EntityManager
	RefreshEntity(e *Entity)
	DeleteEntity(e *Entity)
}

// The primary instance for the framework. It contains all the managers.
// You must use this to create, delete and retrieve entities.
// It is also important to set the delta each game loop iteration.
type EntityWorld struct {
	name          string
	entityManager *EntityManager
	delta         int
	refreshed     *EntityBag
	deleted       *EntityBag
	managers      []Manager
}

func NewEntityWorld() EntityWorld {
	w := EntityWorld{name: "EntityWorld"}
	w.entityManager = NewEntityManager()
	//	w.managers = []Manager{}
	w.managers = append(w.managers, w.entityManager)
	//TODO
	//groupManager
	//tagManager
	return w
}

func (w *EntityWorld) Initialize() {

}

func (w EntityWorld) String() string {
	return fmt.Sprintf("[%s]", w.name)
}

func (w EntityWorld) GetName() string {
	return w.name
}

//Get a entity having the specified id.
func (w EntityWorld) GetEntityManager() *EntityManager {
	return w.entityManager
}

//Ensure all systems are notified of changes to this entity.
func (w *EntityWorld) RefreshEntity(e *Entity) {
	w.refreshed.Add(e)
}

//Delete the provided entity from the world.
func (w *EntityWorld) DeleteEntity(e *Entity) {
	if !w.deleted.Contains(e) {
		w.deleted.Add(e)
	}
}

//Create and return a new or reused entity instance.
func (w *EntityWorld) CreateEntity() *Entity {
	return w.entityManager.Create()
}

//Get a entity having the specified id.
func (w *EntityWorld) GetEntity(id int) *Entity {
	return w.entityManager.GetEntity(id)
}

// Time since last game loop.
// delta in milliseconds.
func (w *EntityWorld) GetDelta() int {
	return w.delta
}

// You must specify the delta for the game here.
//delta time since last game loop.
func (w *EntityWorld) SetDelta(delta int) {
	w.delta = delta
}

// Let framework take care of internal business.
func (w *EntityWorld) LoopStart() {
	if !w.refreshed.IsEmpty() {
		for i := 0; w.refreshed.Size() > i; i++ {
			e := w.refreshed.Get(i)
			for _, manager := range w.managers {
				manager.Refresh(e)
			}
		}
		w.refreshed.Clear()
	}
	if !w.deleted.IsEmpty() {
		for i := 0; w.deleted.Size() > i; i++ {
			e := w.deleted.Get(i)
			for _, manager := range w.managers {
				manager.Remove(e)
			}
		}
		w.deleted.Clear()
	}
}
