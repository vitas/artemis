package aesf

import ()

// you can create your own entity system handling by implementing this.
type System interface {
	Initialize()
	Begin()
	Process()
	IsProcessing() bool
	End()
}
