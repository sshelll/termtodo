# Termtodo

Terminal todo tool.

## Install

`go install github.com/sshelll/termtodo@latest`

## Run

run `termtodo` if GOPATH was configured  
or `$GOPATH/bin/termtodo` instead.

## Attention

The default persistent file path is `~/.local/termtodo/todo.yml`

And you can set `TERMTODO_PERSIST_FILE` env as your customized persistent file path.

Thers's a hack usage of this env key,
such as you can set a file path which is in a `icloud` or `onedrive` directory,
and then you can make it sync. 😄

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
