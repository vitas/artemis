// Package main provides ...
package aesf

import (
	bitset "github.com/willf/bitset"
)

const COMPONENT_BAG_CAP = 16
const COMPONENT_ACTIVE_BAG_CAP = 64
const COMPONENT_REUSED_BAG_CAP = 32
const COMPONENT_BY_TYPE_CAP = 64

type EntityManager struct {
	nextAvailableId     int
	world               World
	count               int
	uniqueEntityId      int
	totalCreated        int
	totalRemoved        int
	disabled            *bitset.BitSet
	entities            *EntityBag
	identPool           *IdentifierPool
	activeEntities      *EntityBag
	removedAndAvailable *EntityBag
	entityComponents    *ComponentBag // Added for debug support.
	componentsByType    map[int]*ComponentBag
}

func NewEntityManager() *EntityManager {
	em := new(EntityManager)
	em.identPool = &IdentifierPool{NewIntList(), 0}
	em.activeEntities = NewEntityBag(COMPONENT_ACTIVE_BAG_CAP)
	em.removedAndAvailable = NewEntityBag(COMPONENT_REUSED_BAG_CAP)
	em.componentsByType = make(map[int]*ComponentBag, COMPONENT_BY_TYPE_CAP)
	em.entityComponents = NewComponentBag(COMPONENT_BAG_CAP)
	return em
}

func (em *EntityManager) Initialize() {

}

func (em *EntityManager) SetWorld(w World) {
	em.world = w
}
func (em *EntityManager) GetWorld() World {
	return em.world
}

func (em *EntityManager) Create() *Entity {
	e := em.removedAndAvailable.RemoveLast()
	if e == nil {
		em.nextAvailableId++
		//e := NewEntity(em.GetWorld(), em.identPool.CheckOut())
		e = NewEntity(em.GetWorld(), em.nextAvailableId)
		//e = new Entity(world, nextAvailableId++);
	} else {
		e.Reset()
	}
	em.uniqueEntityId++
	e.uniqueId = em.uniqueEntityId
	em.activeEntities.Set(e.GetID(), e)
	em.count++
	em.totalCreated++
	return e
}

func (em *EntityManager) GetEntity(entityId int) *Entity {
	return em.activeEntities.Get(entityId)
}

//TODO
func (em *EntityManager) Refresh(e *Entity) {
	/*systemManager := em.world.getSystemManager();
	Bag<EntitySystem> systems = systemManager.GetSystems();
	for(int i = 0, s=systems.size(); s > i; i++) {
		systems.Get(i).change(e);
	}
	*/
}

func (em *EntityManager) Remove(e *Entity) {
	em.activeEntities.Set(e.GetID(), nil)
	e.typeBits = 0
	em.Refresh(e)
	em.RemoveComponentsOfEntity(e)
	em.count--
	em.totalRemoved++
	//em.identPool.CheckIn(e.GetID())
	em.removedAndAvailable.Add(e)
}

func (em *EntityManager) RemoveComponentsOfEntity(e *Entity) {
	for _, eb := range em.componentsByType {
		if eb != nil && e.GetID() < eb.Size() {
			eb.Set(e.GetID(), nil)
		}
	}
}

//Check if this entity is active, or has been deleted, within the framework.
func (em *EntityManager) IsActive(id int) bool {
	return em.activeEntities.Get(id) != nil
}

//how many entities have been created since start.
func (em *EntityManager) GetTotalCreated() int {
	return em.totalCreated
}

//Get how many entities are active in this world.
func (em *EntityManager) GetActiveEntityCount() int {
	return em.count
}

//how many entities have been removed since start.
func (em *EntityManager) GetTotalRemoved() int {
	return em.totalRemoved
}

func (em *EntityManager) GetComponent(e *Entity, ctype *ComponentType) Component {
	components := em.componentsByType[ctype.GetID()]
	if components != nil && e.GetID() < components.Size() {
		return components.Get(e.GetID())
	}
	return nil
}

// Get all components belonging to this entity.
// * WARNING. Use only for debugging purposes, it is dead slow.
// * WARNING. The returned bag is only valid until this method is called again, then it is overwritten.
func (em *EntityManager) GetComponents(e *Entity) *ComponentBag {
	//clear, and GC
	em.entityComponents = NewComponentBag(COMPONENT_BAG_CAP)
	for _, eb := range em.componentsByType {
		if eb != nil && e.GetID() < eb.Size() {
			component := eb.Get(e.GetID())
			if component != nil {
				em.entityComponents.Add(component)
			}

		}
	}
	return em.entityComponents
}

func (em *EntityManager) AddComponent(e *Entity, c Component) {
	ctype := componentTypeManager.getTypeFor(c.GetCType())

	if ctype.GetID() >= len(em.componentsByType) {
		em.componentsByType[ctype.GetID()] = nil
	}
	components := em.componentsByType[ctype.GetID()]
	if components == nil {
		components = NewComponentBag(COMPONENT_BAG_CAP)
		em.componentsByType[ctype.GetID()] = components
	}
	components.Set(e.GetID(), c)
	e.AddTypeBit(ctype.GetBit())
}

func (em *EntityManager) RemoveComponent(e *Entity, c Component) {
	ctype := componentTypeManager.getTypeFor(c.GetCType())
	em.RemoveComponentByType(e, ctype)
}

func (em *EntityManager) RemoveComponentByType(e *Entity, ctype *ComponentType) {
	components := em.componentsByType[ctype.GetID()]
	if components != nil {
		components.Set(e.GetID(), nil)
		e.RemoveTypeBit(ctype.GetBit())
	}
}

// id generation pool
type IdentifierPool struct {
	ids             *IntList
	nextAvailableId int
}

func (ip *IdentifierPool) CheckIn(id int) {
	ip.ids.Add(id)
}

func (ip *IdentifierPool) CheckOut() int {
	ip.ids.Pop()
	ip.nextAvailableId++
	return ip.nextAvailableId
}
