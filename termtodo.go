package main

import (
	"flag"
	"fmt"

	"github.com/sshelll/termtodo/core"
	"github.com/sshelll/termtodo/screen"
	"github.com/sshelll/termtodo/todolist"
)

var isHelp = false

func main() {

	parseFlag()
	if isHelp {
		helpInfo()
		return
	}

	screen.Init()
	defer screen.DefaultScreen.Fini()
	defer todolist.Save()

	core.Init(screen.DefaultScreen.SetCursorLine(1))
	core.Dispatcher.Dispatch()

}

func parseFlag() {
	flag.BoolVar(&isHelp, "h", false, "help info")
	flag.Parse()
}

func helpInfo() {
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
