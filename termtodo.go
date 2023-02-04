package main

import (
	"github.com/sshelll/termtodo/core"
	"github.com/sshelll/termtodo/screen"
	"github.com/sshelll/termtodo/todolist"
)

func main() {

	screen.Init()
	defer screen.DefaultScreen.Fini()
	defer todolist.Save()

	core.Init(screen.DefaultScreen.SetCursorLine(1))
	core.Dispatcher.Dispatch()

}
