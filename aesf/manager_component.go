// Package main provides ...
package aesf

import (
	//	"github.com/willf/bitset"
	"fmt"
	"reflect"
)

type ComponentManager struct {
	//	componentsByType map[*EntityBag
	deletedEntity *EntityBag
}

func (cm *ComponentManager) Initialize() {
	cm.deletedEntity = NewBag()
	v := reflect.TypeOf(em).Elem()
}

func (cm *ComponentManager) RemoveComponentsOfEntity(e Entity) {
	componentBits := e.GetComponentBits()
	for i := componentBits.nextSetBit(0); i >= 0; i = componentBits.nextSetBit(i + 1) {
		cm.componentsByType.Get(i).Set(e.GetID(), nil)
	}
	componentBits.ClearAll()

}

func (cm *ComponentManager) OnDeleted(e Entity) {
	cm.deletedEntity.Add(e)
}

func (cm *ComponentManager) Clean() {
	if cm.deletedEntity.Size() > 0 {
		for i := 0; cm.deletedEntity.Size() > i; i++ {
			cm.RemoveComponentsOfEntity(cm.deletedEntity.Get(i))
		}
		cm.deletedEntity.Clear()
	}
}
