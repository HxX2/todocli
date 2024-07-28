package todo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/HxX2/todo/pkg/pprint"
)

type Todo struct {
	filePath     string
	ListDone     bool
	ListUndone   bool
	ShowProgress bool
}

func Init() *Todo {
	todo := new(Todo)

	configDir := filepath.Join(os.Getenv("HOME"), ".config", "todo")
	filePath := filepath.Join(configDir, "todo.txt")

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			pprint.Error(fmt.Sprintf("Can't create config directory\n%s\n", err))
			return nil
		}

		file, err := os.Create(filePath)
		defer file.Close()
		if err != nil {
			pprint.Error(fmt.Sprintf("Can't create todo.txt file\n%s\n", err))
			return nil
		}

	} else if err != nil {
		pprint.Error(fmt.Sprintf("Can't check file\n%s\n", err))
		return nil
	}

	todo.filePath = filePath
	todo.ListDone = true
	todo.ListUndone = true
	todo.ShowProgress = false

	return todo
}
