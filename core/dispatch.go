package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termtodo/screen"
)

type dispatcher struct {
	s         *screen.Screen
	keyBinder *KeyBinder
}

func (d *dispatcher) Start() {

	s := d.s

	showMain(s)

	for {

		if s.Exited() {
			return
		}

		// update screen
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {

		case *tcell.EventResize:
			s.Sync()

		case *tcell.EventKey:
			if fn := d.keyBinder.Find(ev.Key()); fn != nil {
				fn(ev)
			}

		}

	}

}
