package core

import (
	"github.com/SCU-SJL/termtodo/screen"
	"github.com/SCU-SJL/termtodo/todolist"
	"github.com/gdamore/tcell/v2"
)

var hideDoneFlag bool

type dispatcher struct {
	s         *screen.Screen
	keyBinder *KeyBinder
}

func (d *dispatcher) Dispatch() {

	s := d.s

	show(s)

	for {

		if s.Exited() {
			return
		}

		// update screen
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {

		// resize cell
		case *tcell.EventResize:
			s.Sync()

		case *tcell.EventKey:
			if fn := d.keyBinder.Find(ev.Key()); fn != nil {
				fn(ev)
			}

		}

	}

}

func show(s *screen.Screen) (nLines int) {

	lines := todolist.Content()

	for i, ln := range lines {
		if i == s.CursorLine() && s.IsNormalMode() {
			s.SetContent(0, i, ln+"👈")
		}
		s.SetContent(0, i, ln)
	}

	return len(lines)

}
