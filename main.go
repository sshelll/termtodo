package main

import (
	"github.com/SCU-SJL/termtodo/core"
	"github.com/SCU-SJL/termtodo/screen"
	"github.com/SCU-SJL/termtodo/todolist"
	"os"
	"time"
)

func main() {

	screen.Init()
	defer screen.DefaultScreen.Fini()
	defer todolist.Save()

	//test()

	core.Init(screen.DefaultScreen.SetCursorLine(1))
	core.Dispatcher.Dispatch()

}

func test() {
	s := screen.DefaultScreen
	x, y := screen.DefaultScreen.Size()
	s.SetContent(x/2, y/2, "1. test")
	s.SetContent(x/2, y/2+1, "2. test")
	for {
		s.Show()
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}
}
