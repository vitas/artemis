package aesf

import ()

type EntitySystem interface {
	Initialize()
	SetWorld(w World)
	SetSystemBit(bit int64)
	Change(e *Entity)
	Added(e *Entity)
	Removed(e *Entity)
}

// The most raw entity system.
// It should not typically be used, but you can create your own
//entity system handling by extending this.
// It is recommended that you use the other provided entity system implementations.
type BaseEntitySystem struct {
	world     World
	systemBit int64
	typeFlags int64
	actives   *EntityBag
}

// Create new entity system only via this func
func NewEntitySystem(w World, ctnames ...CTypeName) *BaseEntitySystem {
	bes := new(BaseEntitySystem)
	if w != nil {
		bes.world = w
	}
	bes.actives = NewEntityBag(16)
	bes.registerTypesNames(ctnames...)
	return bes
}

func (em *BaseEntitySystem) registerTypesNames(ctnames ...CTypeName) {
	for _, ctname := range ctnames {
		ctype := gComponentTypeManager.getTypeFor(ctname)
		em.typeFlags |= ctype.GetBit()
	}
}

func (em *BaseEntitySystem) SetWorld(w World) {
	em.world = w
}

func (em *BaseEntitySystem) SetSystemBit(bit int64) {
	em.systemBit = bit
}

//implements system interface
func (em *BaseEntitySystem) Begin() {}

//implements system interface
func (em *BaseEntitySystem) ProcessEntities(actives *EntityBag) {}

//implements system interface
func (em *BaseEntitySystem) IsProcessing() bool {
	return false
}

//implements system interface
func (em *BaseEntitySystem) End() {}

//implements system interface
func (em *BaseEntitySystem) Process() {
	if em.IsProcessing() {
		em.Begin()
		em.ProcessEntities(em.actives)
		em.End()
	}
}

func (em *BaseEntitySystem) Change(e *Entity) {
	contains := (em.systemBit & e.GetSystemBits()) == em.systemBit
	interest := (em.typeFlags & e.GetTypeBits()) == em.typeFlags
	if interest && !contains && em.typeFlags > 0 {
		em.actives.Add(e)
		e.AddSystemBit(em.systemBit)
		em.Added(e)
	} else if !interest && contains && em.typeFlags > 0 {
		em.Remove(e)
	}
}

func (em *BaseEntitySystem) Remove(e *Entity) {
	em.actives.RemoveEntity(e)
	e.RemoveSystemBit(em.systemBit)
	em.Removed(e)
}

// Called if the system has received a entity it is interested in, e.g. created or a component was added to it.
//@param e the entity that was added to this system.
func (em *BaseEntitySystem) Added(e *Entity) {}

// Called if a entity was removed from this system, e.g. deleted or had one of it's components removed.
// @param e the entity that was removed from this system.
func (em *BaseEntitySystem) Removed(e *Entity) {}

//TODO
//func (em *BaseEntitySystem) GetMergedTypes
