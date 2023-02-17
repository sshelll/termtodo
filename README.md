# Termtodo

Terminal todo tool.

## Install

`go install github.com/sshelll/termtodo@latest`

## Run

run `termtodo` if GOPATH was configured  
or `$GOPATH/bin/termtodo` instead.

## Attention

default persistent file is `~/.local/termtodo/todo.yml`

you can set `TERMTODO_PERSIST_FILE` env as your customized persistent file path.

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
