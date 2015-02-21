package aesf

//A typical entity system. Use this when you need to process entities possessing the provided component types.
type EntitySystemProcessingImpl struct {
	*EntitySystemImpl
}

// Create base entity processing system
func NewEntitySystemProcessingImpl(w World, requiredCtname CTypeName, ctnames ...CTypeName) *EntitySystemProcessingImpl {
	eps := new(EntitySystemProcessingImpl)
	mctnames := GetMergedTypes(requiredCtname, ctnames...)
	eps.EntitySystemImpl = NewEntitySystem(w, mctnames...)
	return eps
}

func (em *EntitySystemProcessingImpl) Process(e *Entity)     {}
func (em *EntitySystemProcessingImpl) CheckProcessing() bool { return true }

func (ieps *EntitySystemProcessingImpl) ProcessEntities(entities *EntityBag) {
	for i := 0; entities.Size() > i; i++ {
		ieps.Process(entities.Get(i))
	}
}
