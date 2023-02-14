package screen

import (
	"log"

	"github.com/mattn/go-runewidth"

	"github.com/gdamore/tcell/v2"
)

var (
	DefaultScreen *Screen
	DefaultStyle  tcell.Style
)

type mode int

const (
	def mode = iota
	insert
	doneDoing
)

type Screen struct {
	tcell.Screen
	m       mode
	lineNum int
	exited  bool
}

func (s *Screen) Exit() {
	s.exited = true
}

func (s *Screen) Exited() bool {
	return s.exited
}

func (s *Screen) SetContent(x, y int, content string) {
	for _, c := range content {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.Screen.SetContent(x, y, c, comb, DefaultStyle)
		x += w
	}
}

func (s *Screen) InsertMode() *Screen {
	s.m = insert
	return s
}

func (s *Screen) NormalMode() *Screen {
	s.m = def
	return s
}

func (s *Screen) DoneDoingMode() *Screen {
	s.m = doneDoing
	return s
}

func (s *Screen) IsInsertMode() bool {
	return s.m == insert
}

func (s *Screen) IsNormalMode() bool {
	return s.m == def
}

func (s *Screen) IsDoneDoingMode() bool {
	return s.m == doneDoing
}

func (s *Screen) SetCursorLine(n int) *Screen {
	s.lineNum = n
	return s
}

func (s *Screen) CursorLine() int {
	return s.lineNum
}

func Init() {

	DefaultStyle = tcell.StyleDefault.
		Background(tcell.ColorReset).
		Foreground(tcell.ColorReset)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.SetStyle(DefaultStyle)
	s.EnablePaste()
	s.DisableMouse()
	s.Clear()

	DefaultScreen = &Screen{s, def, 1, false}

}
