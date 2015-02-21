package aesf

//callbacks
type EntitySystemEventDelegate interface {
	// Called if the system has received a entity it is interested in, e.g. created or a component was added to it.
	//@param e the entity that was added to this system.
	Added(e *Entity)
	// Called if a entity was removed from this system, e.g. deleted or had one of it's components removed.
	// @param e the entity that was removed from this system.
	Removed(e *Entity)
}

//extends System interface
type EntitySystem interface {
	System
	EntitySystemEventDelegate
	ProcessEntities(entities *EntityBag)
	SetWorld(w World)
	GetWorld() World
	SetSystemBit(bit int64)
	Change(e *Entity)
	Remove(e *Entity)
	SetEventDelegate(ev EntitySystemEventDelegate)
}

// The most raw entity system.
// It is recommended that you use the other provided entity system implementations.
// delegate some work to it (see IntervalEntitySystem as an example)
type EntitySystemImpl struct {
	world         World
	eventDelegate EntitySystemEventDelegate
	systemBit     int64
	typeFlags     int64
	actives       *EntityBag
}

// Create new entity system only via this func
func NewEntitySystem(world World, ctnames ...CTypeName) *EntitySystemImpl {
	es := new(EntitySystemImpl)
	es.SetWorld(world)
	es.actives = NewEntityBag(16)
	es.registerTypesNames(ctnames...)
	return es
}

func (em *EntitySystemImpl) SetWorld(w World)                              { em.world = w }
func (em *EntitySystemImpl) GetWorld() World                               { return em.world }
func (em *EntitySystemImpl) SetEventDelegate(ev EntitySystemEventDelegate) { em.eventDelegate = ev }
func (em *EntitySystemImpl) Initialize()                                   {}
func (em *EntitySystemImpl) Begin()                                        {}
func (em *EntitySystemImpl) ProcessEntities(actives *EntityBag)            {}
func (em *EntitySystemImpl) CheckProcessing() bool                         { return false }
func (em *EntitySystemImpl) End()                                          {}
func (em *EntitySystemImpl) SetSystemBit(bit int64)                        { em.systemBit = bit }

//implements system interface
func (em *EntitySystemImpl) Process() {
	if em.CheckProcessing() {
		em.Begin()
		em.ProcessEntities(em.actives)
		em.End()
	}
}

//implements EntitySystem
func (em *EntitySystemImpl) Remove(e *Entity) {
	em.actives.RemoveEntity(e)
	e.RemoveSystemBit(em.systemBit)
	em.eventDelegate.Removed(e)
}

//implements EntitySystem
func (em *EntitySystemImpl) Change(e *Entity) {
	contains := (em.systemBit & e.GetSystemBits()) == em.systemBit
	interest := (em.typeFlags & e.GetTypeBits()) == em.typeFlags
	if interest && !contains && em.typeFlags > 0 {
		em.actives.Add(e)
		e.AddSystemBit(em.systemBit)
		em.eventDelegate.Added(e)
	} else if !interest && contains && em.typeFlags > 0 {
		em.Remove(e)
	}
}

func (em *EntitySystemImpl) registerTypesNames(ctnames ...CTypeName) {
	for _, ctname := range ctnames {
		ctype := gComponentTypeManager.getTypeFor(ctname)
		em.typeFlags |= ctype.GetBit()
	}
}

// Merge together a required type and a array of other types. Used in derived systems.
func GetMergedTypes(requiredType CTypeName, otherTypes ...CTypeName) []CTypeName {
	itypes := append([]CTypeName{requiredType}, otherTypes[:1]...)
	return append(itypes, otherTypes[1:]...)
}

// Called if the system has received a entity it is interested in, e.g. created or a component was added to it.
//@param e the entity that was added to this system.
func (em *EntitySystemImpl) Added(e *Entity) {
	if em.eventDelegate != nil {
		em.eventDelegate.Added(e)
	}
}

// Called if a entity was removed from this system, e.g. deleted or had one of it's components removed.
// @param e the entity that was removed from this system.
func (em *EntitySystemImpl) Removed(e *Entity) {
	if em.eventDelegate != nil {
		em.eventDelegate.Removed(e)
	}
}
