// Package main provides ...
package aesf

import ()

type Manager interface {
	Initialize()
	Refresh(e *Entity)
	Remove(e *Entity)
}
