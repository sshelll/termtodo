package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termtodo/screen"
)

var (
	Dispatcher *dispatcher
	keyBinder  *KeyBinder
)

func Init(s *screen.Screen) {
	keyBinder = new(KeyBinder)
	bindKeys(s)
	Dispatcher = &dispatcher{
		s:         s,
		keyBinder: keyBinder,
	}
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

}
