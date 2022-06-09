package player

import (
	"voop/sync"
)

type Player struct {
	*sync.Clock
	*sync.Transport
	*Window
}
