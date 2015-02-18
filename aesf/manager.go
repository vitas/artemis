// Package main provides ...
package aesf

import ()

type Manager interface {
	Initialize()
	SetWorld(w World)
	GetWorld() World
	Refresh(e *Entity)
	Remove(e *Entity)
}
