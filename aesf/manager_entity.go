// Package main provides ...
package aesf

import (
	bitset "github.com/willf/bitset"
)

type EntityManager struct {
  nextAvailableId int
	world *World
	count    int
	uniqueEntityId     int
	totalCreated   int
	totalRemoved   int
	disabled  *bitset.BitSet
	entities  *EntityBag
	identPool *IdentifierPool
	entityComponents []Component // Added for debug support.
	activeEntities *EntityBag 
	removedAndAvailable *EntityBag 
	
	//private Bag<Bag<Component>> componentsByType;

}

func (em *EntityManager) Initialize() {
	em.identPool = &IdentifierPool{NewIntList(), 0}
	em.activeEntities = NewBag()
	em.removedAndAvailable = NewBag()
	em.entityComponents = []Component{}
}

func (em *EntityManager) SetWorld(w *World) {

}
func (em *EntityManager) GetWorld() *World {
	return nil
}

func (em *EntityManager) NewEntity() *BaseEntity {
	e := NewEntity(em.GetWorld(), em.identPool.CheckOut())
	em.added++
	return e
}

func (em *EntityManager) OnAdded(e Entity) {
	em.active++
	em.added++
	em.entities.Set(e.GetID(), e)
}

func (em *EntityManager) OnEnabled(e Entity) {
	em.disabled.Clear(e.GetID())
}

func (em *EntityManager) OnDisabled(e Entity) {
	em.disabled.Set(e.GetID())
}

func (em *EntityManager) OnDeleted(e Entity) {
	em.entities.Set(e.GetID(), nil)
	em.disabled.Clear(e.GetID())
	em.identPool.CheckIn(e.GetID())
	em.active--
	em.deleted++
}

func (em *EntityManager) IsActive(id uint) bool {
	return false
}

func (em *EntityManager) IsEnabled(id uint) bool {
	return !em.disabled.Test(id)
}


func (em *EntityManager) AddComponent(e Entity, c Component) {
	//Bag<Component> bag = componentsByType.get(type.getId());
	//if(bag != null && e.getId() < bag.getCapacity())
	//	return bag.get(e.getId());
}

func (em *EntityManager) GetComponent(e Entity, ctype *ComponentType) Component {
	//Bag<Component> bag = componentsByType.get(type.getId());
	//if(bag != null && e.getId() < bag.getCapacity())
	//	return bag.get(e.getId());
	return nil
}

func (em *EntityManager) GetComponents(e Entity) []Component {
		em.entityComponents.clear();
		for(a: = 0; componentsByType.size() > a; a++) {
			Bag<Component> components = componentsByType.get(a);
			if(components != null && e.getId() < components.size()) {
				Component component = components.get(e.getId());
				if(component != null) {
					entityComponents.add(component);
				}
			}
		}
		return entityComponents;
	}
}

func (em *EntityManager) RemoveComponent(e Entity, c Component) {

}

func (em *EntityManager) RemoveComponentByType(e Entity, ct ComponentType) {

}

func (em *EntityManager) GetEntity(entityId int) Entity{
		return em.activeEntities.Get(entityId);
  }

 //how many entities have been created since start.
func (em *EntityManager) GetTotalCreated() int {
	return em.created
}

//Get how many entities are active in this world.
func (em *EntityManager) GetActiveEntityCount() int{
		return em.count;
}


//how many entities have been removed since start.
func (em *EntityManager) GetTotalRemoved() int {
	return em.removed
}

// id generation pool
type IdentifierPool struct {
	ids             *IntList
	nextAvailableId uint
}

func (ip *IdentifierPool) CheckIn(id uint) {
	ip.ids.Add(id)
}

func (ip *IdentifierPool) CheckOut() uint {
	ip.ids.Pop()
	ip.nextAvailableId++
	return ip.nextAvailableId
}
