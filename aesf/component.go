// Package main provides ...
package aesf

import (
	"fmt"
	"reflect"
)

var (
	componentTypeManager ComponentTypeManager
)

//package initialization
func init() {
	componentTypeManager = ComponentTypeManager{make(map[CTypeName]*ComponentType),
		make(map[CTypeName]int), make(map[CTypeName]int64)}
}

type CTypeName reflect.Value

// static stuct with some methods
type ComponentTypeManager struct {
	componentTypes map[CTypeName]*ComponentType
	ctIDs          map[CTypeName]int
	ctBITs         map[CTypeName]int64
}

func (ctm ComponentTypeManager) getTypeFor(ctname CTypeName) *ComponentType {
	ctype := ctm.componentTypes[ctname]
	if ctype == nil {
		ctype = NewComponentType(ctname, ctm.nextCtID(ctname), ctm.nextCtBIT(ctname))
		ctm.componentTypes[ctname] = ctype
	}
	return ctype
}

func (ctm ComponentTypeManager) getBIT(ctname CTypeName) int64 {
	return ctm.getTypeFor(ctname).GetBit()
}

func (ctm ComponentTypeManager) getID(ctname CTypeName) int {
	return ctm.getTypeFor(ctname).GetID()
}

func (ctm *ComponentTypeManager) nextCtID(ctname CTypeName) int {
	id := ctm.ctIDs[ctname]
	nId := id + 1
	ctm.ctIDs[ctname] = nId
	return id
}

func (ctm *ComponentTypeManager) nextCtBIT(ctname CTypeName) int64 {
	bit := ctm.ctBITs[ctname]
	if 0 == bit {
		bit++
	}
	nBit := bit << 1
	ctm.ctBITs[ctname] = nBit
	return bit
}

// abstract component
type Component interface{}

//create only via NewComponentType func
type ComponentType struct {
	typename CTypeName
	id       int
	bit      int64
}

func NewComponentType(ctype CTypeName, nextId int, nextBit int64) *ComponentType {
	return &ComponentType{ctype, nextId, nextBit}
}

func (ct *ComponentType) GetBit() int64 {
	return ct.bit
}

func (ct *ComponentType) GetID() int {
	return ct.id
}

func (ct ComponentType) String() string {
	return fmt.Sprintf("ComponentType[%s](%d)", ct.typename, ct.id)
}

type ComponentMapper struct {
	ctype         *ComponentType
	ctypename     CTypeName
	entityManager *EntityManager
}

func NewComponentMapper(ctname CTypeName, w *World) *ComponentMapper {
	ctype := componentTypeManager.getTypeFor(ctname)
	return &ComponentMapper{ctype, ctname, w.entityManager}
}

func (cm *ComponentMapper) Get(e Entity) Component {
	return cm.entityManager.getComponent(e, cm.ctype)
}
