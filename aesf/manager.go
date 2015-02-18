// Package main provides ...
package aesf

import ()

type EntityObserver interface {
	OnAdded(e Entity)
	OnChanged(e Entity)
	OnDeleted(e Entity)
	OnEnabled(e Entity)
	OnDisabled(e Entity)
}

type Manager interface {
	EntityObserver
	Initialize()
	SetWorld(w *World)
	GetWorld() World
}
