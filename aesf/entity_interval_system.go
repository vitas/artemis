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
type IntervalEntitySystemImpl struct {
	entitySystem EntitySystem
	acc          int
	interval     int
}

// Create base interval entity system
func NewIntervalEntitySystem(w World, interval int, ctnames ...CTypeName) *IntervalEntitySystemImpl {
	ies := new(IntervalEntitySystemImpl)
	ies.entitySystem = NewEntitySystem(w, ctnames...)
	return ies
}
