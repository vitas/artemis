package aesf

var (
	gComponentTypeManager ComponentTypeManager
)

//package initialization
func init() {
	gComponentTypeManager = ComponentTypeManager{make(map[CTypeName]*ComponentType),
		make(map[CTypeName]int), make(map[CTypeName]int64)}
}

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

func (ctm ComponentTypeManager) getBIT(ctname CTypeName) int64 { return ctm.getTypeFor(ctname).GetBit() }
func (ctm ComponentTypeManager) getID(ctname CTypeName) int    { return ctm.getTypeFor(ctname).GetID() }

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

type ComponentMapper struct {
	ctype         *ComponentType
	ctypename     CTypeName
	entityManager *EntityManager
}

func NewComponentMapper(ctname CTypeName, w World) *ComponentMapper {
	ctype := gComponentTypeManager.getTypeFor(ctname)
	return &ComponentMapper{ctype, ctname, w.GetEntityManager()}
}

func (cm *ComponentMapper) Get(e *Entity) Component { return cm.entityManager.GetComponent(e, cm.ctype) }
