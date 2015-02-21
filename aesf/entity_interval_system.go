package aesf

//extends EntitySystem
//A system that processes entities at a interval in milliseconds.
//A typical usage would be a collision system or physics system.
type IntervalEntitySystem interface {
	EntitySystem
	SetAcc(acc int64)
	GetAcc() int64
	SetInterval(interval int64)
	GetInterval() int64
}

//A system that processes entities at a interval in milliseconds.
//A typical usage would be a collision system or physics system.
type IntervalEntitySystemImpl struct {
	*EntitySystemImpl
	acc      int64
	interval int64
}

// Create base interval entity system
func NewIntervalEntitySystem(w World, interval int64, ctnames ...CTypeName) *IntervalEntitySystemImpl {
	ies := new(IntervalEntitySystemImpl)
	ies.EntitySystemImpl = NewEntitySystem(w, ctnames...)
	return ies
}

func (ies *IntervalEntitySystemImpl) SetAcc(acc int64)           { ies.acc = acc }
func (ies IntervalEntitySystemImpl) GetAcc() int64               { return ies.acc }
func (ies *IntervalEntitySystemImpl) SetInterval(interval int64) { ies.interval = interval }
func (ies IntervalEntitySystemImpl) GetInterval() int64          { return ies.interval }

func (ies *IntervalEntitySystemImpl) CheckProcessing() bool {
	ies.acc += ies.GetWorld().GetDelta()
	if ies.acc >= ies.interval {
		ies.acc -= ies.interval
		return true
	}
	return false
}
