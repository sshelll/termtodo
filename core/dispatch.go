package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termtodo/screen"
	"github.com/sshelll/termtodo/todolist"
)

var hideDoneFlag bool

type dispatcher struct {
	s         *screen.Screen
	keyBinder *KeyBinder
}

func (d *dispatcher) Dispatch() {

	s := d.s

	showDefault(s)

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

func showDefault(s *screen.Screen) (nLines int) {

	lines := todolist.Content()

	for i, ln := range lines {
		s.SetContent(0, i, ln)
		if i == s.CursorLine() && s.IsNormalMode() {
			s.SetContent(0, i, ln+"👈")
		}
	}

	return len(lines)

}

func showDone(s *screen.Screen) (nLines int) {
	return doShowDoneDoing(s, todolist.DoneContent())
}

func showDoing(s *screen.Screen) (nLines int) {
	return doShowDoneDoing(s, todolist.DoingContent())
}

func doShowDoneDoing(s *screen.Screen, lines []string) (nLines int) {

	s.Clear()

	for i, ln := range lines {
		s.SetContent(0, i, ln)
	}

	return len(lines)

}
