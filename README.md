# Termtodo
Terminal todo tool.  

### Install
`` go get github.com/SCU-SJL/termtodo``

### Run
run `` termtodo `` if GOPATH was configured  
or `` $GOPATH/bin/termtodo`` instead.  

### Attention:  
persistent file is ``~/local/lib/todo.yml``


### Help
|   Key           |  Desc  |
|  ----           |  ----  |
| Ctrl-r          | drop all changes since the program was run |
| Ctrl-k          | create new category |
| Ctrl-n          | create new todo and put it in the current category |
| Enter           | 1. fold / unfold category 2. change todo status |
| BackSpace / Del | del todo or category |
| Esc             | quit |
