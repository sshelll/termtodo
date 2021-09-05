package core

import (
	"github.com/SCU-SJL/termtodo/screen"
	"github.com/SCU-SJL/termtodo/todolist"
	"github.com/SCU-SJL/termtodo/util"
	"github.com/gdamore/tcell/v2"
)

type todoService struct {
	s     *screen.Screen
	input string
}

func (srv *todoService) pressRune(evk *tcell.EventKey) {
	s := srv.s
	if !s.IsInsertMode() {
		return
	}
	srv.input += util.GetKeyRune(evk.Name())
	srv.refreshInsertMode()
}

func (srv *todoService) pressDelOrBS(*tcell.EventKey) {
	s := srv.s
	input := srv.input

	if s.IsNormalMode() {
		cursorLn := s.CursorLine()
		todolist.DelByIdx(cursorLn - 1)
		if cursorLn > 1 {
			s.SetCursorLine(cursorLn - 1)
		}
		srv.refreshNormalMode()
		return
	}

	if s.IsInsertMode() {
		if len(input) == 0 {
			return
		}

		rs := []rune(input)
		rs = rs[:len(rs)-1]
		srv.input = string(rs)

		srv.refreshInsertMode()
	}

}

func (srv *todoService) pressEnter(*tcell.EventKey) {

	s := srv.s

	if s.IsNormalMode() { // switch item status
		curLn := s.CursorLine()
		todolist.SwitchStatusByIdx(curLn - 1)
		srv.refreshNormalMode()
		return
	}

	if s.IsInsertMode() { // save new item
		s.NormalMode()
		todolist.Add("default", srv.input)
		srv.input = ""
		srv.refreshNormalMode()
		return
	}

}

func (srv *todoService) pressCtrlN(*tcell.EventKey) {
	s := srv.s
	if s.IsNormalMode() {
		srv.input = ""
		s.InsertMode()
		srv.refreshInsertMode()
	}
}

func (srv *todoService) pressUpDown(evk *tcell.EventKey) {

	s := srv.s

	if !s.IsNormalMode() {
		return
	}

	k := evk.Key()

	curLn := s.CursorLine()

	flag := false
	doneCnt, todoCnt := todolist.Count()

	// 1st line is title
	if k == tcell.KeyUp && curLn > 1 {
		s.SetCursorLine(curLn - 1)
		flag = true
	} else if k == tcell.KeyDown && curLn < util.If(hideDoneFlag, todoCnt, todoCnt+doneCnt).(int) {
		s.SetCursorLine(curLn + 1)
		flag = true
	}

	if !flag {
		return
	}

	srv.refreshNormalMode()
}

func (srv *todoService) pressTab(*tcell.EventKey) {
	s := srv.s
	if !s.IsNormalMode() {
		return
	}

	hideDoneFlag = !hideDoneFlag

	s.SetCursorLine(1)

	srv.refreshNormalMode()
}

func (srv *todoService) pressCtrlR(*tcell.EventKey) {
	if !srv.s.IsNormalMode() {
		return
	}
	util.WithFatalf(todolist.Reload, "reload")
	srv.refreshNormalMode()
}

func (srv *todoService) refreshInsertMode() {
	srv.s.Clear()
	srv.s.SetContent(0, setMainContent(srv.s), "New: "+srv.input)
}

func (srv *todoService) refreshNormalMode() {
	s := srv.s
	s.Clear()
	setMainContent(s)
}
