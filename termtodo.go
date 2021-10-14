package main

import (
	"github.com/SCU-SJL/termtodo/core"
	"github.com/SCU-SJL/termtodo/screen"
	"github.com/SCU-SJL/termtodo/todolist"
)

func main() {

	screen.Init()
	defer screen.DefaultScreen.Fini()
	defer todolist.Save()

	core.Init(screen.DefaultScreen.SetCursorLine(1))
	core.Dispatcher.Dispatch()

}
