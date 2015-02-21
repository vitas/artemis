package aesf

var (
	gSystemBitManager SystemBitManager
)

//static initialization , package level
func init() {
	gSystemBitManager = SystemBitManager{0, make(map[CTypeName]int64)}
}

//static struct, global in the package
type SystemBitManager struct {
	POS uint32
	//CTName of EntitySystem
	systemBits map[CTypeName]int64
}

func (sbm SystemBitManager) GetBitFor(ctname CTypeName) int64 {
	bit := sbm.systemBits[ctname]
	if bit == 0 {
		var b int64
		b = 1
		//    1L
		bit = b << sbm.POS
		sbm.POS++
		sbm.systemBits[ctname] = bit
	}
	return bit
}

//If you need to communicate with systems from other system, then look it up here.
//Use the world instance to retrieve a instance.
type SystemManager struct {
	world   World
	systems map[CTypeName]EntitySystem
	bagged  []EntitySystem
}

func NewSystemManager(w World) *SystemManager {
	sm := SystemManager{world: w}
	sm.systems = make(map[CTypeName]EntitySystem)
	return &sm
}

func (sm *SystemManager) Initialize()               {}
func (sm *SystemManager) Refresh(e *Entity)         {}
func (sm *SystemManager) Remove(e *Entity)          {}
func (sm SystemManager) GetSystems() []EntitySystem { return sm.bagged }

func (sm *SystemManager) SetSystem(entitySystem EntitySystem) EntitySystem {
	entitySystem.SetWorld(sm.world)

	//ignore error..likely nothing happens :)
	ctname, _ := ConvertTypeToString(entitySystem)
	sm.systems[ctname] = entitySystem

	isExists := false
	for _, sys := range sm.bagged {
		if sys == entitySystem {
			isExists = true
			break
		}
	}
	if !isExists {
		sm.bagged = append(sm.bagged, entitySystem)
	}
	b := gSystemBitManager.GetBitFor(ctname)
	entitySystem.SetSystemBit(b)

	return entitySystem
}

func (sm *SystemManager) InitializeAll() {
	for _, system := range sm.bagged {
		system.Initialize()
	}
}
