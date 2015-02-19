package aesf

import (
	"fmt"
)

//Cannot be instantiated outside the framework,
//you must create new entities using World.
type Entity struct {
	id            int
	uniqueId      int
	systemBits    int64
	typeBits      int64
	world         World
	entityManager *EntityManager
}

// use thiss method to create entity
func NewEntity(world World, id int) *Entity {
	return &Entity{id, 0, 0, 0, world, world.GetEntityManager()}
}

func (e *Entity) GetID() int {
	return e.id
}

func (e *Entity) GetUniqueID() int {
	return e.uniqueId
}

func (e *Entity) AddTypeBit(bit int64) {
	e.typeBits |= bit
}

func (e *Entity) AddSystemBit(bit int64) {
	e.systemBits |= bit
}

func (e *Entity) RemoveSystemBit(bit int64) {
	e.typeBits &= ^bit
}

func (e *Entity) RemoveTypeBit(bit int64) {
	e.typeBits &= ^bit
}

func (e *Entity) GetSystemBits() int64 {
	return e.systemBits
}

func (e *Entity) GetTypeBits() int64 {
	return e.typeBits
}

func (e *Entity) Reset() {
	e.typeBits = 0
	e.systemBits = 0
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity[%d]", e.id)
}

//this is the preferred method to use when retrieving a component from a entity.
//It will provide good performance.
func (e *Entity) GetComponent(ctype *ComponentType) Component {
	return e.entityManager.GetComponent(e, ctype)
}

//Slower retrieval of components from this entity.
//Minimize usage of this, but is fine to use e.g. when creating new entities and setting data in components.
func (e *Entity) GetComponentByType(ctname CTypeName) Component {
	return e.GetComponent(gComponentTypeManager.getTypeFor(ctname))
}

// Get all components belonging to this entity.
// * WARNING. Use only for debugging purposes, it is dead slow.
// * WARNING. The returned bag is only valid until this method is called again, then it is overwritten.
func (e *Entity) GetAllComponents() *ComponentBag {
	return e.entityManager.GetComponents(e)
}

//Add a component to this entity.
func (e *Entity) AddComponent(c Component) {
	e.entityManager.AddComponent(e, c)
}

//Removes the component from this entity.
func (e *Entity) RemoveComponent(c Component) {
	e.entityManager.RemoveComponent(e, c)
}

//Faster removal of components from a entity.
func (e *Entity) RemoveComponentByType(ctype *ComponentType) {
	e.entityManager.RemoveComponentByType(e, ctype)
}

func (e *Entity) IsActive() bool {
	return e.entityManager.IsActive(e.GetID())
}

// Refresh all changes to components for this entity.
//After adding or removing components, you must call this method.
//It will update all relevant systems.
// It is typical to call this after adding components to a newly created entity.
func (e *Entity) Refresh() {
	e.world.RefreshEntity(e)
}

// Delete this entity from the world.
func (e *Entity) Delete() {
	e.world.DeleteEntity(e)
}

//Set the group of the entity. Same as World.setGroup().
func (e *Entity) SetGroup(group string) {
	//e.world.groupManager().Set(group, e)
}

// Assign a tag to this entity. Same as World.setTag().
func (e *Entity) SetTag(tag string) {
	//	e.world.tagManager().Register(tag, e)
}
