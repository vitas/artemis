package aesf

import ()

//callbacks
type EntitySystemEvent interface {
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
	processEntities(entities *EntityBag)
	SetWorld(w World)
	SetSystemBit(bit int64)
	Change(e *Entity)
	Remove(e *Entity)
}

// The most raw entity system.
// It is recommended that you use the other provided entity system implementations.
// delegate some work to it (see IntervalEntitySystem as an example)
type EntitySystemHelper struct {
	systemBit int64
	typeFlags int64
	actives   *EntityBag
}

// Create new entity system only via this func
func NewEntitySystemHelper(ctnames ...CTypeName) *EntitySystemHelper {
	bes := new(EntitySystemHelper)
	bes.actives = NewEntityBag(16)
	bes.registerTypesNames(ctnames...)
	return bes
}

func (em *EntitySystemHelper) registerTypesNames(ctnames ...CTypeName) {
	for _, ctname := range ctnames {
		ctype := gComponentTypeManager.getTypeFor(ctname)
		em.typeFlags |= ctype.GetBit()
	}
}

//implements EntitySystem
func (em *EntitySystemHelper) Change(e *Entity) {
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

// Merge together a required type and a array of other types. Used in derived systems.
func GetMergedTypes(requiredType CTypeName, otherTypes ...CTypeName) []CTypeName {
	itypes := append([]CTypeName{requiredType}, otherTypes[:1]...)
	return append(itypes, otherTypes[1:]...)
}
