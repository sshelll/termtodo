package core

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termtodo/screen"
	"github.com/sshelll/termtodo/todolist"
	"github.com/sshelll/termtodo/util"
)

type todoService struct {
	s                  *screen.Screen
	input              []rune
	inputCursorPos     int
	cateKey            string
	isCreatingCategory bool
}

func (srv *todoService) pressEscOrCtrlC(*tcell.EventKey) {

	s := srv.s

	if s.IsNormalMode() {
		s.Exit()
		return
	}

	srv.inputCursorPos = 0

	if s.IsInsertMode() {
		s.NormalMode()
		srv.input = nil
		srv.refreshNormalMode()
		return
	}

	if s.IsDoneDoingMode() {
		s.NormalMode()
		srv.input = nil
		srv.refreshNormalMode()
		return
	}

}

func (srv *todoService) pressRune(evk *tcell.EventKey) {

	s := srv.s
	if !s.IsInsertMode() {
		return
	}

	curRune := util.GetKeyRune(evk.Name())
	// clone old runes and append new rune, because append will change old runes
	newRunes := append(util.CloneRuneSlice(srv.input[:srv.inputCursorPos]), []rune(curRune)...)
	newRunes = append(newRunes, srv.input[srv.inputCursorPos:]...)
	srv.input = newRunes

	srv.incrInputCursorPos()
	srv.refreshInsertMode()

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

		newRunes := []rune(input)
		delPos := util.Max(srv.inputCursorPos-1, 0)
		delPos = util.Min(delPos, len(newRunes)-1)
		newRunes = append(newRunes[:delPos], newRunes[delPos+1:]...)

		srv.input = newRunes

		srv.decrInputCursorPos()

		srv.refreshInsertMode()

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

		// cursor is at item line -> change item status
		todolist.ChangeItemStatus(cateKey, offset)
		srv.refreshNormalMode()
		return

	}

	if s.IsInsertMode() { // save new item
		s.NormalMode()
		if srv.isCreatingCategory {
			todolist.AddNewCategory(string(srv.input))
		} else {
			todolist.AddNewItem(srv.cateKey, string(srv.input))
		}
		srv.input = nil
		srv.inputCursorPos = 0
		srv.refreshNormalMode()
		return
	}

}

func (srv *todoService) pressCtrlN(*tcell.EventKey) {
	s := srv.s
	if s.IsNormalMode() {
		srv.input = nil
		srv.isCreatingCategory = false
		srv.cateKey, _ = srv.getCategoryByCursor()
		s.InsertMode()
		srv.refreshInsertMode()
	}
}

func (srv *todoService) pressCtrlK(*tcell.EventKey) {
	s := srv.s
	if s.IsNormalMode() {
		srv.input = nil
		srv.isCreatingCategory = true
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

	// 1st line is title
	if k == tcell.KeyUp && curLn > 1 {
		s.SetCursorLine(curLn - 1)
		flag = true
	} else if k == tcell.KeyDown && curLn < todolist.ContentLineCnt()-1 {
		s.SetCursorLine(curLn + 1)
		flag = true
	}

	if !flag {
		return
	}

	srv.refreshNormalMode()

}

func (srv *todoService) pressLeft(*tcell.EventKey) {

	s := srv.s

	if !s.IsInsertMode() {
		return
	}

	srv.decrInputCursorPos()

	srv.refreshInsertMode()

}

func (srv *todoService) pressRight(*tcell.EventKey) {

	s := srv.s

	if !s.IsInsertMode() {
		return
	}

	srv.incrInputCursorPos()

	srv.refreshInsertMode()

}

// TODO: what the fuck is this? i forgot to finish this function in 2021???
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

	showDoneItems(srv.s)
	srv.s.DoneDoingMode()

}

func (srv *todoService) pressCtrlZ(*tcell.EventKey) {

	if !srv.s.IsNormalMode() {
		return
	}

	showDoingItems(srv.s)
	srv.s.DoneDoingMode()

}

func (srv *todoService) refreshInsertMode() {

	prefixMsg := fmt.Sprintf("New TODO - (%s): ", srv.cateKey)
	if srv.isCreatingCategory {
		prefixMsg = "New Category: "
	}

	srv.s.Clear()
	lineCnt := showMain(srv.s)
	srv.s.SetContent(0, lineCnt, prefixMsg+string(srv.input))

	// cursor pos x is at prefix msg + input cursor pos
	cellX := srv.cellCnt(prefixMsg) + srv.convPosToCell()
	srv.s.ShowCursor(cellX, lineCnt)

}

func (srv *todoService) refreshNormalMode() {

	s := srv.s
	s.Clear()
	showMain(s)

	srv.s.HideCursor()

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

func (srv *todoService) incrInputCursorPos() {
	srv.inputCursorPos = util.Min(srv.inputCursorPos+1, len(srv.input))
}

func (srv *todoService) decrInputCursorPos() {
	srv.inputCursorPos = util.Max(srv.inputCursorPos-1, 0)
}

func (srv *todoService) convPosToCell() int {
	return srv.cellCnt(string(srv.input[:srv.inputCursorPos]))
}

func (srv *todoService) cellCnt(s string) int {

	cnt := 0

	for _, r := range s {
		// some unicode char takes 2 cells, but its strlen is 3
		cnt += util.Min(2, util.StrLen(r))
	}

	return cnt

}
