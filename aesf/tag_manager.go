package aesf

import (
	"fmt"
)

//If you need to tag any entity, use this. A typical usage would be to tag
//entities such as "PLAYER". After creating an entity call register().
type TagManager struct {
	world       World
	entityByTag map[string]*Entity
}

func NewTagManager(w World) *TagManager {
	tm := TagManager{world: w}
	tm.entityByTag = make(map[string]*Entity)
	return &tm
}

func (em *TagManager) Initialize() {
}

func (tm *TagManager) GetTags() map[string]*Entity {
	return tm.entityByTag
}

func (tm TagManager) String() string {
	return fmt.Sprintf("TagManager")
}

func (tm *TagManager) Register(tag string, e *Entity) {
	tm.entityByTag[tag] = e
}

func (tm *TagManager) Unregister(tag string) {
	delete(tm.entityByTag, tag)
}

func (tm TagManager) IsRegistered(tag string) bool {
	_, ok := tm.entityByTag[tag]
	return ok
}
func (tm TagManager) GetEntity(tag string) *Entity {
	return tm.entityByTag[tag]
}

func (tm *TagManager) Remove(e *Entity) {
	for tag, ei := range tm.entityByTag {
		if ei == e {
			tm.Unregister(tag)
			break
		}
	}
}

//do nothing, just implement manager
func (tm *TagManager) Refresh(e *Entity) {}
