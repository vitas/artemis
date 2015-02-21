package aesf

import (
	"fmt"
)

//If you need to group your entities together, e.g. tanks going into "units" group or explosions into "effects",
//then use this manager. You must retrieve it using world instance.
//A entity can only belong to one group at a time.
type GroupManager struct {
	world           World
	emptyBag        *EntityBag
	entitiesByGroup map[string]*EntityBag
	groupByEntity   map[int]string
}

func NewGroupManager(w World) *GroupManager {
	gm := GroupManager{world: w}
	gm.entitiesByGroup = make(map[string]*EntityBag)
	gm.groupByEntity = make(map[int]string)
	gm.emptyBag = NewEntityBag(0)
	return &gm
}

func (gm *GroupManager) Initialize() {}

//Set the group of the entity.
func (gm *GroupManager) Set(group string, e *Entity) {
	gm.Remove(e) // Entity can only belong to one group.

	entities := gm.entitiesByGroup[group]
	if entities == nil {
		entities = NewEntityBag(COMPONENT_ACTIVE_BAG_CAP)
		gm.entitiesByGroup[group] = entities
	}
	entities.Add(e)
	gm.groupByEntity[e.GetID()] = group
}

func (gm GroupManager) String() string { return fmt.Sprintf("GroupManager") }

//Get all entities that belong to the provided group.
func (gm *GroupManager) getEntities(group string) *EntityBag {
	bag, ok := gm.entitiesByGroup[group]
	if !ok {
		return gm.emptyBag
	}
	return bag

}

//the name of the group that this entity belongs to, null if none.
func (gm GroupManager) GetGroupOf(e *Entity) string { return gm.groupByEntity[e.GetID()] }

//Checks if the entity belongs to any group.
func (gm GroupManager) IsGrouped(e *Entity) bool { return len(gm.GetGroupOf(e)) > 0 }

//Check if the entity is in the supplied group.
func (gm GroupManager) IsInGroup(group string, e *Entity) bool {
	return len(group) > 0 && group == gm.GetGroupOf(e)
}

//Removes the provided entity from the group it is assigned to, if any.
func (gm *GroupManager) Remove(e *Entity) {
	group, ok := gm.groupByEntity[e.GetID()]
	if ok {
		delete(gm.groupByEntity, e.GetID())

		entities, ok := gm.entitiesByGroup[group]
		if ok {
			entities.RemoveEntity(e)
		}
	}
}

//do nothing, just implement manager
func (tm *GroupManager) Refresh(e *Entity) {}
