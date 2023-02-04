# Termtodo

Terminal todo tool.

## Install

`go get github.com/sshelll/termtodo`

## Run

run `termtodo` if GOPATH was configured  
or `$GOPATH/bin/termtodo` instead.

## Attention

persistent file is `~/.local/termtodo/todo.yml`

## Help

| Key             | Desc                                                  |
| --------------- | ----------------------------------------------------- |
| Ctrl-r          | drop all changes since the program was run            |
| Ctrl-k          | create new category                                   |
| Ctrl-n          | create new todo and put it in the current category    |
| Ctrl-z          | show all doing items                                  |
| Ctrl-x          | show all done items                                   |
| Enter           | 1. fold / unfold category <br/> 2. change todo status |
| BackSpace / Del | del todo or category                                  |
| Esc             | quit                                                  |
