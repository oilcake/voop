package sync

type Link struct {
	st  chan Status
	lnk Linker
}

type Linker interface {
	ProvideSync() Status
}

func NewLink(lnk Linker) *Link {
	l := &Link{
		st:  make(chan Status, 3),
		lnk: lnk,
	}
	l.spin()
	return l
}

func (l *Link) spin() {
	var (
		watch    Status
		oldTempo float64
		newTempo float64
	)
	watch.Bpm = 0.0
	go func() {
		for {
			oldTempo = watch.Bpm
			watch = l.lnk.ProvideSync()
			newTempo = watch.Bpm
			if oldTempo != newTempo {
				watch.D = true
				oldTempo = newTempo
			}
			l.st <- watch
			watch.D = false
		}
	}()
}

func (l *Link) Dock() chan Status {
	return l.st
}
