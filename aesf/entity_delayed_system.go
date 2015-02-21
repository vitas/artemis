package aesf

type DelayedEntitySystem interface {
	EntitySystem
	StartDelayedRun(delay int64)
	ProcessEntitiesAccumulated(entities *EntityBag, accumulatedDelta int64)
	Stop()
}

// The purpose of this class is to allow systems to execute at varying intervals.
//
// An example system would be an ExpirationSystem, that deletes entities after a certain
// lifetime. Instead of running a system that decrements a timeLeft value for each
// entity, you can simply use this system to execute in a future at a time of the shortest
// lived entity, and then reset the system to run at a time in a future at a time of the
// shortest lived entity, etc.
//
// Another example system would be an AnimationSystem. You know when you have to animate
// a certain entity, e.g. in 300 milliseconds. So you can set the system to run in 300 ms.
// to perform the animation.
//
// This will save CPU cycles in some scenarios.
//
// Make sure you detect all circumstances that change. E.g. if you create a new entity you
// should find out if you need to run the system sooner than scheduled, or when deleting
// a entity, maybe something changed and you need to recalculate when to run. Usually this
// applies to when entities are created, deleted, chan.
type DelayedEntitySystemImpl struct {
	*EntitySystemImpl
	acc     int64
	delay   int64
	running bool
}

// Create base interval entity system
func NewDelayedEntitySystem(w World, ctnames ...CTypeName) *DelayedEntitySystemImpl {
	des := new(DelayedEntitySystemImpl)
	des.EntitySystemImpl = NewEntitySystem(w, ctnames...)
	return des
}

//Start processing of entities after a certain amount of milliseconds.
//Cancels current delayed run and starts a new one.
func (des *DelayedEntitySystemImpl) StartDelayedRun(delay int64) {
	des.delay = delay
	des.acc = 0
	des.running = true
}

func (des DelayedEntitySystemImpl) ProcessEntities(entities *EntityBag) {
	des.ProcessEntitiesAccumulated(entities, des.acc)
	des.Stop()
}

// The entities to process with accumulated delta.
func (des DelayedEntitySystemImpl) ProcessEntitiesAccumulated(entities *EntityBag, accumulatedDelta int64) {
}

func (des *DelayedEntitySystemImpl) CheckProcessing() bool {
	if des.running {
		des.acc += des.world.GetDelta()

		if des.acc >= des.delay {
			return true
		}
	}
	return false
}

//Aborts running the system in the future and stops it. Call delayedRun() to start it again.
func (des *DelayedEntitySystemImpl) Stop() {
	des.running = false
	des.acc = 0
}

//Check if the system is counting down towards processing.
func (des *DelayedEntitySystemImpl) IsRunning() bool { return des.running }

//Get the initial delay that the system was ordered to process entities after.
func (des *DelayedEntitySystemImpl) GetInitialTimeDelay() int64 { return des.delay }
func (des *DelayedEntitySystemImpl) GetRemainingTimeUntilProcessing() int64 {
	if des.running {
		return des.delay - des.acc
	}
	return 0
}
