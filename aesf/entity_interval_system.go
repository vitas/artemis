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

func (em *IntervalEntitySystemImpl) GetWorld() World {
	return em.entitySystem.GetWorld()
}

func (ies *IntervalEntitySystemImpl) SetAcc(acc int) {
	ies.acc = acc
}

func (ies IntervalEntitySystemImpl) GetAcc() int {
	return ies.acc
}

func (ies *IntervalEntitySystemImpl) SetInterval(interval int) {
}

func (ies IntervalEntitySystemImpl) GetInterval() int {
	return ies.interval
}

func (ies *IntervalEntitySystemImpl) CheckProcessing() bool {
	ies.acc += ies.GetWorld().GetDelta()
	if ies.acc >= ies.interval {
		ies.acc -= ies.interval
		return true
	}
	return false
}
