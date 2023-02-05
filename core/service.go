package core

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termtodo/screen"
	"github.com/sshelll/termtodo/todolist"
	"github.com/sshelll/termtodo/util"
)

type todoService struct {
	s               *screen.Screen
	input           string
	cateKey         string
	isInputCategory bool
}

func (srv *todoService) pressEscOrCtrlC(*tcell.EventKey) {
	s := srv.s

	if s.IsNormalMode() {
		s.Exit()
		return
	}

	if s.IsInsertMode() {
		s.NormalMode()
		srv.input = ""
		srv.refreshNormalMode()
		return
	}

	if s.IsDoneDoingMode() {
		s.NormalMode()
		srv.input = ""
		srv.refreshNormalMode()
		return
	}

}

func (srv *todoService) pressRune(evk *tcell.EventKey) {
	s := srv.s
	if !s.IsInsertMode() {
		return
	}

	srv.input += util.GetKeyRune(evk.Name())
	if srv.isInputCategory {
		srv.refreshInsertMode("New Category: ")
	} else {
		srv.refreshInsertMode(fmt.Sprintf("New TODO (%s): ", srv.cateKey))
	}
}

func (srv *todoService) pressDelOrBS(*tcell.EventKey) {
	s := srv.s
	input := srv.input

	if s.IsNormalMode() {
		cursorLn := s.CursorLine()
		cateKey, cursorOffset := srv.getCategoryByCursor()
		todolist.DelWithCategoryKey(cateKey, cursorOffset)

		// reset cursor to avoid over height limit of screen
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

		if srv.isInputCategory {
			srv.refreshInsertMode("New Category: ")
		} else {
			srv.refreshInsertMode(fmt.Sprintf("New TODO (%s): ", srv.cateKey))
		}
	}

}

func (srv *todoService) pressEnter(*tcell.EventKey) {

	s := srv.s

	if s.IsNormalMode() { // switch item status

		cateKey, offset := srv.getCategoryByCursor()
		if util.IsBlank(cateKey) {
			return
		}

		// cursor is at category line -> unfold / fold category
		if offset == 0 {
			todolist.SwitchFoldStatus(cateKey)
			srv.refreshNormalMode()
			return
		}

		// cursor is at item line -> remark item status
		todolist.RemarkItemStatus(cateKey, offset)
		srv.refreshNormalMode()
		return

	}

	if s.IsInsertMode() { // save new item
		s.NormalMode()
		if srv.isInputCategory {
			todolist.AddNewCategory(srv.input)
		} else {
			todolist.AddNewItem(srv.cateKey, srv.input)
		}
		srv.input = ""
		srv.refreshNormalMode()
		return
	}

}

func (srv *todoService) pressCtrlN(*tcell.EventKey) {
	s := srv.s
	if s.IsNormalMode() {
		srv.input = ""
		srv.isInputCategory = false
		srv.cateKey, _ = srv.getCategoryByCursor()
		s.InsertMode()
		srv.refreshInsertMode(fmt.Sprintf("New TODO (%s): ", srv.cateKey))
	}
}

func (srv *todoService) pressCtrlK(*tcell.EventKey) {
	s := srv.s
	if s.IsNormalMode() {
		srv.input = ""
		srv.isInputCategory = true
		s.InsertMode()
		srv.refreshInsertMode("New Category: ")
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

	// 1st line is title
	if k == tcell.KeyUp && curLn > 1 {
		s.SetCursorLine(curLn - 1)
		flag = true
	} else if k == tcell.KeyDown && curLn < len(todolist.Content())-1 { // TODO: do not call Content()
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
	srv.s.SetCursorLine(1)
	srv.refreshNormalMode()
}

func (srv *todoService) pressCtrlX(*tcell.EventKey) {
	if !srv.s.IsNormalMode() {
		return
	}
	showDone(srv.s)
	srv.s.DoneDoingMode()
}

func (srv *todoService) pressCtrlZ(*tcell.EventKey) {
	if !srv.s.IsNormalMode() {
		return
	}
	showDoing(srv.s)
	srv.s.DoneDoingMode()
}

func (srv *todoService) refreshInsertMode(msg string) {
	srv.s.Clear()
	srv.s.SetContent(0, showDefault(srv.s), msg+srv.input+"_")
}

func (srv *todoService) refreshNormalMode() {
	s := srv.s
	s.Clear()
	showDefault(s)
}

func (srv *todoService) getCategoryByCursor() (cateKey string, cursorOffset int) {
	s := srv.s

	curLn := s.CursorLine()

	cateKeys := todolist.CateKeys()

	// 1st line is title
	idx := 0

	for _, cateKey := range cateKeys {

		cate := todolist.GetCategory(cateKey)
		if cate == nil {
			idx++
			if idx == curLn {
				return cateKey, 0
			}
			continue
		}

		// cate title line
		idx++
		cateLn := idx

		// cate content line
		if !cate.Fold() {
			idx += cate.Size()
		}

		if idx >= curLn {
			return cate.Key(), curLn - cateLn
		}

	}

	return "", 0
}
