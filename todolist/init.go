package todolist

import (
	"os"

	"github.com/sshelll/termtodo/util"
)

const PersistFileEnvKey = "TERMTODO_PERSIST_FILE"

var inst *todoList

func Init() {

	inst = new(todoList)

	inst.filePath = calPersistentFilePath()

	if _, err := os.Stat(inst.filePath); err == nil {
		util.WithFatalf(Reload, "reload")
	}

}

func calPersistentFilePath() (path string) {

	util.WithFatalf(func() error {

		filePath, err := util.JoinHomePath(`/.local/termtodo/todo.yml`)
		if err != nil {
			return err
		}

		if path, ok := os.LookupEnv(PersistFileEnvKey); ok {
			filePath = path
		}

		path = filePath

		return nil

	}, "cal persistent file path")

	return

}
