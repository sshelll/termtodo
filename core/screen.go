package core

import (
	"github.com/sshelll/termtodo/screen"
	"github.com/sshelll/termtodo/todolist"
)

var hideDoneFlag bool

func showMain(s *screen.Screen) (nLines int) {

	lines := todolist.Content()

	for i, ln := range lines {
		s.SetContent(0, i, ln)
		if i == s.CursorLine() && s.IsNormalMode() {
			s.SetContent(0, i, ln+"👈")
		}
	}

	return len(lines)

}

func showDoneItems(s *screen.Screen) (nLines int) {
	return showLines(s, todolist.DoneContent())
}

func showDoingItems(s *screen.Screen) (nLines int) {
	return showLines(s, todolist.DoingContent())
}

func showLines(s *screen.Screen, lines []string) (nLines int) {

	s.Clear()

	for i, ln := range lines {
		s.SetContent(0, i, ln)
	}

	return len(lines)

}
