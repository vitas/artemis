package aesf

import ()

//extends EntitySystem
//A system that processes entities at a interval in milliseconds.
//A typical usage would be a collision system or physics system.
type IntervalEntitySystem interface {
	EntitySystem
	SetAcc(acc int)
	GetAcc() int
	SetInterval(interval int)
	GetInterval() int
}

//A system that processes entities at a interval in milliseconds.
//A typical usage would be a collision system or physics system.
type BaseIntervalEntitySystem struct {
	//deletega some work
	entitySystemHelper *EntitySystemHelper
	world              World
	acc                int
	interval           int
}

// Create base interval entity system
func NewBaseIntervalEntitySystemtySystem(w World, interval int, ctnames ...CTypeName) *BaseIntervalEntitySystem {
	ies := new(BaseIntervalEntitySystem)
	if w != nil {
		ies.world = w
	}
	ies.entitySystemHelper = NewEntitySystemHelper(ctnames...)
	return ies
}

func (em *BaseIntervalEntitySystem) SetWorld(w World) {
	em.world = w
}

//implements system interface
func (em *BaseIntervalEntitySystem) Begin() {}

//implements system interface
func (em *BaseIntervalEntitySystem) ProcessEntities(actives *EntityBag) {}

//implements system interface
func (em *BaseIntervalEntitySystem) IsProcessing() bool {
	return false
}

//implements system interface
func (em *BaseIntervalEntitySystem) End() {}

//implements system interface
func (em *BaseIntervalEntitySystem) Process() {
	if em.IsProcessing() {
		em.Begin()
		em.ProcessEntities(em.entitySystemHelper.actives)
		em.End()
	}
}

//implements EntitySystem
func (em *BaseIntervalEntitySystem) Remove(e *Entity) {
	em.entitySystemHelper.actives.RemoveEntity(e)
	e.RemoveSystemBit(em.entitySystemHelper.systemBit)
	em.Removed(e)
}

func (em *BaseIntervalEntitySystem) Change(e *Entity) {
	em.entitySystemHelper.Change(e)
}

// Called if the system has received a entity it is interested in, e.g. created or a component was added to it.
//@param e the entity that was added to this system.
func (em *BaseIntervalEntitySystem) Added(e *Entity) {}

// Called if a entity was removed from this system, e.g. deleted or had one of it's components removed.
// @param e the entity that was removed from this system.
func (em *BaseIntervalEntitySystem) Removed(e *Entity) {}
