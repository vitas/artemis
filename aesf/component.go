package aesf

import (
	"fmt"
)

type CTypeName string

const CTYPE_NAME_UNKNOWN = "Unknown"

//A tag component. All components in the system must extend this class.
type Component interface {
	GetCType() CTypeName
}

func (ctn CTypeName) String() string { return fmt.Sprintf("CTypeName[%s]", ctn) }

type ComponentType struct {
	typename CTypeName
	id       int
	bit      int64
}

func NewComponentType(ctype CTypeName, nextId int, nextBit int64) *ComponentType {
	return &ComponentType{ctype, nextId, nextBit}
}

func (ct *ComponentType) GetBit() int64 { return ct.bit }
func (ct *ComponentType) GetID() int    { return ct.id }
func (ct ComponentType) String() string {
	return fmt.Sprintf("ComponentType[%s](%d)", ct.typename, ct.id)
}
