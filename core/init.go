package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termtodo/screen"
)

var (
	keyBinder *KeyBinder
)

func Start(s *screen.Screen) {
	keyBinder = new(KeyBinder)
	bindKeys(s)
	dispatcher := &dispatcher{
		s:         s,
		keyBinder: keyBinder,
	}
	dispatcher.Start()
}

func bindKeys(s *screen.Screen) {

	srv := &todoService{s: s}

	keyBinder.Bind(
		srv.pressEscOrCtrlC,
		tcell.KeyEscape, tcell.KeyESC, tcell.KeyEsc, tcell.KeyCtrlC,
	)

	keyBinder.Bind(
		srv.pressRune,
		tcell.KeyRune,
	)

	keyBinder.Bind(
		srv.pressDelOrBS,
		tcell.KeyDEL, tcell.KeyBackspace,
	)

	keyBinder.Bind(
		srv.pressEnter,
		tcell.KeyEnter,
	)

	keyBinder.Bind(
		srv.pressCtrlN,
		tcell.KeyCtrlN,
	)

	keyBinder.Bind(
		srv.pressCtrlK,
		tcell.KeyCtrlK,
	)

	keyBinder.Bind(
		srv.pressUpDown,
		tcell.KeyUp, tcell.KeyDown,
	)

	keyBinder.Bind(
		srv.pressTab,
		tcell.KeyTab, tcell.KeyTAB,
	)

	keyBinder.Bind(
		srv.pressCtrlR,
		tcell.KeyCtrlR,
	)

	keyBinder.Bind(
		srv.pressCtrlX,
		tcell.KeyCtrlX,
	)

	keyBinder.Bind(
		srv.pressCtrlZ,
		tcell.KeyCtrlZ,
	)

	keyBinder.Bind(
		srv.pressLeft,
		tcell.KeyLeft,
	)

	keyBinder.Bind(
		srv.pressRight,
		tcell.KeyRight,
	)

}
