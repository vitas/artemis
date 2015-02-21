package aesf

// If you need to process entities at a certain interval then use this.
// A typical usage would be to regenerate ammo or health at certain intervals, no need
// to do that every game loop, but perhaps every 100 ms. or every second.
type IntervalEntitySystemProcessingImpl struct {
	*IntervalEntitySystemImpl
	acc      int64
	interval int64
}

// Create base interval entity processing system
func NewIntervalEntitySystemProcessingImpl(w World, interval int64, requiredCtname CTypeName, ctnames ...CTypeName) *IntervalEntitySystemProcessingImpl {
	ieps := new(IntervalEntitySystemProcessingImpl)
	mctnames := GetMergedTypes(requiredCtname, ctnames...)
	ieps.IntervalEntitySystemImpl = NewIntervalEntitySystem(w, interval, mctnames...)
	return ieps
}

func (em *IntervalEntitySystemProcessingImpl) Process(e *Entity)     {}
func (em *IntervalEntitySystemProcessingImpl) CheckProcessing() bool { return true }

func (ieps *IntervalEntitySystemProcessingImpl) ProcessEntities(entities *EntityBag) {
	for i := 0; entities.Size() > i; i++ {
		ieps.Process(entities.Get(i))
	}
}
