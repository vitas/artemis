// Entity.
//Cannot be instantiated outside the framework
// must be created using World
package aesf

import (
	"fmt"
)

type Entity interface {
	GetID() uint
	GetUniqueID() uint
	GetSystemBits() int64
	GetTypeBits() int64
	AddTypeBit(bit int64)
	AddSystemBit(bit int64)
	RemoveTypeBit(bit int64)
	RemoveSystemBit(bit int64)
	Reset()
	AddComponent(c Component)
	RemoveComponent(c Component)
	IsActive() bool
}

type BaseEntity struct {
	id            uint
	uniqueId      uint
	systemBits    int64
	typeBits      int64
	world         *World
	entityManager *EntityManager
}

func NewEntity(world *World, id uint) *BaseEntity {
	return &BaseEntity{id, 0, 0, 0, world, world.entityManager}
}

func (e *BaseEntity) GetID() uint {
	return e.id
}

func (e *BaseEntity) GetUniqueID() uint {
	return e.uniqueId
}

func (e *BaseEntity) AddTypeBit(bit int64) {
	e.typeBits |= bit
}

func (e *BaseEntity) AddSystemBit(bit int64) {
	e.systemBits |= bit
}

func (e *BaseEntity) RemoveSystemBit(bit int64) {
	e.typeBits &= ^bit
}

func (e *BaseEntity) RemoveTypeBit(bit int64) {
	e.typeBits &= ^bit
}

func (e *BaseEntity) GetSystemBits() int64 {
	return e.systemBits
}

func (e *BaseEntity) GetTypeBits() int64 {
	return e.typeBits
}

func (e *BaseEntity) Reset() {
	e.typeBits = 0
	e.systemBits = 0
}

func (e *BaseEntity) String() string {
	return fmt.Sprintf("Entity[%d]", e.id)
}

//this is the preferred method to use when retrieving a component from a entity.
//It will provide good performance.
func (e *BaseEntity) GetComponent(ctype ComponentType) Component {
	return e.entityManager.GetComponent(e, &ctype)
}

//Slower retrieval of components from this entity.
//Minimize usage of this, but is fine to use e.g. when creating new entities and setting data in components.
func (e *BaseEntity) GetComponentByType(ct ComponentType) Component {
	//return e.entityManager.GetComponent(e, ct)
	return nil
}

// Get all components belonging to this entity.
// * WARNING. Use only for debugging purposes, it is dead slow.
// * WARNING. The returned bag is only valid until this method is called again, then it is overwritten.
func (e *BaseEntity) GetAllComponents() []Component {
	return e.entityManager.GetComponents()
}

func (e *BaseEntity) AddComponent(c Component) {
	e.entityManager.AddComponent(e, c)
}

func (e *BaseEntity) RemoveComponent(c Component) {
	e.entityManager.RemoveComponent(e, c)
}

func (e *BaseEntity) RemoveComponentByType(ct ComponentType) {
	e.entityManager.RemoveComponentByType(e, ct)
}

func (e *BaseEntity) IsActive() bool {
	return e.entityManager.IsActive(e.GetID())
}
