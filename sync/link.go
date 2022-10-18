package sync

type Linker interface {
	ProvideSync() *Status
}

func NewLink(st chan Status, crbnr Carabiner) {
	var (
		watch    *Status
		oldTempo float32
		newTempo float32
	)
	watch.Bpm = 0.0
	go func() {
		for {
			oldTempo = watch.Bpm
			watch = crbnr.ProvideSync()
			newTempo = watch.Bpm
			if oldTempo != newTempo {
				watch.D = true
				oldTempo = newTempo
			}
			st <- *watch
			watch.D = false
		}
	}()

}
