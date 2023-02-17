package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sshelll/termtodo/core"
	"github.com/sshelll/termtodo/screen"
	"github.com/sshelll/termtodo/todolist"
	"github.com/sshelll/termtodo/util"
)

var (
	flagShowKeyMapping      = flag.Bool("k", false, "show key mapping")
	flagShowPersistFilePath = flag.Bool("p", false, "show persist file path")
)

func main() {

	if handleFlag() {
		return
	}

	todolist.Init()
	screen.Init()

	defer func() {
		r := recover()
		screen.DefaultScreen.Fini()
		todolist.Save()
		if r != nil {
			log.Fatalf("[termtodo] fatal error: %v", r)
		}
	}()

	core.Start(screen.DefaultScreen.SetCursorLine(1))

}

func handleFlag() (shouldExit bool) {

	flag.Parse()

	if util.BoolPtrVal(flagShowPersistFilePath) {
		fmt.Println(todolist.PersistFilePath())
		return true
	}

	if util.BoolPtrVal(flagShowKeyMapping) {
		printKeyMapping()
		return true
	}

	return false

}

func printKeyMapping() {
	fmt.Println(`| Key             | Desc                                                  |
| --------------- | ----------------------------------------------------- |
| Ctrl-r          | drop all changes since the program was run            |
| Ctrl-k          | create new category                                   |
| Ctrl-n          | create new todo and put it in the current category    |
| Ctrl-z          | show all doing items                                  |
| Ctrl-x          | show all done items                                   |
| Enter           | 1. fold / unfold category <br/> 2. change todo status |
| BackSpace / Del | del todo or category                                  |
| Esc             | quit                                                  |`)
}
